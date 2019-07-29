package envflags

import (
	"os"
	"path"
	"strings"

	"github.com/spf13/pflag"
)

var (
	DefaultConfig Config
	DefaultFlags  *pflag.FlagSet
	DefaultEnv    []string
)

func init() {
	DefaultConfig = NewConfig()
	DefaultFlags = pflag.CommandLine
	DefaultEnv = os.Environ()
}

func NewConfig() Config {
	cfg := Config{
		TransformFn: defaultTransformFn,
		NameFn:      defaultNameFn,
		Override:    false,
	}
	if len(os.Args) > 0 {
		cfg.Prefix = defaultTransformFn(path.Base(os.Args[0]))
	}
	return cfg
}

func Apply() error {
	return ApplyTo(DefaultFlags)
}

func ApplyTo(flags *pflag.FlagSet) error {
	return ApplyWith(DefaultEnv, flags)
}

func ApplyWith(env []string, flags *pflag.FlagSet) error {
	return New(DefaultConfig).ApplyWith(env, flags)
}

func defaultNameFn(prefix string, flagName string) string {
	switch {
	case flagName == "":
		return ""
	case prefix == "":
		return flagName
	default:
		return prefix + "_" + flagName
	}
}

func defaultTransformFn(name string) string {
	return strings.Map(func(c rune) rune {
		switch {

		// [A-Z_] -> passthrough
		case c >= 'A' && c <= 'Z':
			return c
		case c == '_':
			return '_'

		// toupper
		case c >= 'a' && c <= 'z':
			return c + ('A' - 'a')

		// {'-','.'} -> '_'
		case c == '-' || c == '.':
			return '_'
		}

		return -1
	}, name)
}
