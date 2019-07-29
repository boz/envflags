package envflags

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func Test_ApplyWith(t *testing.T) {
	env := []string{
		"ENVFLAGS_TEST_FOO=value",
		"ENVFLAGS_TEST_BAR=100",
		"ENVFLAGS_TEST_QUUX=do-not-want",
	}

	flags := pflag.NewFlagSet("cli", pflag.ContinueOnError)

	flags.String("foo", "", "the foo")
	flags.Int("bar", 10, "the bar")
	flags.String("quux", "", "the quux")

	assert.NoError(t, flags.Parse([]string{
		"--quux", "quux-want",
	}))
	assert.NoError(t, ApplyWith(env, flags))

	{
		val, err := flags.GetString("foo")
		assert.NoError(t, err)
		assert.Equal(t, "value", val)
	}

	{
		val, err := flags.GetInt("bar")
		assert.NoError(t, err)
		assert.Equal(t, 100, val)
	}

	{
		val, err := flags.GetString("quux")
		assert.NoError(t, err)
		assert.Equal(t, "quux-want", val)
	}
}

func Test_defaultTransformFn(t *testing.T) {
	tests := [][]string{
		{"ab", "AB"},
		{"a-b", "A_B"},
		{"a.b", "A_B"},
		{"a b", "AB"},
		{"a/b", "AB"},
		{"", ""},
		{"~", ""},
		{"_", "_"},
		{"-", "_"},
	}
	for _, test := range tests {
		assert.Equal(t, test[1], defaultTransformFn(test[0]), test[0])
	}
}

func Test_defaultNameFn(t *testing.T) {
	tests := [][]string{
		{"a", "b", "a_b"},
		{"", "b", "b"},
		{"a", "", ""},
	}
	for _, test := range tests {
		assert.Equal(t, test[2], defaultNameFn(test[0], test[1]), test[0], test[1])
	}
}
