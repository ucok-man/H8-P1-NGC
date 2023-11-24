package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func filelog(fpath string) (loggerfile io.Writer, closefn func()) {
	dirpath := filepath.Dir(fpath)
	err := os.MkdirAll(dirpath, 0777)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(fpath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}

	return file, func() {
		file.Close()
	}
}
