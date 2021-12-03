package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadByLines() {
	file, err := os.Open(`H:\workspaces\go\src\autoconf\resources\cfg.ini`)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// This is our buffer now
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("read lines:")
	for _, line := range lines {
		fmt.Println(line)
	}
}
