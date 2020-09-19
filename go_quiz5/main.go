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

func (group Groups) DisplayGroup() {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member : %d", len(group.Members))
	fmt.Println("")

	fmt.Println("Users name : ")
	for _, user := range group.Members {
		fmt.Println(user.firstName)
	}
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

type Groups struct {
	Name string
	Admin Users
	Members []Users
	isAvailable bool
}

func main() {
	// cara ketiga
	user := Users{1, "Rohman Nur", "Haqiqi", "rohmannurhaqiqi@gmail.com", true}
	user2 := Users{2, "Linka Ayu", "Ningsih", "linka@gmail.com", true}

	// pemanggilan method user
	result := user.display()
	fmt.Println(result)
	fmt.Println(user2.display())

	// struct lanjutan menggunakan embeded
	users := []Users{user, user2}

	group := Groups{"Gamer", user, users, true}

	group.DisplayGroup()
	// displayGroups(group)
}