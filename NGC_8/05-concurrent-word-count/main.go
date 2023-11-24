package main

import (
	"embed"
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ucok-man/H8-NGC/PHASE_1/NGC_8/05-concurrent-word-count/sample"
)

func countWord(fname string, fs embed.FS) (int, error) {
	content, err := fs.ReadFile(fname)
	if err != nil {
		return -1, fmt.Errorf("error %s: %v", fname, err)
	}

	contents := strings.Split(string(content), " ")
	return len(contents), nil
}

func main() {
	// sample input filename
	files := getSampleFile(sample.Samplefile)

	var wg sync.WaitGroup
	wg.Add(len(files))

	wordcount := make(map[string]int)
	for _, filename := range files {
		go func(fname string) {
			defer wg.Done()

			count, err := countWord(fname, sample.Samplefile)
			if err != nil {
				fmt.Println(err)
				return
			}

			wordcount[filepath.Base(fname)] = count
		}(filename)
	}

	wg.Wait()

	// display
	for fname, count := range wordcount {
		fmt.Printf("%s: %d words\n", fname, count)
	}
}
