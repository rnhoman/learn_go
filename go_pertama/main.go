package main

import "fmt"
import "pertama/calculation"

func main()  {
	fmt.Println("Halo, belajar golang")

	result := calculation.Add(10, 11)

	fmt.Println(result)
}