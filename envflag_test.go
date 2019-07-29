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

	assert.NoError(t, flags.Parse([]string{}))
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
}
