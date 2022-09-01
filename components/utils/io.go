package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(fmt.Sprintf("data/%s", path))
	if HandleError(err) {
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
	if HandleError(err) {
		return
	}

	_, err = File.WriteString(Content + "\n")
	if HandleError(err) {
		return
	}

	File.Close()
}

func RemoveDuplicateStr(strSlice []string) []string {
    allKeys := make(map[string]bool)
    list := []string{}
    for _, item := range strSlice {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}