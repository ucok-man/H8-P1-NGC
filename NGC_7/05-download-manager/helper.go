package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

// No Error handling in case error exit. and log.
func downloadFile(destinationDir, url string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, "err: url -", url, err)
		return
	}
	defer resp.Body.Close()

	// make 05-download-manager dir as root.
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	root = filepath.Join(root, "NGC_7/05-download-manager/")

	// create dir if necessary
	err = os.MkdirAll(filepath.Join(root, destinationDir), 0777)
	if err != nil {
		fmt.Fprintln(os.Stderr, "err: url - ", url, err)
		return
	}

	// make out target file to be root/destinationDir/targetfile
	outname := filepath.Base(url)
	outpath := filepath.Join(root, destinationDir, outname)

	// get file info to check file is exists
	finfo, err := os.Stat(outpath)
	if err != nil && !os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, "err: url -", url, err)
		return
	}

	// file already exists - add indexing to front of filename.
	if finfo != nil {
		listentry, err := os.ReadDir(destinationDir)
		if err != nil {
			fmt.Fprintln(os.Stderr, "err: url -", url, err)
			return
		}

		var count int
		for _, entry := range listentry {
			if entry.IsDir() {
				continue
			}

			match, err := regexp.MatchString(finfo.Name(), entry.Name())
			if err != nil {
				fmt.Fprintln(os.Stderr, "err: url -", url, err)
				return
			}

			if match {
				count++
			}
		}

		outname = fmt.Sprintf("(%v)_", count) + outname
		outpath = filepath.Join(destinationDir, outname)
	}

	// make out file
	out, err := os.Create(outpath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "err: url -", url, err)
		return
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "err: url -", url, err)
		return
	}
	fmt.Println(outpath)
}
