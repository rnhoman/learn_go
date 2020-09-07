package main

import (
	"fmt"
)

func main() {
	/**
	/* map[string] = string yg di dalam kurung kurawal adalah key nya
	/* map[string]int = int yg di luar kurung kurawal atau stelah kurung kurawal adalah type datanya
	*/
	var myMap map[string]int
	myMap = map[string]int{}

	/** 
	/* jika ada 2 key yang sama dengan nilai value yang berbeda maka yang di ambil adalah key dan value yg terakhir
	*/
	myMap["Ruby"] = 9
	myMap["Go"] = 9
	myMap["JavaScript"] = 8
	myMap["Go"] = 10

	/*
	/* jika ingin mengambil nilai value dari key yg ada
	/* panggil panggil nama variabel nya, lalu masukan kurung kurawal dan masukan key nya untuk mengetahui hasil dari key tersebut
	/* note : kalau mau mengetahui hasil dari key tersebut pastikan nama key nya sesuai, jika tidak sesuai maka hasil yg ditampilkan akan 0 atau default dari hasil key nya
	*/

	fmt.Println(myMap["Ruby"])
	fmt.Println(myMap)

	myMap2 := map[string]string{
		"Go": "is the best fast", 
		"Ruby": "is the best language",
		"JavaScript": "is in language web",
	}
	fmt.Println(myMap2)

	for index, console := range myMap2 {
		fmt.Println("index : ", index, " value : ", console)
	}

	fmt.Println("==============")

	delete(myMap2, "Go")

	for key, value :=range myMap2 {
		fmt.Println("Key : ", key, " value : ", value)
	}

	fmt.Println("===============")
	value, isAvaibel := myMap2["phyton"]

	fmt.Println(isAvaibel)
	fmt.Println(value)
}