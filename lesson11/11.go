package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program finds the sum of digits of N.")
	n := 321
	for ; n != 0; n /= 10 {
		fmt.Println(n)
		s := n % 10
		fmt.Println(s)

	}

}
