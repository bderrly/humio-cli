package main

import (
	"github.com/spf13/cobra"
)

func newUsersRotateApiTokenCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "rotate-api-token <user-id>",
		Short: "Rotate and retrieve a user's API token [Root Only]",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			userID := args[0]

			client := NewApiClient(cmd)
			newToken, apiErr := client.Users().RotateToken(userID)
			exitOnError(cmd, apiErr, "Error updating user")

			cmd.Printf("New API Token: %s\n", newToken)
		},
	}

	return &cmd
}
