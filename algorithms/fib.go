package main

import (
	"fmt"
	"time"
)

// Рекурсия без оптимизаций
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// Рекурсия с мемоизацией
func fibCache(n int, cache map[int]int) int {
	if n <= 1 {
		cache[n] = n
	}
	if a, ok := cache[n]; ok {
		return a
	}
	f := fibCache(n-1, cache) + fibCache(n-2, cache)
	cache[n] = f
	return f
}

// Генератор последовательности фибоначи
func fibGenerator() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	N := 20
	cache := map[int]int{}
	for i := 0; i < N; i++ {
		fmt.Println(fibCache(i, cache))
	}
	
    // or
	// nextFib := fibGenerator()
	// for i := 0; i < N; i++ {
	// 	fmt.Println(nextFib())
	// }
}
