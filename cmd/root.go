package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "f2f",
	Short: "f2f - Format-to-Format converter",
	Long:  "f2f is a CLI tool to convert files between formats like JSON, YAML, and more.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
