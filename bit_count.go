// Code from wikipedia adopted for Go.
package main

import (
	"fmt"
	"math"
)

const (
	m1  = 0x5555555555555555 // binary: 0101...
	m2  = 0x3333333333333333 // binary: 00110011..
	m4  = 0x0f0f0f0f0f0f0f0f // binary:  4 zeros,  4 ones ...
	m8  = 0x00ff00ff00ff00ff // binary:  8 zeros,  8 ones ...
	m16 = 0x0000ffff0000ffff // binary: 16 zeros, 16 ones ...
	m32 = 0x00000000ffffffff // binary: 32 zeros, 32 ones
	h01 = 0x0101010101010101 // the sum of 256 to the power of 0,1,2,3...
)

func BitCount(x int) int {
	fmt.Printf("%d = %b\n", x, x)
	x = (x & m1) + ((x >> 1) & m1)    // put count of each  2 bits into those  2 bits
	x = (x & m2) + ((x >> 2) & m2)    // put count of each  4 bits into those  4 bits
	x = (x & m4) + ((x >> 4) & m4)    // put count of each  8 bits into those  8 bits
	x = (x & m8) + ((x >> 8) & m8)    // put count of each 16 bits into those 16 bits
	x = (x & m16) + ((x >> 16) & m16) // put count of each 32 bits into those 32 bits
	x = (x & m32) + ((x >> 32) & m32) // put count of each 64 bits into those 64 bits
	return x
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(BitCount(i))
	}
	fmt.Println(BitCount(int(math.Pow(3, 64) + 1)))
}
