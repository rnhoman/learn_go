package main

import (
	"fmt"
)

func main() {
	// sentence := printResult("Saya sedang")
	// fmt.Println(sentence)
	// result := add(10, 20)
	// fmt.Println(result)

	luas, keliling := calculate(10, 20)
	fmt.Println(luas)
	fmt.Println(keliling)
}

func calculate(panjang int, lebar int) (int, int)  {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	return luas, keliling
}
func printResult(sentence string)string {
	newSentence := sentence + " belajar golang"
	return newSentence
}

func add(number int, numberTwo int)int{
	result := number + numberTwo
	return result
}

// 1. input
// 2. proses
// 3. output