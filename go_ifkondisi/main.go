package main

import (
	"fmt"
)

func main(){
	// age := 9

	// if (age > 10) {
	// 	fmt.Println("boleh bermain game")
	// } else {
	// 	fmt.Println("kamu belum boleh")
	// }

	score := 49.5

	var grade string

	if (score <= 50) {
		grade = "E"
	}else if (score <= 60) {
		grade = "D"
	}else if (score <= 70) {
		grade = "C"
	}else {
		grade = "NULL"
	}

	fmt.Println(grade)
}