package main

import "fmt"

func main() {

	/*
		cara pertama
	*/
	pengujian_nama("Karuniawan Putra", pengecekan_nama)

	cek_nama := pengecekan_nama
	pengujian_nama("Karuniawan Putra", cek_nama)

	/*
		cara kedua
	*/

	pengujian_nilai("Karuniawan Putra", 80, pengecekan_nilai)
}

func pengujian_nama(nama string, cek func(string) bool) {
	cek_nama := cek(nama)
	fmt.Println("Apakah nama", nama, "terlalu panjang?", cek_nama)
}

func pengecekan_nama(nama string) bool {
	if len(nama) > 20 {
		return false
	} else {
		return true
	}
}

// ============================================================================================== //

type FuncTipe func(int) bool // deklarasi tipe digunakan agar mempermudah penulisan parameter

func pengujian_nilai(nama string, nilai int, cek FuncTipe) {
	fmt.Println("Apakah", nama, "lulus?", cek(nilai))
}

func pengecekan_nilai(nilai int) bool {
	if nilai >= 75 {
		return true
	} else {
		return false
	}
}
