package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(fmt.Sprintf("data/%s", path))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func AppendFile(FileName string, Content string) {
	File, err := os.OpenFile(fmt.Sprintf("data/%s", FileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}

	_, err = File.WriteString(Content + "\n")
	if err != nil {
		return
	}

	File.Close()
}