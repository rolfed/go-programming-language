package main

import ( 
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	location := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, location)
	} else {
		for _, arg := range files {
			file, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Duplicate error: %v\n", err)
				continue
			} 

			countLines(file, counts, location)
			file.Close()
		}

		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%v\t%s\n", n, location[line], line)
			}
		}
	}

}

func isInFiles(expected string, files []string) bool {
	isInFiles := false;
	for _, actual := range files {
		if expected == actual {
			isInFiles = true
		}
	}

	return isInFiles 
}

func countLines(file *os.File, counts map[string]int, location map[string][]string) {
	input := bufio.NewScanner(file)

	for input.Scan() {
		line := input.Text()
		counts[line]++

		if !isInFiles(file.Name(), location[line]) {
			location[line] = append(location[line], file.Name())
		}
	}
	// NOTE: ignoring potential erros from input.Err()

}
