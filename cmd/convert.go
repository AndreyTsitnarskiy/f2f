package cmd

import (
	"fmt"
	"os"

	"github.com/AndreyTsitnarskiy/f2f/internal/converter"
	"github.com/spf13/cobra"
)

var (
	inPath   string
	outPath  string
	toFormat string
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert file between formats",
	Run: func(cmd *cobra.Command, args []string) {
		if inPath == "" || outPath == "" || toFormat == "" {
			fmt.Println("Error: missing required flags --in, --out, or --to")
			cmd.Usage()
			return
		}

		err := converter.ConvertFile(inPath, outPath, toFormat)
		if err != nil {
			fmt.Println("Conversion failed:", err)
			os.Exit(1)
		}
		fmt.Println("✅ Conversion successful!")
	},
}

func init() {
	convertCmd.Flags().StringVar(&inPath, "in", "", "Path to input file")
	convertCmd.Flags().StringVar(&outPath, "out", "", "Path to output file")
	convertCmd.Flags().StringVar(&toFormat, "to", "", "Target format: json or yaml")

	rootCmd.AddCommand(convertCmd)
}
