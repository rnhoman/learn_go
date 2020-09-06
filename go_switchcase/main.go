package main

import (
	"fmt"
)

func main() {
	number := 10

	switch number {
	case 1:
		fmt.Println("Satu")
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
	default:
		fmt.Println("Default")
	}
}