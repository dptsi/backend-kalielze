package main

import (
	"fmt"
	"math"
)

func main() {
	const phi float64 = 3.14
	d := 6.7
	fmt.Println("diameter", d)

	r := d / 2
	fmt.Println("jari-jari", r)

	l := phi * math.Pow(float64(r), 2) // math.Pow() unutk pangkat
	fmt.Println("luas", l)

	round := math.Round(l) // pembulatan otomatis tergantung yang paling terdekat
	floor := math.Floor(l) // pembulatan ke bawah
	ceil := math.Ceil(l)   // pembulatan ke atas

	fmt.Println("round", round)
	fmt.Println("floor", floor)
	fmt.Println("ceil", ceil)

	/*
		masih banyak documentation mengenai package math

		http://golang.org/pkg/math/
	*/
}
