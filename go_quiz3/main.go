package main

import (
	"fmt"
)

func main()  {
	/** soal nomor 1
	/* Menghitung nilai rata-rata
	/* score := [8]int[100,80,75,92,70,93,88,67]
	*/
	// scores := [8]int{100,80,75,92,70,93,88,67}

	// var total int

	// for _, score := range scores {
	// 	total = total + score
	// }
	// length := len(scores)
	// average := float64(total) / float64(length)
	// fmt.Println(total)
	// fmt.Println(average)
	
	/** soal nomor 2
	/* Mengambil nilai terbesar dari scores
	*/
	scores := [8]int{100,80,75,92,70,93,88,67}
	var goodScore []int

	for _, score := range scores {
		if score >= 90 {
			goodScore = append(goodScore, score)
		}
	}

	fmt.Println(goodScore)
} 