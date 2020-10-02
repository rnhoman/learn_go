package main

import (
	"fmt"
)


type BangunDatar interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar int
}

func (persegipanjang PersegiPanjang) HitungLuas() int {
	return persegipanjang.Panjang * persegipanjang.Lebar
}

type Asal struct {
	Angka int
}

func (asal Asal) HitungLuas() int {
	return 1001
}

func main()  {
	bangunDatar := PersegiPanjang{Lebar: 5, Panjang: 6}
	luas := SeberapaLuas(bangunDatar)
	fmt.Println("Luas Persegi Panjang: ", luas)

	persegi := Persegi{Sisi: 4}
	luasPersegi := SeberapaLuas(persegi)
	fmt.Println("Luas Persegi : ", luasPersegi)

	asal := Asal{Angka: 5}
	luas = SeberapaLuas(asal)
	fmt.Println("Luas Asal :", luas)
}

func SeberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}