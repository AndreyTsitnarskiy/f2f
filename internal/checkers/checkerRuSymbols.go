package checkers

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func CheckFileForRussianChars(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Filed to read %s: %v\n", path, err)
		return
	}

	lines := strings.Split(string(content), "\n")
	re := regexp.MustCompile(`[А-Яа-яЁё]`)

	found := false
	for i, line := range lines {
		if re.MatchString(line) {
			if !found {
				fmt.Printf("📁 %s:\n", path)
				found = true
			}
			fmt.Printf("  Line %d: %s\n", i+1, line)
		}
	}
	if !found {
		fmt.Printf("✅ No Russian characters found in %s\n", path)
	}
}
