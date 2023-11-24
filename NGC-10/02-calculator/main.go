package main

import (
	"H8-P1-NGC/NGC-10/utils"
	"H8-P1-NGC/common/path"
	"fmt"
	"io"
	"os"
)

func main() {
	closeCpu := utils.CPUProfile(path.NGC10_CALCULATOR, "profile/cpu.pprof")
	defer closeCpu()

	// main code
	logger, closefn := filelog(path.Join(path.NGC10_CALCULATOR, "panic/log.log"))
	defer closefn()
	defer utils.RecoverPanic(logger)

	if err := run(os.Stdin); err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
	// end main code

	closeMem := utils.MEMProfile(path.NGC10_CALCULATOR, "profile/mem.pprof")
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
