package main

import "bitbucket.org/zombiezen/yaml"

import "io"
import "bufio"
import "fmt"
import "os"

func main() {
	file, err := os.Open("t.yaml")
	if err != nil {
		fmt.Println("Can't read t.yaml")
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file))
	
	parser := new yaml.Parser
	var lines []string
	scanner := yaml.NewScanner(bufio.NewReader(file))
	for {
		tok, err := scanner.Scan()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("scan error: ", err, tok)
			break
		}
		lines = append(lines, tok.Value)
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}
