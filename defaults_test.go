package envflags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_defaultNameFn(t *testing.T) {
	tests := []struct {
		prefix   string
		flagName string
		name     string
	}{
		{"foo", "bar", "foo_bar"},
		{"", "foo", "foo"},
		{"foo", "", ""},
	}

	for _, test := range tests {
		assert.Equal(t, test.name, defaultNameFn(test.prefix, test.flagName))
	}
}

func Test_defaultTransformFn(t *testing.T) {
	tests := []struct {
		name   string
		result string
	}{
		{"foo", "FOO"},
		{"FOO", "FOO"},
		{"A_B", "A_B"},
		{"a-b", "A_B"},
		{"a.b", "A_B"},
		{"a?b", "AB"},
		{"_A", "_A"},
		{"A_", "A_"},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, defaultTransformFn(test.result))
	}
}
