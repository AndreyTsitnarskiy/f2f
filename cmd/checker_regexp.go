package cmd

import (
	"fmt"
	"os"

	"github.com/AndreyTsitnarskiy/f2f/internal/checkers"
	"github.com/spf13/cobra"
)

var (
	re          string
	fileToCheck string
)

var checkRegCmd = &cobra.Command{
	Use:   "checkreg",
	Short: "Check file with regex (wrap your regex in quotes!)",
	Long: `Check a text file line-by-line using a regular expression.

Example:
  f2f checkreg --reg "(?i:select\\s+.+from)" --file "input.txt"`,
	Run: func(cmd *cobra.Command, args []string) {
		if re == "" || fileToCheck == "" {
			fmt.Println("❌ Please provide both --reg and --file")
			return
		}

		// Создаем временный файл
		tmpFile, err := os.CreateTemp("", "regex_*.txt")
		if err != nil {
			fmt.Printf("❌ Failed to create temp file: %v\n", err)
			return
		}
		defer os.Remove(tmpFile.Name()) // Удалим файл после выполнения

		_, err = tmpFile.WriteString(re)
		if err != nil {
			fmt.Printf("❌ Failed to write to temp file: %v\n", err)
			return
		}
		tmpFile.Close()

		// Используем функцию из пакета checkers
		err = checkers.CheckRegExpInFileFromFile(tmpFile.Name(), fileToCheck)
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
		}
	},
}

func init() {
	checkRegCmd.Flags().StringVar(&re, "reg", "", "Regular expression (wrap in quotes!)")
	checkRegCmd.Flags().StringVar(&fileToCheck, "file", "", "File to check")
	rootCmd.AddCommand(checkRegCmd)
}
