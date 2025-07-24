package checkers

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// ANSI color
const (
	colorStart = "\033[33m" // Yellow
	colorEnd   = "\033[0m"
)

func CheckRegExpInFileFromFile(regFile, textFile string) error {
	regBytes, err := os.ReadFile(regFile)
	if err != nil {
		return fmt.Errorf("could not read regex file: %v", err)
	}
	pattern := string(regBytes)

	re, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("invalid regex: %v", err)
	}

	file, err := os.Open(textFile)
	if err != nil {
		return fmt.Errorf("failed to open text file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	found := false

	for scanner.Scan() {
		line := scanner.Text()
		loc := re.FindStringIndex(line)
		if loc != nil {
			found = true
			match := line[loc[0]:loc[1]]
			highlighted := line[:loc[0]] + colorStart + match + colorEnd + line[loc[1]:]
			fmt.Printf("🔍 Line %d: %s\n", lineNum, highlighted)
		}
		lineNum++
	}

	if !found {
		fmt.Println("✅ No matches found.")
	}
	return nil
}
