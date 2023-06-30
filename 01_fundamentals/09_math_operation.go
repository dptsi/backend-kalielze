package main

import (
	"fmt"
	"math"
)

func main() {
	
	d := 28
	const phi = 3.14
	const val = 2

	r := d / val
	fmt.Println("Jari-jari lingkaran", r)

	l := phi * math.Pow(float64(r), 2)
	fmt.Println("Luas lingkaran", l)

	k := phi * float64(val) * float64(r)
	fmt.Println("Keliling lingkaran", k)

	l += float64(r)
	fmt.Println("luas + jari-jari", l)

	l -= k
	fmt.Println("luas - keliling", fmt.Sprintf("%.2f", l))
	
}