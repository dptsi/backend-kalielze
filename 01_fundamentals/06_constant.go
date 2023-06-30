package main

import "fmt"

func main() {
	/*
		Constant hampir sama dengan variable.
		Perbedaannya adalah constant tidak mewajibkan
		harus digunakan ketika sudah dibuat.

		Dan constant tidak bisa menggunakan :=, karena
		untuk membedakan variabel dan constan adalah
		pendeklarasian const dan var. Tanda := adalah
		pendeklarasian dari variabel bukan constant
	*/

	const phi = 3.14
	const percent = 100
	fmt.Println(phi)

	/*
		Value constant tidak dapat dirubah seperti
		pada variabel

		Misal kita ingin merubah phi menjadi 10, maka

		phi = 10

		maka deklarasi tersebut akan error
	*/

	// phi = 10

	/*
		Kesimpulan, constant hanya dibuat untuk sekali 
		pendeklarasian dan value akan bersifat konsisten
		atau tidak dapat dirubah. Biasanya digunakan sebagai
		pendeklarasian nilai yang sifatnya tetap seperti phi
		
		Constant dapat juga dideklarasikan secara multiple
	*/

	const (
		rupiah = 1200
		lulus = true
	)
}
