package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/suzuki-shunsuke/yaml2json/pkg/cli"
)

var (
	version = ""
	commit  = "" //nolint:gochecknoglobals
	date    = "" //nolint:gochecknoglobals
)

func main() {
	if err := core(); err != nil {
		log.Fatal(err)
	}
}

func core() error {
	runner := cli.Runner{
		Stdout: os.Stdout,
		LDFlags: &cli.LDFlags{
			Version: version,
			Commit:  commit,
			Date:    date,
		},
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if err := runner.Run(ctx); err != nil {
		return err //nolint:wrapcheck
	}
	return nil
}
