package main

import (
	"fmt"
)

type Student struct {
	ID int
	Name string
	GPA float32
}

func (student *Student) graduate() {
	student.Name = student.Name + " S.T"
}

func main()  {
	student := Student{1, "Rohman Nur Haqiqi", 3.24}
	fmt.Println(student.Name)

	// graduate(&student)
	student.graduate()
	
	fmt.Println(student.Name)
}