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

func (user Users) display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.firstName,user.lastName,user.email)
}

func displayUserStruct(user Users) string {
	result := fmt.Sprintf("Name : %s %s, Email : %s", user.firstName,user.lastName,user.email)
	return result
}

type Groups struct {
	Name string
	Admin Users
	Members []Users
	isAvailable bool
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
	// user := Users{ID: 1, firstName: "Rohman Nur", lastName: "Haqiqi", email: "rohmannurhaqiqi@gmail.com", isActive: true}

	// cara ketiga
	user := Users{1, "Rohman Nur", "Haqiqi", "rohmannurhaqiqi@gmail.com", true}
	user2 := Users{2, "Linka Ayu", "Ningsih", "linka@gmail.com", true}

	// pemanggilan method user
	result := user.display()
	fmt.Println(result)
	fmt.Println(user2.display())

	// displayUser1 := displayUserStruct(user)
	// displayUser2 := displayUserStruct(user2)
	
	// fmt.Println(displayUser1)
	// fmt.Println(displayUser2)

	// struct lanjutan menggunakan embeded
	// users := []Users{user, user2}

	// group := Groups{"Gamer", user, users, true}

	// displayGroups(group)
}

func displayGroups(group Groups) {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member : %d", len(group.Members))
	fmt.Println("")

	fmt.Println("Users name : ")
	for _, user := range group.Members {
		fmt.Println(user.firstName)
	}
}