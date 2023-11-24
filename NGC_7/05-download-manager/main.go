package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	if len(os.Args[1:]) < 2 {
		fmt.Fprintln(os.Stderr, "usage: <destination dir> [url ...]")
		os.Exit(1)
	}

	if (os.Args[2]) == "-help" || (os.Args[2]) == "--help" {
		fmt.Fprintln(os.Stderr, "usage: <destination dir> [url ...]")
		os.Exit(1)
	}

	urls := os.Args[2:]
	for _, url := range urls {
		wg.Add(1)
		go downloadFile(os.Args[1], url, &wg)
	}

	wg.Wait()
	fmt.Println("Done!")
}

// https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_1mb.mp4
