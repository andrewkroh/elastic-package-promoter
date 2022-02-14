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
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"strings"

	"github.com/coreos/go-semver/semver"
)

var templateFuncs = map[string]interface{}{
	"latestVersion":      latestVersion,
	"changesSince":       changesSince,
	"makePromoteCommand": makePromoteCommand,
}

var htmlTmpl = template.Must(template.New("github").
	Option("missingkey=error").
	Funcs(templateFuncs).
	ParseFS(templates, "*/*.gohtml"))

func getTemplate(name string) (*template.Template, error) {
	if strings.HasPrefix(name, "embed:") {
		t := htmlTmpl.Lookup(strings.TrimPrefix(name, "embed:"))
		if t == nil {
			return nil, errors.New(name + " template not found")
		}
		return t, nil
	}

	data, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New("").
		Option("missingkey=error").
		Funcs(templateFuncs).
		Parse(string(data))
	if err != nil {
		return nil, fmt.Errorf("failed to parse %v: %w", name, err)
	}

	return tmpl, nil
}

func renderTemplate(tmpl *template.Template, data *SummaryOfChanges) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, data); err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.Bytes(), nil
}

func latestVersion(branch Branch, packageName string) (*Package, error) {
	p, err := listLatestPackage(branch.worktree, packageName)
	if err != nil {
		if errors.Is(err, errPackageNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return p, nil
}

func changesSince(p Package, sinceVersion *semver.Version) Changelog {
	return p.Changelog.ChangesSince(sinceVersion)
}

func makePromoteCommand(origin, target string, packages []Package) template.HTML {
	var sb strings.Builder
	sb.WriteString("elastic-package promote -d=")
	sb.WriteString(origin)
	sb.WriteString("-")
	sb.WriteString(target)
	sb.WriteString(` -n -p "`)
	for i, p := range packages {
		sb.WriteString(p.Name)
		sb.WriteString("-")
		sb.WriteString(p.Version.String())
		if i < len(packages)-1 {
			sb.WriteString(",")
		}
	}
	sb.WriteString(`"`)

	return template.HTML(sb.String())
}
