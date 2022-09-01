package main

import (
	"bufio"
	"fmt"
	"os"
)

func Check() {
	
}

func PrintLogo() {
	fmt.Println(`
██╗  ██╗██████╗ ██████╗  ██████╗ ██╗  ██╗██╗   ██╗
╚██╗██╔╝██╔══██╗██╔══██╗██╔═══██╗╚██╗██╔╝╚██╗ ██╔╝
 ╚███╔╝ ██████╔╝██████╔╝██║   ██║ ╚███╔╝  ╚████╔╝ 
 ██╔██╗ ██╔═══╝ ██╔══██╗██║   ██║ ██╔██╗   ╚██╔╝  
██╔╝ ██╗██║     ██║  ██║╚██████╔╝██╔╝ ██╗   ██║   
╚═╝  ╚═╝╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝   ╚═╝   
                                                  
`)
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
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


func main() {
	PrintLogo()
	fmt.Println("hi")
}