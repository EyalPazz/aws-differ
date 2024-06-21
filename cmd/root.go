package cmd

import (
    "github.com/spf13/cobra"
)

var (
    // Used for flags.

    rootCmd = &cobra.Command{
        Use:   "aws-differ",
        Short: "aws-differ is a cli tool to check differences between aws resources",
        Long: "aws-differ is a cli tool to check differences between aws resources",
        Run: func(cmd *cobra.Command, args []string) {
        },
    }
)

// Execute executes the root command.
func Execute() error {
    Extend(rootCmd)
    return rootCmd.Execute()
}

func init() {
}


