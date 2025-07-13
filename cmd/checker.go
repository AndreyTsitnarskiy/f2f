package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/AndreyTsitnarskiy/f2f/internal/checkers"
	"github.com/spf13/cobra"
)

var checkInPath string

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check files for Russian symbols",
	Run: func(cmd *cobra.Command, args []string) {

		if checkInPath == "" {
			fmt.Println("Error: missing --in flag")
			return
		}

		info, err := os.Stat(checkInPath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if info.IsDir() {
			err := filepath.Walk(checkInPath, func(path string, info fs.FileInfo, err error) error {
				if !info.IsDir() {
					checkers.CheckFileForRussianChars(path)
				}
				return nil
			})
			if err != nil {
				fmt.Println("Error walking directory:", err)
			}
		} else {
			checkers.CheckFileForRussianChars(checkInPath)
		}
	},
}

func init() {
	checkCmd.Flags().StringVar(&checkInPath, "in", "", "File or directory to check")
	rootCmd.AddCommand(checkCmd)
}
