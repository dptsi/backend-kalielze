package main

import "fmt"

func main() {
	/*
		Anonymous func lebih digunakan untuk mendeklarasikan langung
		tanpa membuat deklarasi func terlebih dahulu
	*/

	pengujian_nilai("Karuniawan Putra", 85, func(nilai int) bool {
		if nilai >= 75 {
			return true
		} else {
			return false
		}
	})

	nilai := func(nilai int) bool {
		if nilai >= 75 {
			return true
		} else {
			return false
		}
	}
	pengujian_nilai("Karuniawan Putra", 85, nilai)

	/*
		Kesimpulan

		Anonymous func dapat langsung diletakkan pada parameter atau
		variabel tanpa harus dideklarasikan
	*/

}

type FuncTipe func(int) bool // deklarasi tipe digunakan agar mempermudah penulisan parameter

func pengujian_nilai(nama string, nilai int, cek FuncTipe) {
	fmt.Println("Apakah", nama, "lulus?", cek(nilai))
}
