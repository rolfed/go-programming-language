package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var name string
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		name = filename
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Duplicate error: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		// NOTE: line is not empty
		if line != "" {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", name, n, line)
			}
		}
	}
}
