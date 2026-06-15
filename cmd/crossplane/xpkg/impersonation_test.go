/*
Copyright 2026 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package xpkg

import (
	"testing"

	"github.com/alecthomas/kong"
	"github.com/google/go-cmp/cmp"

	"github.com/crossplane/cli/v2/cmd/crossplane/common/kube"
)

func TestInstallImpersonationFlagsParse(t *testing.T) {
	cases := map[string]struct {
		reason string
		args   []string
		want   kube.ImpersonationFlags
	}{
		"None": {
			reason: "Without impersonation flags the fields should be empty.",
			args:   []string{"provider", "example.org/provider-foo:v1.0.0"},
			want:   kube.ImpersonationFlags{},
		},
		"Group": {
			reason: "--as-group should populate the embedded flags.",
			args:   []string{"--as-group=team-a-admins", "provider", "example.org/provider-foo:v1.0.0"},
			want:   kube.ImpersonationFlags{AsGroup: []string{"team-a-admins"}},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var c installCmd

			p, err := kong.New(&c)
			if err != nil {
				t.Fatalf("%s\nkong.New(): unexpected error: %v", tc.reason, err)
			}

			if _, err := p.Parse(tc.args); err != nil {
				t.Fatalf("%s\nParse(%v): unexpected error: %v", tc.reason, tc.args, err)
			}

			if diff := cmp.Diff(tc.want, c.Impersonation); diff != "" {
				t.Errorf("%s\nParse(%v): -want, +got:\n%s", tc.reason, tc.args, diff)
			}
		})
	}
}

func TestUpdateImpersonationFlagsParse(t *testing.T) {
	cases := map[string]struct {
		reason string
		args   []string
		want   kube.ImpersonationFlags
	}{
		"None": {
			reason: "Without impersonation flags the fields should be empty.",
			args:   []string{"provider", "example.org/provider-foo:v1.0.1"},
			want:   kube.ImpersonationFlags{},
		},
		"User": {
			reason: "--as should populate the embedded flags.",
			args:   []string{"--as=jane", "provider", "example.org/provider-foo:v1.0.1"},
			want:   kube.ImpersonationFlags{As: "jane"},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var c updateCmd

			p, err := kong.New(&c)
			if err != nil {
				t.Fatalf("%s\nkong.New(): unexpected error: %v", tc.reason, err)
			}

			if _, err := p.Parse(tc.args); err != nil {
				t.Fatalf("%s\nParse(%v): unexpected error: %v", tc.reason, tc.args, err)
			}

			if diff := cmp.Diff(tc.want, c.Impersonation); diff != "" {
				t.Errorf("%s\nParse(%v): -want, +got:\n%s", tc.reason, tc.args, diff)
			}
		})
	}
}
