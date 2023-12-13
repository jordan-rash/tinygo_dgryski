package main

import (
	"os"
	"syscall"
	"time"
)

func main() {
	println("hello, world", time.Now().UnixNano())
	time.Sleep(1 * time.Second)
	println("hello, world", time.Now().UnixNano())
	println("args")
	for i, s := range os.Args {
		println(i, s)
	}

	println("env")
	for i, s := range syscall.Environ() {
		println(i, s)
	}

	syscall.Exit(1)
}
