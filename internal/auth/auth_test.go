package auth

import (
	"reflect"
	"testing"
)

func TestSplitAuthHeader(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected []string
	}{
		"simple": {input: "ApiKey e84358b6-54e7-46eb-ac4c-f47f13610d0a", expected: []string{"ApiKey", "e84358b6-54e7-46eb-ac4c-f47f13610d0a"}},
	}

	for name, tc := range tests {
		got, _ := SplitAuthHeader(tc.input)
		if !reflect.DeepEqual(tc.expected, got) {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.expected, got)
		}
	}
}
