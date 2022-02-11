package main

import (
	"strings"
	"testing"
	"time"

	"github.com/coreos/go-semver/semver"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/stretchr/testify/assert"
)

func TestTemplateVersionUpdate(t *testing.T) {
	targetBranch := &Branch{
		Name:      "production",
		Commit:    "abc1234",
		Timestamp: time.Date(2022, 1, 11, 8, 9, 18, 0, time.UTC),

		worktree: &git.Worktree{Filesystem: osfs.New("testdata")},
	}

	packages := []Package{
		{
			PackageManifest: PackageManifest{
				Name:       "aws",
				Title:      "AWS",
				Version:    semver.Must(semver.NewVersion("1.8.0")),
				Conditions: Conditions{KibanaVersion: "^8.0.0"},
			},
			Changelog: []ReleaseChanges{
				{
					Version: semver.New("1.8.0"),
					Changes: []Change{
						{
							Type:        "enhancement",
							Description: "Update ECS",
							Link:        "https://github.com/elastic/integrations/pull/123",
						},
						{
							Type:        "bugfix",
							Description: "Fix bug",
							Link:        "https://github.com/elastic/integrations/pull/124",
						},
					},
				},
			},
		},
	}

	out, err := summarizeChanges(targetBranch, packages)
	if err != nil {
		t.Fatal(err)
	}

	const expected = `
## Summary of Changes

Comparisons were made to production branch commit
abc1234 from 2022-01-11 08:09:18 +0000 UTC.

- AWS - 1.8.0
  - Requires ^8.0.0
  - Changes since 1.7.0
    - 1.8.0
      - enhancement: Update ECS [PR](https://github.com/elastic/integrations/pull/123)
      - bugfix: Fix bug [PR](https://github.com/elastic/integrations/pull/124)
`

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(string(out)))
}

func TestTemplateNewPackage(t *testing.T) {
	targetBranch := &Branch{
		Name:      "production",
		Commit:    "abc1234",
		Timestamp: time.Date(2022, 1, 11, 8, 9, 18, 0, time.UTC),

		worktree: &git.Worktree{Filesystem: osfs.New("testdata")},
	}

	packages := []Package{
		{
			PackageManifest: PackageManifest{
				Name:       "vaporware",
				Title:      "Vaporware",
				Version:    semver.Must(semver.NewVersion("1.0.0")),
				Conditions: Conditions{KibanaVersion: "^8.0.0"},
			},
			Changelog: []ReleaseChanges{
				{
					Version: semver.New("1.0.0"),
					Changes: []Change{
						{
							Type:        "enhancement",
							Description: "Update ECS",
							Link:        "https://github.com/elastic/integrations/pull/123",
						},
					},
				},
			},
		},
	}

	out, err := summarizeChanges(targetBranch, packages)
	if err != nil {
		t.Fatal(err)
	}

	const expected = `
## Summary of Changes

Comparisons were made to production branch commit
abc1234 from 2022-01-11 08:09:18 +0000 UTC.

- Vaporware - 1.0.0
  - Requires ^8.0.0
  - New Package
    - 1.0.0
      - enhancement: Update ECS [PR](https://github.com/elastic/integrations/pull/123)
`

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(string(out)))
}

func TestCheckoutPackageStorage(t *testing.T) {
	repo, err := clonePackageStorage()
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, repo)
}

func TestMakePromoteCommand(t *testing.T) {
	cmd := makePromoteCommand("snapshot", "staging", []Package{
		{
			PackageManifest: PackageManifest{
				Name:    "foo",
				Version: semver.Must(semver.NewVersion("1.0.0")),
			},
		},
		{
			PackageManifest: PackageManifest{
				Name:    "vaporware",
				Version: semver.Must(semver.NewVersion("1.0.0")),
			},
		},
	})

	assert.Equal(t, `elastic-package promote -d=snapshot-staging -n -p "foo-1.0.0,vaporware-1.0.0"`, cmd)
}
