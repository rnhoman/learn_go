package main

import (
	"fmt"
)

// struct adalah modelnya

type Users struct {
	ID int
	firstName string
	lastName string
	email string
	isActive bool
}

func main() {
	// cara pertama
	/** user := Users{}
	user.ID = 1
	user.firstName = "Rohman Nur"
	user.lastName = "Haqiqi"
	user.email = "rohmannurhaqiqi@gmail.com"
	user.isActive = true

	user2 := Users{}
	user2.ID = 2
	user2.firstName = "Herli"
	user2.lastName = "Agustiani"
	user2.email = "herliagustiani@gmail.com"
	user2.isActive = true

	fmt.Println(user)
	fmt.Println(user2)
	*/

	// cara kedua
	user := Users{ID: 1, firstName: "Rohman Nur", lastName: "Haqiqi", email: "rohmannurhaqiqi@gmail.com", isActive: true}

	fmt.Println(user)
}