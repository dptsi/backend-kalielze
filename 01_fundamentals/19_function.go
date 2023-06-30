package main

import (
	"fmt"
	"math"
)

func main() {
	helloworld()

	// r, l := luas_lingkaran(67.8)
	// fmt.Println("Jari-jari lingkaran", fmt.Sprintf("%.2f", r))
	// fmt.Println("Luas lingkaran", fmt.Sprintf("%.2f", l))

	_, l := luas_lingkaran(67.8)
	// fmt.Println("Jari-jari lingkaran", fmt.Sprintf("%.2f", r))
	fmt.Println("Luas lingkaran", fmt.Sprintf("%.2f", l))

	_, l1 := luas_lingkaran1(61.8)
	fmt.Println("Luas lingkaran1", fmt.Sprintf("%.2f", l1))
}

func helloworld() {
	fmt.Println("Menghitung luas lingkaran")
}

/*
	Membuat function paramenter dan return
*/
// cara pertama
func luas_lingkaran(d float64) (float64, float64) {
	const phi = 3.14
	r := d / 2
	// fmt.Println("Jari-jari lingkaran", r)

	l := phi * math.Pow(float64(r), 2)
	// fmt.Println("Luas lingkaran", l)

	return r, l
}

// cara kedua
func luas_lingkaran1(d float64) (r float64, l float64) {
	const phi = 3.14
	r = d / 2
	// fmt.Println("Jari-jari lingkaran", r)

	l = phi * math.Pow(float64(r), 2)
	// fmt.Println("Luas lingkaran", l)

	return
	/*
		return tidak perlu deklarasikan variabel,
		karena variabel sudah dideklarasikan pada
		function

		cara kedua ini mempermudah ketika ada banyak
		return variable yang akan dikembalikan
	*/
}
