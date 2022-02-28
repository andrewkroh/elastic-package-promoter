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
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"

	"github.com/coreos/go-semver/semver"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var updateGolden = flag.Bool("update", false, "update golden files")

func TestTemplateSummaryOfChanges(t *testing.T) {
	originBranch := &Branch{
		Name:      "snapshot",
		Commit:    "fe2bc77a4f4294b236899b9b9c61d2d118e059b0",
		Timestamp: time.Date(2022, 1, 17, 11, 9, 3, 0, time.UTC),
	}

	targetBranch := &Branch{
		Name:      "production",
		Commit:    "abc1234",
		Timestamp: time.Date(2022, 1, 11, 8, 9, 18, 0, time.UTC),

		worktree: &git.Worktree{Filesystem: osfs.New("testdata")},
	}

	testTemplate(t, "new_package.md", &SummaryOfChanges{
		OriginBranch: originBranch,
		TargetBranch: targetBranch,
		Packages: []Package{
			{
				PackageManifest: PackageManifest{
					Name:       "vaporware",
					Title:      "Vaporware",
					Owner:      Owner{Github: "elastic/foo-team"},
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
		},
	})

	testTemplate(t, "with_changes.md", &SummaryOfChanges{
		OriginBranch: originBranch,
		TargetBranch: targetBranch,
		Packages: []Package{
			{
				PackageManifest: PackageManifest{
					Name:       "aws",
					Title:      "AWS",
					Owner:      Owner{Github: "elastic/foo-team"},
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
		},
	})

	testTemplate(t, "no_pending_changes.md", &SummaryOfChanges{
		OriginBranch: originBranch,
		TargetBranch: targetBranch,
		Packages:     nil,
	})
}

func testTemplate(t *testing.T, templateName string, soc *SummaryOfChanges) {
	t.Run(templateName, func(t *testing.T) {
		tmpl, err := getTemplate("embed:summary_of_changes.md.gohtml")
		require.NoError(t, err)

		out, err := renderTemplate(tmpl, soc)
		require.NoError(t, err)

		t.Log("Output:\n", string(out))

		goldenPath := filepath.Join("testdata/golden", templateName)
		if *updateGolden {
			require.NoError(t, ioutil.WriteFile(goldenPath, out, 0o644))
		}

		expected, err := ioutil.ReadFile(goldenPath)
		require.NoError(t, err)

		assert.Equal(t, string(expected), string(out))
	})
}
