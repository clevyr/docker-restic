package main

import (
	"errors"
	"log/slog"
	"os"
	"os/exec"

	"github.com/gabe565/docker-restic/cmd/dumpdb/mariadb"
	"github.com/gabe565/docker-restic/cmd/dumpdb/mongodb"
	"github.com/gabe565/docker-restic/cmd/dumpdb/postgres"
	"github.com/gabe565/docker-restic/cmd/dumpdb/sqlite"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dumpdb",
		Short: "Database utilities",
	}

	cmd.AddCommand(
		mariadb.New(),
		mongodb.New(),
		postgres.New(),
		sqlite.New(),
	)

	return cmd
}

func main() {
	if err := New().Execute(); err != nil {
		slog.Error(err.Error())

		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			os.Exit(exitErr.ExitCode())
		}
		os.Exit(1)
	}
}
