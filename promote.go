// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package main

import (
	"embed"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/coreos/go-semver/semver"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/gomarkdown/markdown"
	"gopkg.in/yaml.v2"
)

const (
	appName          = "elastic-package-promote"
	remoteRepository = "https://github.com/elastic/package-storage"
)

// Build-time parameters.
var (
	version string
	commit  string
)

// Parameters
var (
	originBranch      string
	targetBranch      string
	teamFilter        string
	includeDeprecated bool
	templateFile      string
	markdownToHTML    bool
)

func init() {
	flag.StringVar(&originBranch, "origin", "snapshot", "origin branch")
	flag.StringVar(&targetBranch, "target", "production", "target branch")
	flag.StringVar(&teamFilter, "team", "", "select packages owned by this team (e.g. elastic/security-external-integrations)")
	flag.BoolVar(&includeDeprecated, "d", false, "include deprecated packages")
	flag.StringVar(&templateFile, "tmpl", "embed:summary_of_changes.md.gohtml", "template file to render")
	flag.BoolVar(&markdownToHTML, "md-to-html", false, "convert markdown to HTML")
}

//go:embed templates/*.gohtml
var templates embed.FS

var errPackageNotFound = errors.New("package not found")

type Package struct {
	PackageManifest
	Changelog Changelog
}

type PackageManifest struct {
	FormatVersion   string            `yaml:"format_version"`
	Name            string            `yaml:"name"`
	Title           string            `yaml:"title"`
	Version         *semver.Version   `yaml:"version"`
	License         string            `yaml:"license"`
	Description     string            `yaml:"description"`
	Type            string            `yaml:"type"`
	Categories      []string          `yaml:"categories"`
	Release         string            `yaml:"release"`
	Conditions      Conditions        `yaml:"conditions"`
	Screenshots     []Screenshots     `yaml:"screenshots"`
	Icons           []Icons           `yaml:"icons"`
	PolicyTemplates []PolicyTemplates `yaml:"policy_templates"`
	Owner           Owner             `yaml:"owner"`
}

type Conditions struct {
	KibanaVersion string `yaml:"kibana.version"`
}

type Screenshots struct {
	Src   string `yaml:"src"`
	Title string `yaml:"title"`
	Size  string `yaml:"size"`
	Type  string `yaml:"type"`
}

type Icons struct {
	Src   string `yaml:"src"`
	Title string `yaml:"title"`
	Size  string `yaml:"size"`
	Type  string `yaml:"type"`
}

type Vars struct {
	Name        string      `yaml:"name"`
	Type        string      `yaml:"type"`
	Title       string      `yaml:"title"`
	Description string      `yaml:"description,omitempty"`
	ShowUser    bool        `yaml:"show_user"`
	Required    bool        `yaml:"required"`
	Default     interface{} `yaml:"default,omitempty"`
	Multi       bool        `yaml:"multi,omitempty"`
}

