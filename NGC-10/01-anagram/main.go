package main

import (
	"H8-P1-NGC/NGC-10/utils"
	"H8-P1-NGC/common/path"
	"fmt"
	"log"
)

func main() {
	closeCpu := utils.CPUProfile(path.NGC10_ANAGRAM, "profile/cpu.pprof")
	defer closeCpu()

	/* ---------------------------------------------------------------- */
	/*                             main code                            */
	/* ---------------------------------------------------------------- */
	logger, closefn := filelog(path.Join(path.NGC10_ANAGRAM, "panic/log.log"))
	defer closefn()
	defer utils.RecoverPanic(logger)

	// if len(os.Args[1:]) != 2 {
	// 	showhelp(os.Stderr)
	// 	os.Exit(1)
	// }

	// if strings.ToLower(os.Args[1]) == "-help" {
	// 	showhelp(os.Stderr)
	// 	os.Exit(1)
	// }

	// if err := run(os.Args[1], os.Args[2]); err != nil {
	// 	fmt.Fprint(os.Stderr, err)
	// 	os.Exit(1)
	// }

	for i := 0; i < 1_000_000; i++ {
		err := run("HelloWorldHowdy", "HowdyHelloWorld")
		if err != nil {
			log.Fatal(err)
		}
	}

	/* ------------------------- end main code ------------------------ */

	closeMem := utils.MEMProfile(path.NGC10_ANAGRAM, "profile/mem.pprof")
	defer closeMem()
}

func run(input_1, input_2 string) error {
	if err := validate(input_1, input_2); err != nil {
		return err
	}

	result := checkAnagram(input_1, input_2)
	fmt.Println(result)

	return nil
}
