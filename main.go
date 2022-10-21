package main

import "runtime"

func main() {
	/*
		Setting max runtime process to 8 core
	*/
	runtime.GOMAXPROCS(runtime.NumCPU())

	/*
		This is basic concurrency concept
	*/
	basic()

	/*
		This is wait group concurrency concept
	*/
	group()
}