type Inputs struct {
	Type        string `yaml:"type"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Vars        []Vars `yaml:"vars"`
}

type PolicyTemplates struct {
	Name        string   `yaml:"name"`
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Inputs      []Inputs `yaml:"inputs"`
}

type Owner struct {
	Github string `yaml:"github"`
}

type Changelog []ReleaseChanges

type ReleaseChanges struct {
	Version *semver.Version `json:"version"`
	Changes []Change        `json:"changes"`
}

type Change struct {
	Description string `json:"description"`
	Type        string `json:"type"`
	Link        string `json:"link"`
}

func (cl Changelog) ChangesSince(since *semver.Version) Changelog {
	var out Changelog
	for _, change := range cl {
		if change.Version.Compare(*since) > 0 {
			out = append(out, change)
		}
	}

	return out
}

func clonePackageStorage() (*git.Repository, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	url, err := url.Parse(remoteRepository)
	if err != nil {
		return nil, err
	}

	appDir := filepath.Join(home, "."+appName)
	repoName := path.Base(url.Path)
	repoDir := filepath.Join(appDir, "git", repoName)

	// Open or clone.
	repo, err := git.PlainOpen(repoDir)
	if errors.Is(err, git.ErrRepositoryNotExists) {
		log.Printf("Cloning %v into %v.", repoName, repoDir)
		repo, err = git.PlainClone(repoDir, false, &git.CloneOptions{
			URL: remoteRepository,
		})
	}
	if err != nil {
		return nil, err
	}

	log.Println("Fetching latest changes.")
	err = repo.Fetch(&git.FetchOptions{})
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		return nil, fmt.Errorf("failed in git fetch: %w", err)
	}
	log.Println("Fetch completed.")

	return repo, nil
}

func checkoutBranch(repo *git.Repository, branch string) error {
	remoteRef, err := repo.Reference(plumbing.NewRemoteReferenceName("origin", branch), false)
	if err != nil {
		return fmt.Errorf("failed to get remote reference for %v: %w", branch, err)
	}
	log.Printf("Branch <%s> is at commit %s", branch, remoteRef.Hash())

	wt, err := repo.Worktree()
	if err != nil {
		return err
	}

	log.Printf("Checking out %v.", branch)
	err = wt.Checkout(&git.CheckoutOptions{
		Hash: remoteRef.Hash(),
	})
	if err != nil {
		return fmt.Errorf("checkout failed: %w", err)
	}
	log.Println("Checkout completed.")

	log.Println("Cleaning repo.")
	err = wt.Clean(&git.CleanOptions{
		Dir: true,
	})
	if err != nil {
		return fmt.Errorf("clean failed: %w", err)
	}

	return nil
}

func listLatestPackages(wt *git.Worktree) ([]Package, error) {
	packageDirs, err := wt.Filesystem.ReadDir("packages")
	if err != nil {
		return nil, err
	}

	packages := make([]Package, 0, len(packageDirs))
	for _, dir := range packageDirs {
		p, err := listLatestPackage(wt, dir.Name())
		if err != nil {
			return nil, fmt.Errorf("failed to get latest package for %v: %w", dir.Name(), err)
		}

		packages = append(packages, *p)
	}

	return packages, nil
}

func listLatestPackage(wt *git.Worktree, packageName string) (*Package, error) {
	if packageName == "" {
		return nil, errors.New("package name must be non-empty")
	}

	packageDir := filepath.Join("packages", packageName)

	versionDirs, err := wt.Filesystem.ReadDir(packageDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errPackageNotFound
		}
		return nil, err
	}

	versions := make(semver.Versions, 0, len(versionDirs))
	for _, dir := range versionDirs {
		ver, err := semver.NewVersion(filepath.Base(dir.Name()))
		if err != nil {
			log.Printf("Ignoring invalid version directory %v: %v", filepath.Join(packageDir, dir.Name()), err)
			continue
		}

		versions = append(versions, ver)
	}

	if len(versions) == 0 {
		return nil, errors.New("no versions found")
	}
	sort.Sort(sort.Reverse(versions))

	latestVersion := versions[0].String()

	manifest, err := readPackageManifest(wt, packageName, latestVersion)
	if err != nil {
		return nil, err
	}
	cl, err := readChangelog(wt, packageName, latestVersion)
	if err != nil {
		return nil, err
	}

	return &Package{
		PackageManifest: *manifest,
		Changelog:       cl,
	}, nil
}

func readPackageManifest(wt *git.Worktree, packageName, version string) (*PackageManifest, error) {
	manifestPath := filepath.Join("packages", packageName, version, "manifest.yml")

	f, err := wt.Filesystem.Open(manifestPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var pm PackageManifest
	dec := yaml.NewDecoder(f)
	if err = dec.Decode(&pm); err != nil {
		return nil, err
	}

	return &pm, nil
}

func readChangelog(wt *git.Worktree, packageName, version string) (Changelog, error) {
	manifestPath := filepath.Join("packages", packageName, version, "changelog.yml")

	f, err := wt.Filesystem.Open(manifestPath)
	if err != nil {
		// Optional.
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	defer f.Close()

	var cl Changelog
	dec := yaml.NewDecoder(f)
	if err = dec.Decode(&cl); err != nil {
		return nil, err
	}

	return cl, nil
}

func filterByGithubOwner(team string, packages []Package) []Package {
	return filterPackages(packages, func(p *Package) bool {
		return p.Owner.Github == team
	})
}

func filterDeprecated(packages []Package) []Package {
	return filterPackages(packages, func(p *Package) bool {
		return !strings.HasPrefix(p.Description, "Deprecated")
	})
}

func filterPackages(packages []Package, predicate func(*Package) bool) []Package {
	var out []Package
	for _, p := range packages {
		if predicate(&p) {
			out = append(out, p)
		}
	}
	return out
}

type SummaryOfChanges struct {
	TeamFilter        string
	IncludeDeprecated bool
	OriginBranch      *Branch
	TargetBranch      *Branch
	Packages          []Package
	Version           string
}

type Branch struct {
	Name      string
	Commit    string
	Timestamp time.Time

	worktree *git.Worktree
}

func newBranch(repo *git.Repository, name string) (*Branch, error) {
	targetHead, err := repo.Head()
	if err != nil {
		return nil, err
	}

	// Latest commit.
	itr, err := repo.Log(&git.LogOptions{})
	if err != nil {
		return nil, err
	}

	commit, err := itr.Next()
	if err != nil {
		return nil, err
	}

	wt, err := repo.Worktree()
	if err != nil {
		return nil, err
	}

	return &Branch{
		Name:      name,
		Commit:    targetHead.Hash().String(),
		Timestamp: commit.Committer.When.UTC(),
		worktree:  wt,
	}, nil
}

func run() error {
	log.Printf("%s %s (%s)", appName, version, commit)

	tmpl, err := getTemplate(templateFile)
	if err != nil {
		return err
	}

	repo, err := clonePackageStorage()
	if err != nil {
		return err
	}

	if err = checkoutBranch(repo, originBranch); err != nil {
		return err
	}

	origin, err := newBranch(repo, originBranch)
	if err != nil {
		return err
	}

	wt, err := repo.Worktree()
	if err != nil {
		return err
	}

	packages, err := listLatestPackages(wt)
	if err != nil {
		return err
	}

	if !includeDeprecated {
		packages = filterDeprecated(packages)
	}
	if teamFilter != "" {
		packages = filterByGithubOwner(teamFilter, packages)
	}

	if len(packages) == 0 {
		log.Printf("No matching packages found on <%v>.", originBranch)
		return nil
	} else {
		log.Printf("Found %d matching packages on <%v>.", len(packages), originBranch)
	}

	for _, p := range packages {
		log.Println(p.Name, p.Version)
	}

	if err = checkoutBranch(repo, targetBranch); err != nil {
		return err
	}

	target, err := newBranch(repo, targetBranch)
	if err != nil {
		return err
	}

	out, err := renderTemplate(tmpl, &SummaryOfChanges{
		TeamFilter:        teamFilter,
		IncludeDeprecated: includeDeprecated,
		OriginBranch:      origin,
		TargetBranch:      target,
		Packages:          packages,
		Version:           version,
	})
	if err != nil {
		return err
	}

	if markdownToHTML {
		out = markdown.ToHTML(out, nil, nil)
	}

	fmt.Println(string(out))
	return nil
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
