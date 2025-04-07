package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newReposUnblockCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "unblock <repo>",
		Short: "Unblock ingestion on a repository.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			repoName := args[0]
			client := NewApiClient(cmd)

			err := client.Repositories().Unblock(repoName)
			exitOnError(cmd, err, "Error unblocking ingestion on repository")

			fmt.Fprintf(cmd.OutOrStdout(), "Successfully remove ingestion block on repo %s\n", repoName)
		},
	}

	return &cmd
}
