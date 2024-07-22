package cli

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/spf13/pflag"
	"github.com/suzuki-shunsuke/yaml2json/pkg/controller"
)

type Runner struct {
	Stdout  io.Writer
	LDFlags *LDFlags
}

type LDFlags struct {
	Version string
	Commit  string
	Date    string
}

func (flags *LDFlags) ShowVersion() string {
	if flags.Version == "" {
		if flags.Commit == "" {
			return ""
		}
		return flags.Commit
	}
	return flags.Version + " (" + flags.Commit + ")"
}

func parseFlag(verFlag, helpFlag *bool, indentFlag *string) {
	pflag.BoolVarP(verFlag, "version", "v", false, "show yaml2json's version")
	pflag.BoolVarP(helpFlag, "help", "h", false, "show the help message")
	pflag.StringVarP(indentFlag, "indent", "i", "", "indent")
	pflag.Parse()
}

const helpMsg = `NAME:
   yaml2json - Convert YAML to JSON

   https://github.com/suzuki-shunsuke/yaml2json

USAGE:
   yaml2json [--help, -h] [--version, -v] [--indent, -i ""]

VERSION:
   %s

OPTIONS:
   --help, -h           show help (default: false)
   --version, -v        print the version (default: false)
	 --indent, -i         indent (default: "")`

var errFileRequired = errors.New("YAML file is required")

func (runner *Runner) Run(ctx context.Context, _ ...string) error {
	ctrl := controller.New(os.Stdout)
	param := &controller.RunParam{}
	verFlag, helpFlag := false, false
	parseFlag(&verFlag, &helpFlag, &param.Indent)
	if helpFlag {
		fmt.Fprintf(runner.Stdout, helpMsg, runner.LDFlags.ShowVersion())
		return nil
	}
	if verFlag {
		fmt.Fprintln(runner.Stdout, runner.LDFlags.ShowVersion())
		return nil
	}
	param.Path = pflag.Arg(0)
	if param.Path == "" {
		return errFileRequired
	}
	return ctrl.Run(ctx, param) //nolint:wrapcheck
}
