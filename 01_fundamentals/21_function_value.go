package main

import (
	"fmt"
	"math"
)

func main() {

	luas := luas_lingkaran // variabel luas secara otomatis menjadi fungsi, seperti luas_lingkaran
	fmt.Println(luas(54.5))
	fmt.Println(luas_lingkaran(54.5))

}

func luas_lingkaran(d float64) (l float64) {
	const phi = 3.14
	r := d / 2
	l = phi * math.Pow(float64(r), 2)
	return
}
