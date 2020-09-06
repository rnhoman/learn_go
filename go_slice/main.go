package main

import (
	"fmt"
)

func main() {
	/**
	/* append merupakan menambahkan suatu element ke base slice menjadi slice yang baru
	/* append itu bersifat dinamis
	*/

	var gameConsole []string
	gameConsole = append(gameConsole, "PlayStation4")
	gameConsole = append(gameConsole, "Nintendo Switch")
	gameConsole = append(gameConsole, "Xbox One")
	
	fmt.Println(gameConsole)
	
	/**
	/* array yang langsung di isi
	/* slice yang langsung di isi bersifat static
	*/
	gameConsoles := []string{"PlayStation4", "Xbox One", "Nintendo Switch"}
	fmt.Println(gameConsoles)

	/**
	/* mengambil data array dengan foreach
	/* index nya tidak perlu di isi, tapi di ganti dengan tanda "_" garis bawah
	/* karena kita akan mencetak nama array nya
	*/
	for _, console := range gameConsole {
		fmt.Println(console)
	}
}