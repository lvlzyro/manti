package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	fmt.Println("this is the start point of manti.")
	fmt.Println("you're using " + os)
}
