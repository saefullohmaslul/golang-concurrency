package main

import (
	"fmt"
	"sync"
	"time"
)

func paralelFunc(wg *sync.WaitGroup, message string) {
	if wg != nil {
		defer wg.Done()
	}

	for i := 0; i <= 5; i++ {
		/*
			We using time.Sleep to represent heavy logic
		*/
		time.Sleep(100 * time.Millisecond)
		fmt.Println(fmt.Sprintf("paralelFunc %s, %d", message, i))
	}
}

func group() {
	// We declare wg to use wait group function
	var wg sync.WaitGroup

	// set total process. in this case bcs we will running function one and function two, we set paralel process to 2
	paralelProcess := 2
	wg.Add(paralelProcess)

	/*
		We running function one and function two using paralel process
	*/
	go paralelFunc(&wg, "function one")
	go paralelFunc(&wg, "function two")

	/*
		and wait them until finish
	*/
	wg.Wait()

	/*
		then, function three will execute after finished function one and function two

		Example the result is:

		paralelFunc function one, 0
		paralelFunc function two, 0
		paralelFunc function one, 1
		paralelFunc function two, 1
		paralelFunc function two, 2
		paralelFunc function one, 2
		paralelFunc function one, 3
		paralelFunc function two, 3
		paralelFunc function two, 4
		paralelFunc function one, 4
		paralelFunc function one, 5
		paralelFunc function two, 5
		paralelFunc function three, 0
		paralelFunc function three, 1
		paralelFunc function three, 2
		paralelFunc function three, 3
		paralelFunc function three, 4
		paralelFunc function three, 5
	*/
	paralelFunc(nil, "function three")
}
