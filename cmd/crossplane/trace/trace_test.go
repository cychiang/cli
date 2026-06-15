package trace

import (
	"testing"

	"github.com/alecthomas/kong"
	"github.com/google/go-cmp/cmp"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/crossplane-runtime/v2/pkg/test"

	"github.com/crossplane/cli/v2/cmd/crossplane/common/kube"
)

func TestCmd_getResourceAndName(t *testing.T) {
	type args struct {
		Resource string
		Name     string
	}

	type want struct {
		resource string
		name     string
		err      error
	}

	tests := map[string]struct {
		reason string
		fields args
		want   want
	}{
		"Splitted": {
			reason: "Should return the resource and name if both are provided",
			fields: args{
				Resource: "resource",
				Name:     "name",
			},
			want: want{
				resource: "resource",
				name:     "name",
				err:      nil,
			},
		},
		"Empty": {
			// should never happen, resource is required by kong
			reason: "Should return an error if no resource is provided",
			fields: args{
				Resource: "",
				Name:     "",
			},
			want: want{
				err: errors.New(errInvalidResource),
			},
		},
		"Combined": {
			reason: "Should return the resource and name if both are provided combined as resource",
			fields: args{
				Resource: "resource/name",
				Name:     "",
			},
			want: want{
				resource: "resource",
				name:     "name",
			},
		},
		"MoreSlashes": {
			reason: "Should return an error if the resource contains more than one slashes",
			fields: args{
				Resource: "resource/name/other",
				Name:     "",
			},
			want: want{
				err: errors.New(errInvalidResource),
			},
		},
		"BothAndCombined": {
			reason: "Should return an error if a name is provided both in the resource and separately",
			fields: args{
				Resource: "resource/name",
				Name:     "name",
			},
			want: want{
				err: errors.New(errNameDoubled),
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := &Cmd{
				Resource: tt.fields.Resource,
				Name:     tt.fields.Name,
			}

			gotResource, gotName, err := c.getResourceAndName()
			if diff := cmp.Diff(tt.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("Cmd.getResourceAndName() error = %v, wantErr %v", err, tt.want.err)
			}

			if diff := cmp.Diff(tt.want.resource, gotResource); diff != "" {
				t.Errorf("Cmd.getResourceAndName() resource = %v, want %v", gotResource, tt.want.resource)
			}

			if diff := cmp.Diff(tt.want.name, gotName); diff != "" {
				t.Errorf("Cmd.getResourceAndName() name = %v, want %v", gotName, tt.want.name)
			}
		})
	}
}

func TestImpersonationFlagsParse(t *testing.T) {
	cases := map[string]struct {
		reason string
		args   []string
		want   kube.ImpersonationFlags
	}{
		"None": {
			reason: "Without impersonation flags the fields should be empty.",
			args:   []string{"configuration.example.org"},
			want:   kube.ImpersonationFlags{},
		},
		"UserAndGroup": {
			reason: "--as and --as-group should populate the embedded flags.",
			args:   []string{"--as=jane", "--as-group=team-a", "configuration.example.org"},
			want:   kube.ImpersonationFlags{As: "jane", AsGroup: []string{"team-a"}},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var c Cmd

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
