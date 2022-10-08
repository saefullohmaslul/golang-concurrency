package main

import (
	"fmt"
	"time"
)

func basicConcurrentFunc(message string) {
	for i := 0; i <= 5; i++ {
		/*
			We using time.Sleep to represent heavy logic
		*/
		time.Sleep(100 * time.Millisecond)
		fmt.Println(fmt.Sprintf("basicConcurrentFunc %s, %d", message, i))
	}
}

func basic() {
	/*
		In this function, we call basicConcurrentFunc with 2 method, synchronous and asynchronous
		If we call basicConcurrentFunc only synchronous, read book will execute after get water finish.
		But if we call get water with asynchronous, it will running by concurrency.

		Example the result is:

		basicConcurrentFunc get water, 0
		basicConcurrentFunc read book, 0
		basicConcurrentFunc get water, 1
		basicConcurrentFunc read book, 1
		basicConcurrentFunc get water, 2
		basicConcurrentFunc read book, 2
		basicConcurrentFunc get water, 3
		basicConcurrentFunc read book, 3
		basicConcurrentFunc get water, 4
		basicConcurrentFunc read book, 4
		basicConcurrentFunc get water, 5
		basicConcurrentFunc read book, 5
	*/
	go basicConcurrentFunc("get water")
	basicConcurrentFunc("read book")
}
