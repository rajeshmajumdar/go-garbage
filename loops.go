package main

import "fmt"

func main() {
		for i := 100000; i < 100100; i++ {
				fmt.Printf("%d \t %b \t %x", i, i, i)
		}
}
