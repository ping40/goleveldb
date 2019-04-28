package main

import "fmt"

func main() {

	var x int

	defer func() {
		if x := kkrecover(); x > 0 {
			fmt.Println("123")
		}

		fmt.Printf("x2: %d \n", x)
		if x != 0 {
			fmt.Println("1234")
		}
		fmt.Println("123 5")
	}()
	fmt.Println(22)
}

func kkrecover() int {
	return 8
}
