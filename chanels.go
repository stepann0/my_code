package main

import (
	"fmt"
)

func greet(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
}

func greet2(c chan int) {
	for i := -6; i > -1; i++ {
		c <- i
	}
}

func main() {
	c := make(chan int)
	go greet(c)
	go greet2(c)

	for ch := range c {
		fmt.Println(ch)
	}

}
