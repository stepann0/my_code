package main

import "fmt"

func main() {
	mem := [30000]byte{}
	p := 0
	
    mem[p] += 8
    for mem[p] != 0 {
        p++
        mem[p] += 4
        for mem[p] != 0 {
            p++
            mem[p] += 2
            p++
            mem[p] += 3
            p++
            mem[p] += 3
            p++
            mem[p]++
            p -= 4
            mem[p]--
        }
        p++
        mem[p]++
        p++
        mem[p]++
        p++
        mem[p]--
        p += 2
        mem[p]++
        for mem[p] != 0 {
            p--
        }
        p--
        mem[p]--
    }
    p += 2
    fmt.Printf("%c", mem[p])
    p++
    mem[p] -= 3
    fmt.Printf("%c", mem[p])
    mem[p] += 7
    fmt.Printf("%c", mem[p])
    fmt.Printf("%c", mem[p])
    mem[p] += 3
    fmt.Printf("%c", mem[p])
    p += 2
    fmt.Printf("%c", mem[p])
    p--
    mem[p]--
    fmt.Printf("%c", mem[p])
    p--
    fmt.Printf("%c", mem[p])
    mem[p] += 3
    fmt.Printf("%c", mem[p])
    mem[p] -= 6
    fmt.Printf("%c", mem[p])
    mem[p] -= 8
    fmt.Printf("%c", mem[p])
    p += 2
    mem[p]++
    fmt.Printf("%c", mem[p])
    p++
    mem[p] += 2
    fmt.Printf("%c", mem[p])
}
