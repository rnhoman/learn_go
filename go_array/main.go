package main

import (
	"fmt"
)

func main() {
	// var language [5]string
	// language[0] = "Go"
	// language[1] = "Ruby"
	// language[2] = "Javascript"
	// language[3] = "C#"
	// language[4] = "Phython"

	/**
	/* array [5] = untuk menentukan jumlah array nya
	/* array [...] = untuk menentukan jumlah array yang ada pada jumlah string nya
	*/
	language := [...]string{
		"Go", 
		"Ruby", 
		"JavaScript",
		"C#",
		"Phyton",
		"Kotlin",
	}

	// fmt.Println(language)

	// length := len(language)
	// fmt.Println(length)

	/**
	/* mengambil data array tertentu dengan menggunakan looping
	*/

	for index, lang := range language {
		fmt.Println("index : ", index, " language : ", lang)
	}
}