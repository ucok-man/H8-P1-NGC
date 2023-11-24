package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func filelog(fpath string) (loggerfile io.Writer, closefn func()) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dirpath := filepath.Join(cwd, filepath.Dir(fpath))
	err = os.MkdirAll(dirpath, 0777)
	if err != nil {
		log.Fatal(err)
	}

	fpath = filepath.Join(cwd, fpath)
	file, err := os.OpenFile(fpath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatal(err)
	}

	return file, func() {
		file.Close()
	}
}
