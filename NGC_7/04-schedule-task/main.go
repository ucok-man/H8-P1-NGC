package main

import (
	"fmt"
	"sync"
	"time"
)

func scheduleTask(wg *sync.WaitGroup, listTaskfn ...func(wg *sync.WaitGroup)) {
	defer wg.Done()
	fmt.Println("Executing schedule task...")

	for _, taskfn := range listTaskfn {
		wg.Add(1)
		go taskfn(wg)
	}
}

func main() {
	var wg sync.WaitGroup

	// func sample task
	task_1 := func(wg *sync.WaitGroup) {
		defer wg.Done()

		fmt.Println("Task One started...")
		time.Sleep(2 * time.Second)
		fmt.Println("Task One done...")
	}
	task_2 := func(wg *sync.WaitGroup) {
		defer wg.Done()

		fmt.Println("Task Two started...")
		time.Sleep(2 * time.Second)
		fmt.Println("Task Two done...")
	}
	task_3 := func(wg *sync.WaitGroup) {
		defer wg.Done()

		fmt.Println("Task Three started...")
		time.Sleep(2 * time.Second)
		fmt.Println("Task Three done...")
	}

	wg.Add(1)
	go scheduleTask(&wg, task_1, task_2, task_3)

	// operation on main
	fmt.Println("Main process start...")
	fmt.Println("Main process done...")

	// wait all process to done.
	wg.Wait()
	fmt.Println("All proccess done!")
}
