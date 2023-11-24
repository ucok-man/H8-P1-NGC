package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
)

func showhelp(writer io.Writer) {
	var buff bytes.Buffer
	buff.WriteString("~anagram-cli~\n")
	buff.WriteString("Usage: anagram-cli [-help] | [<word1> <word2>]\n")

	_, err := buff.WriteTo(writer)
	if err != nil {
		panic(err)
	}
}

func filelog(fpath string) (loggerfile io.Writer, closefn func()) {
	dirpath := filepath.Join(filepath.Dir(fpath))
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
