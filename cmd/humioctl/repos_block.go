package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func newReposBlockCmd() *cobra.Command {
	var durationFlag time.Duration

	cmd := cobra.Command{
		Use:   "block <repo> --duration=<duration>",
		Short: "Block ingestion on a repository.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			repoName := args[0]
			client := NewApiClient(cmd)

			if durationFlag.Seconds() == 0 {
				exitOnError(cmd.Parent(), fmt.Errorf("must supply a non-zero duration"), "No duration specified")
			}

			err := client.Repositories().Block(repoName, durationFlag)
			exitOnError(cmd, err, "Error blocking ingestion on repository")

			fmt.Fprintf(cmd.OutOrStdout(), "Successfully blocked ingestion on repo %s for %s\n", repoName, durationFlag.String())
		},
	}

	cmd.Flags().DurationVarP(&durationFlag, "duration", "d", 0*time.Second, "how long to block ingest")

	return &cmd
}
