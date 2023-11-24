package main

import (
	"embed"
	"log"
)

func getSampleFile(fs embed.FS) []string {
	entrys, err := fs.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	var samplefile []string
	for _, entry := range entrys {
		samplefile = append(samplefile, entry.Name())
	}

	return samplefile
}
