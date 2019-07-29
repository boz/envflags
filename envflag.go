package envflags

import (
	"os"
	"strings"

	"github.com/spf13/pflag"
)

type Envflags interface {
	ApplyTo(flags *pflag.FlagSet) error
	ApplyWith(env []string, flags *pflag.FlagSet) error
}

type TransformFn func(name string) string
type NameFn func(prefix string, flagName string) string

type Config struct {
	Prefix      string
	TransformFn TransformFn
	NameFn      NameFn
	Override    bool
}

type envflags struct{ Config }

func New(cfg Config) Envflags {
	return &envflags{Config: cfg}
}

func (ef *envflags) ApplyTo(flags *pflag.FlagSet) error {
	return ef.ApplyWith(os.Environ(), flags)
}

func (ef *envflags) ApplyWith(env []string, flags *pflag.FlagSet) error {
	var err error

	envm := parseEnv(env)

	flags.VisitAll(func(flag *pflag.Flag) {
		if err != nil || (flag.Changed && !ef.Override) {
			return
		}

		name := ef.NameFn(ef.TransformFn(ef.Prefix), ef.TransformFn(flag.Name))
		if name == "" {
			return
		}

		if val, ok := envm[name]; ok {
			err = flag.Value.Set(val)
		}
	})

	return err
}

func parseEnv(env []string) map[string]string {
	envm := make(map[string]string, len(env))
	for _, entry := range env {
		parts := strings.Split(entry, "=")
		if len(parts) != 2 {
			continue
		}
		envm[parts[0]] = parts[1]
	}
	return envm
}
