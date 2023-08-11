package main

import (
	"fmt"
)

func Rombo(n int) {
	if n%2 == 0 {
		println("La linea central debe tener una cantidad impar de asteriscos")
	} else {
		for i := 0; i < n; i++ {
			for j := 0; j < n-i-1; j++ {
				fmt.Print(" ")
			}
			for j := 0; j < 2*i+1; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}

		for i := n - 2; i >= 0; i-- {
			for j := 0; j < n-i-1; j++ {
				fmt.Print(" ")
			}
			for j := 0; j < 2*i+1; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}
}

func main() {
	Rombo(9)
}
