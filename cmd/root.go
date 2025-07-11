package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/AndreyTsitnarskiy/f2f/internal/converter"
)

func Execute() {
	var inputPath, outputPath, toFormat string

	flag.StringVar(&inputPath, "in", "", "Input file path, example: /home/user/file.json")
	flag.StringVar(&outputPath, "out", "", "Output file path, example: /home/user/file.yaml")
	flag.StringVar(&toFormat, "to", "", "Choice convert format, example: yaml")

	flag.Parse()

	if inputPath == "" || outputPath == "" || toFormat == "" {
		fmt.Println("No some args")
		flag.Usage()
		os.Exit(1)
	}

	ext := filepath.Ext(inputPath)

	switch {
	case ext == ".json" && toFormat == "yaml":
		err := converter.JsonToYaml(inputPath, outputPath)
		checkErr(err)
	case (ext == ".yaml" || ext == ".yml") && toFormat == "json":
		err := converter.YamlToJson(inputPath, outputPath)
		checkErr(err)
	default:
		log.Fatalf("Conversion from %s to %s not supported yet", ext, toFormat)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Conversion successful.")
	}
}
