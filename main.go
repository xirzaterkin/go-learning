package main

import "fmt"

func main() {
	msg := "hello world!"
	fmt.Println(msg)

	for i := 1; i < 5; i++ {
		fmt.Println("Day1 -Step %d\n", i)
	}
}
