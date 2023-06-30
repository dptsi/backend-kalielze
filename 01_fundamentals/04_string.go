package main

import "fmt"

func main() {
	fmt.Println("Helloworld, I starting learn golang program language")

	// untuk mencari panjang string
	fmt.Println(len("Helloworld, I starting learn golang program language"))

	// untuk mencari karakter pada string
	fmt.Println("Helloworld, I starting learn golang program language"[9])

	/*
		Catatan:

		Pada hasil terakhir menunjukkan angka 100, hal ini
		menunjukkan byte karakter yang dipilih.

		Hasil di atas dapat diverifikasi pada tabel ASCII
	*/
}
