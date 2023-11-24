package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func RecoverPanic(logger io.Writer) {
	if err := recover(); err != nil {
		// log to file actual error
		logmsg := fmt.Sprintf("PANIC %v: %v\n", time.Now().Format(time.DateOnly), err)
		_, err := logger.Write([]byte(logmsg))
		if err != nil {
			log.Fatal(err)
		}

		// print msg to user
		fmt.Fprintf(os.Stderr, "Sorry we have problem in our application!\n")
		os.Exit(1)
		return
	}
}
