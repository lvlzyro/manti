package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	fmt.Println("this is the start point of manti.")
	fmt.Println("you're using " + os)
	switch os {
	case "windows":
		fmt.Println("you'll need .exe")
	case "linux":
		fmt.Println("you'll need .deb or .rpm")
	case "darwin":
		fmt.Println("you'll need .app")
	}
}
