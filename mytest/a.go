package main

import "fmt"

func main() {

	k := make([]byte, 10, 20)
	for i := 0; i < 10; i++ {
		k[i] = byte(i)
	}

	copy(k, k[5:])
	fmt.Printf("haha :  %v, %v, \n", len(k), cap(k))

	fmt.Printf("haha %v  \n", k)
}
