package main

import (
	"fmt"
	"time"
)

func simple(str string) {
	for i := 1; ; i++ {
		fmt.Printf("%s %d\n", str, i)
		time.Sleep(time.Second)
	}
}

func main() {
	go simple("basic")
	fmt.Println("In main function!")
	time.Sleep(time.Second * 2)
	fmt.Println("Leaving. You're just basic!")

	// NOTE: if `main` function exits then it will terminate all
	// the running goroutines
}
