package checkers

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// ANSI color
const (
	highlightStart = "\033[1;31m" // ярко-красный
	highlightEnd   = "\033[0m"
)

func CheckRegExpInFile(re, filePath string) {

	reg, err := regexp.Compile(re)
	if err != nil {
		fmt.Printf("❌ Invalid regular expression: %v\n", err)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("❌ Failed to open file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	matchesFound := false

	for scanner.Scan() {
		line := scanner.Text()
		if reg.MatchString(line) {
			matchesFound = true

			highlighted := reg.ReplaceAllStringFunc(line, func(m string) string {
				return highlightStart + m + highlightEnd
			})
			fmt.Printf("🔍 Match at line %d: %s\n", lineNum, highlighted)
		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("❌ Error reading file: %v\n", err)
	}

	if !matchesFound {
		fmt.Println("✅ No matches found.")
	}
}
