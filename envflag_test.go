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
	}

	flags := pflag.NewFlagSet("cli", pflag.ContinueOnError)
	flags.String("foo", "", "the foo")
	flags.Int("bar", 10, "the bar")
	flags.String("baz", "", "the baz")

	assert.NoError(t, flags.Parse([]string{}))
	assert.NoError(t, ApplyWith(env, flags))

	{
		assert.True(t, flags.Changed("foo"))
		val, err := flags.GetString("foo")
		assert.NoError(t, err)
		assert.Equal(t, "value", val)
	}

	{
		assert.True(t, flags.Changed("bar"))
		val, err := flags.GetInt("bar")
		assert.NoError(t, err)
		assert.Equal(t, 100, val)
	}

	{
		assert.False(t, flags.Changed("baz"))
		val, err := flags.GetString("baz")
		assert.NoError(t, err)
		assert.Equal(t, "", val)
	}
}

func Test_InvalidEnv(t *testing.T) {
	env := []string{
		"ENVFLAGS_TEST_FOO=value",
		"ENVFLAGS_TEST_BAR=not-an-int",
		"ENVFLAGS_TEST_BAZ=value",
	}

	flags := pflag.NewFlagSet("cli", pflag.ContinueOnError)
	flags.String("foo", "", "the foo")
	flags.Int("bar", 10, "the bar")
	flags.String("baz", "", "the baz")

	assert.NoError(t, flags.Parse([]string{}))
	assert.Error(t, ApplyWith(env, flags))

	assert.False(t, flags.Changed("bar"))
}

func Test_parseEnv(t *testing.T) {
	env := []string{
		"FOO=foo",
		"BAR=bar=bar",
		"BAZ=",
		"ZAB",
	}

	envm := parseEnv(env)

	assert.Equal(t, map[string]string{
		"FOO": "foo",
		"BAR": "bar=bar",
		"BAZ": "",
	}, envm)
}
