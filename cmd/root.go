package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var zeroCmd = &cobra.Command{
	Use:   "zero",
	Short: "zero is a cli tool for performing basic mathematical operations",
	Long:  "zero is a cli tool for performing basic mathematical operations - addition, multiplication, division and subtraction.",
	Run: func(cmd *cobra.Command, args []string) {
		println("Hello world")
	},
}

func Execute() {
	if err := zeroCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Zero '%s'\n", err)
		os.Exit(1)
	}
}

func init() {

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(zeroCmd)
}
