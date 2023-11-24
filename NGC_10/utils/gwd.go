package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetRootPath() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(filepath.Join(wd, "NGC_10/01-anagram/"))
	return filepath.Join(wd, "NGC_10/01-anagram/")
}
