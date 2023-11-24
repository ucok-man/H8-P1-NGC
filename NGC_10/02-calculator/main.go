package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/ucok-man/H8-NGC/PHASE_1/NGC_10/utils"
)

func main() {
	closeCpu := utils.CPUProfile(filepath.Join(utils.GetRootPath(), "profile/cpu.pprof"))
	defer closeCpu()

	// main code
	logger, closefn := filelog(filepath.Join(utils.GetRootPath(), "panic/log.log"))
	defer closefn()
	defer utils.RecoverPanic(logger)

	if err := run(os.Stdin); err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
	// end main code

	closeMem := utils.MEMProfile(filepath.Join(utils.GetRootPath(), "profile/mem.pprof"))
	defer closeMem()
}

func run(reader io.Reader) error {
	var calcApp Calculator

	// show menu
	fmt.Println(calcApp.Welcome())

	// prompt operation
	calcApp.operation = calcApp.promptOperation(reader)

	// prompt number
	calcApp.val_a = calcApp.promptNumber(reader, "input num 1: ")
	calcApp.val_b = calcApp.promptNumber(reader, "input num 2: ")

	// check division by zero
	if calcApp.operation == div && calcApp.val_b == 0 {
		return errDivisionZero("cannot compute the result")
	}

	// calculate
	result := calcApp.calculate(calcApp.operation, calcApp.val_a, calcApp.val_b)
	fmt.Println("RESULT:", result)

	return nil
}
