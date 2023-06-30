package main

import "fmt"

func main() {

	var nama string = "Karuniawan Putra"
	fmt.Println(nama)

	/*
		Catatan:

		Pada cara pertama pendeklarasian variabel jelas,
		dengan menyebutkan jenis variabel string.

		Variabel yang sudah dibuat, tidak dapat dibuat
		kembali pada tahap berikutnya dalam satu file.

		Contoh:
		var nama = "El Rafif Zavier Putra" -> error

		karena variabel nama sudah dibuat sebelumnya,
		lain halnya jika

		nama = "El Rafif Zavier Putra"

		maka variabel nama tidak error, namun value
		dari variabel nama akan berubah menjadi
		"El Rafif Zavier Putra"
	*/

	nama = "El Rafif Zavier Putra"
	fmt.Println(nama)

	/*
		Variabel yang sudah dibuat, harus benar-benar
		digunakan, sehingga ketika proses compiler
		sistem golang akan mengecek, apakah ada variabel
		yang tidak digunakan. Jika ada, maka proses
		compiler akan menghasilkan error.

		Contoh:
		var alias string

		jika variabel alias tidak dipanggil/digunakan,
		misal pada fmt, maka akan error
	*/

	// var alias string

	/*
		Variabel dapat dibuat tanpa mendefinisikan jenis
		variabel. Bahasa GO dapat secara otomatis mendeteksi
		jenis variabel berdasarkan value variabel.

		Namun, tetap nantinya variabel akan secara default
		mengikuti jenis variabel yang telah didevinisikan.
	*/

	var kota = "Surabaya"
	/*
		jika dihover pada kota, akan menunjukkan, bahwa
		variable kota adalah string.
	*/
	fmt.Println(kota)

	/*
		Cara kedua di atas memiliki kelemahan, karena jika
		value dideklarasikan setelahnya akan error

		Contoh:
		var kota
		kota = "Surabaya" -> error

		var kota string
		kota = "Surabaya" -> execute

		Cara selanjutnya adalah membuat variabel
		tanpa menuliskan var, dengan cara :=
	*/

	alamat := "Ngagel Dadi 3 No 6 Surabaya"
	kantor := "ITS Surabaya Kampus Sukolilo"
	fmt.Println(alamat)
	fmt.Println(kantor)

	/*
		Namun cara := tidak bisa digunakan ketika kita
		ingin mendeklarasikan variabel khusus. Misal

		angka64 int64 := 12135 -> error

		jika ingin mendeklarasikan variable dengan 
		deklarasi khusus, maka tetap menggunakan var

		Cara terakhir adalah mendefinisikan variabel
		secara multiple
	*/

	var (
		umur               = 30
		tanggal_lahir      = "16 Nopember 1992"
		menikah       bool = true
	)
	fmt.Println(umur)
	fmt.Println(tanggal_lahir)
	fmt.Println(menikah)

}
