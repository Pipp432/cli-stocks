package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


func Execute() {
	if err := queryCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Zero '%s'\n", err)
		os.Exit(1)
	}
}

func init() {

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(queryCmd)
}
