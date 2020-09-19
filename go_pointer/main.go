package main

import (
	"fmt"
)

func main() {
	// numberA := 5
	// numberB := &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// *numberB = 10
	// fmt.Println(*numberB)
	// fmt.Println(numberA)

	number := 5
	// fmt.Println("Alamat memory : ", &number)
	fmt.Println("Nilai awal : ", number)

	change(&number, 100)
	// fmt.Println("Alamat memory : ", &number)
	fmt.Println("Nilai akhir : ", number)
	
}

func change(old *int, new int)  {
	*old = new
	// fmt.Println("Alamat memory : ", &old)
	fmt.Println("Di dalam function : ", *old)
}