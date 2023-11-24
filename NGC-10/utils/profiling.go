package utils

import (
	"H8-P1-NGC/common/path"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func CPUProfile(dir path.NGC, fpath string) (closefn func()) {
	fpath = path.Join(dir, fpath)

	f, err := os.Create(fpath)
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}

	return func() {
		defer f.Close()
		defer pprof.StopCPUProfile()
	}
}

func MEMProfile(dir path.NGC, fpath string) (closeFn func()) {
	fpath = path.Join(dir, fpath)

	f, err := os.Create(fpath)
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}

	runtime.GC()
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}

	return func() {
		defer f.Close()
	}
}
