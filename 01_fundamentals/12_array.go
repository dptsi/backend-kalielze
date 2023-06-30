package main

import "fmt"

func main() {

	// cara pertama
	var arr_hari [7]string
	arr_hari[0] = "Senin"
	arr_hari[1] = "Selasa"
	arr_hari[2] = "Rabu"
	arr_hari[3] = "Kamis"
	arr_hari[4] = "Jum'at"
	arr_hari[5] = "Sabtu"
	arr_hari[6] = "Minggu"
	fmt.Println(arr_hari)    // menampilkan semua isi array
	fmt.Println(arr_hari[1]) // menampilkan value array index ke 1

	/*
		Value array tidak boleh melebihi kapasitas
		pendeklarasian array.

		Kapasitas array yang sudah dibuat, maka
		tidak boleh ditambah.
	*/

	// cara kedua
	var arr_angka = [10]int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
	}
	fmt.Println(arr_angka)

	/*
		Value data array dapat diubah, sama seperti variabel
	*/

	arr_angka[8] = 6 // merubah value data arr_angka index ke 8 menjadi 6
	fmt.Println(arr_angka)

	// cara ketiga
	arr_peserta := [...]string{
		"anton", "budi", "charlie", "david", "elsa", "fiska",
	}

	/*
		cara ketiga adalah mendeklarasikan array tanpa
		mengetahui jumlah data yang masuk dengan [...]
	*/

	fmt.Println("Jumlah peserta", len(arr_peserta), "orang") // mengetahui jumlah data arr_peserta

}
