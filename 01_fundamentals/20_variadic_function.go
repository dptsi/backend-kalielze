package main

import "fmt"

func main() {

	fmt.Println(hitung_total(9, 5, 6, 7, 8, 9, 10))

	fmt.Println(hitung_perkalian("hasil", 7.5, 7.6, 8.9))

	/*
		variadic function dapat diisi dengan slice, dengan
		mengisikan pada parameter fungsi dan ditambahi argumen
		...
	*/

	slc := []float64{
		5.6, 7.8, 9.9, 0.8, 5.5,
	}

	fmt.Println(hitung_perkalian("hasil slice", slc...))

}

/*
	Variadic memungkinkan banyak parameter yang ingin
	dimasukkan pada sebuah fungsi, namun harus yang
	memiliki kesamaan tipe
*/

func hitung_total(nilai ...int64) (r int64) {
	r = 0

	for _, total := range nilai {
		r += total
	}

	return
}

/*
	Apabila ingin menggunakan dua parameter yang
	berbeda, maka parameter variadic harus terletak
	di bagian paling akhir
*/

func hitung_perkalian(str string, nilai ...float64) (s string, r float64) {
	r = 1

	for _, total := range nilai {
		r *= total
	}

	s = str

	return
}
