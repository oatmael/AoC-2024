package aocutils

import (
	"bufio"
	"iter"
	"log"
	"os"
)

func ReadInput(path string) iter.Seq[string] {
	return func(yield func(string) bool) {
		file, err := os.Open(path)
		if (err != nil) {
			log.Fatal(err);
		}
		
		scanner := bufio.NewScanner(file);

		defer file.Close()
		
		for (scanner.Scan()) {
			if !yield(scanner.Text()) {
				return
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err);
		}
	}
}