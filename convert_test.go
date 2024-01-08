package virshjson

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    []map[string]any
		expectedErr error
	}{
		{
			name:        "no input results in a malformed input error",
			input:       "",
			expected:    nil,
			expectedErr: ErrMalformedHeader,
		},
		{
			name: "virsh list output can be parsed",
			input: ` Id   Name        State
---------------------------
 7    nix-visor   running

`,
			expected: []map[string]any{
				{
					"Id":    "7",
					"Name":  "nix-visor",
					"State": "running",
				},
			},
		},
		{
			name: "virsh domifaddr output can be parsed",
			input: ` Name       MAC address          Protocol     Address
-------------------------------------------------------------------------------
 vnet1      52:54:00:c9:ae:a5    ipv4         192.168.122.110/24

`,
			expected: []map[string]any{
				{
					"Name":        "vnet1",
					"MAC address": "52:54:00:c9:ae:a5",
					"Protocol":    "ipv4",
					"Address":     "192.168.122.110/24",
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := Convert(strings.NewReader(test.input))
			if err != test.expectedErr {
				t.Errorf("expected error %v, got %v", test.expectedErr, err)
			}
			if diff := cmp.Diff(test.expected, actual); diff != "" {
				t.Error(diff)
			}
		})
	}
}
