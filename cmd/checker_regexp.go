package cmd

import (
	"fmt"

	"github.com/AndreyTsitnarskiy/f2f/internal/checkers"
	"github.com/spf13/cobra"
)

var (
	re          string
	fileToCheck string
)

var checkRegCmd = &cobra.Command{
	Use:   "checkreg",
	Short: "Check file",
	Long:  "Checking a file for a match from a regular expression and line-by-line output",
	Run: func(cmd *cobra.Command, args []string) {
		if re == "" || fileToCheck == "" {
			fmt.Println("❗ Both --reg and --file flags are required")
			cmd.Usage()
			return
		}
		checkers.CheckRegExpInFile(re, fileToCheck)
	},
}

func init() {
	checkRegCmd.Flags().StringVar(&re, "reg", "", "regular expression")
	checkRegCmd.Flags().StringVar(&fileToCheck, "file", "", "filregular expressione")
	rootCmd.AddCommand(checkRegCmd)
}
