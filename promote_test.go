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
	"html/template"
	"testing"

	"github.com/coreos/go-semver/semver"
	"github.com/stretchr/testify/assert"
)

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

	assert.Equal(t, template.HTML(`elastic-package promote -d=snapshot-staging -n -p "foo-1.0.0,vaporware-1.0.0"`), cmd)
}
