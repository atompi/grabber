package options

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Version string = "v0.0.1"

type FileOptions struct {
	Src  string `yaml:"src"`
	Dest string `yaml:"dest"`
}

type SourceOptions struct {
	Url   string        `yaml:"url"`
	Auth  string        `yaml:"auth"`
	Files []FileOptions `yaml:"files"`
}

func NewOptions() (opts []SourceOptions) {
	optsSource := viper.Get("sources")
	err := createOptions(optsSource, &opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "create options failed:", err)
		os.Exit(1)
	}
	return
}
