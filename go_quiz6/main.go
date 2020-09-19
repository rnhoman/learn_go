package main

import (
	"fmt"
)

type Gamer struct{
	Name string
	Games []string
}

func (gemer *Gamer)AddGame(game string) {
	gemer.Games = append(gemer.Games, game)
}

func main() {
	gamer := Gamer{Name: "Zelda"}

	// menambahkan gamernya
	gamer.AddGame("Mario")
	gamer.AddGame("Fifa 2020")
	gamer.AddGame("Soccer 2020")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}