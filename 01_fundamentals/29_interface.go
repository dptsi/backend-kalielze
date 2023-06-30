package main

import (
	"fmt"
	"math"
)

// Definisikan interface area
type Area interface {
	Luas() float64
}

// Struct Lingkaran yang mengimplementasikan interface area
type Lingkaran struct {
	jari_jari float64
}

// Struct Persegi yang mengimplementasikan interface area
type Persegi struct {
	panjang float64
	lebar   float64
}

func main() {
	lingkaran := Lingkaran{jari_jari: 5}
	persegi := Persegi{panjang: 3, lebar: 4}

	PrintArea(lingkaran, "lingkaran")
	PrintArea(persegi, "persegi")
}

// Fungsi yang menggunakan parameter berupa interface area
func PrintArea(a Area, j string) {
	fmt.Println("Area luas", j, ":", a.Luas())
}

func (luas Lingkaran) Luas() float64 {
	const phi = 3.14
	return phi * math.Pow(luas.jari_jari, 2)
}

func (luas Persegi) Luas() float64 {
	return luas.panjang * luas.lebar
}
