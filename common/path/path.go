package path

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
)

type NGC string

// if change the directory this must be change also.
const (
	NGC07_DOWNLOAD_MANAGER NGC = "NGC-07/05-download-manager"
	NGC10_ANAGRAM          NGC = "NGC-10/01-anagram"
	NGC10_CALCULATOR       NGC = "NGC-10/02-calculator"
)

func getwd() string {
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return root
}

func Join(dir NGC, targetpath string) string {
	joinedpath := filepath.Join(getwd(), string(dir), targetpath)
	parts := strings.Split(joinedpath, string(filepath.Separator))
	// fmt.Println(parts)

	// cleaning the same path name.
	temp := map[string]int{}
	for i, path := range parts {
		temp[path] = i
	}

	// fmt.Println(temp)
	var idxs []int
	for key := range temp {
		idxs = append(idxs, temp[key])
	}
	slices.Sort(idxs)

	var cleanpaths []string
	for _, idx := range idxs {
		cleanpaths = append(cleanpaths, parts[idx])
	}

	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		res := filepath.Join(cleanpaths...)
		res = "/" + res
		return res
	}

	return filepath.Join(cleanpaths...)
}
