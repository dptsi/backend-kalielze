package main

import "fmt"

func main() {

	/*
		Map secara fungsi hampir sama seperti array
		hanya saja, map index dapat didefinisikan
		sendiri

		Misal ingin menggunakan penamaan index dengan
		string. Jika di array index di default dengan
		integer dan dimulai dari 0
	*/

	var arr_hari [7]string = [7]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jum'at",
		"Sabtu",
		"Minggu",
	}
	fmt.Println(arr_hari[0])

	var map_hari map[string]string = map[string]string{
		"hari_pertama": "Senin",
		"hari_kedua":   "Selasa",
		"hari_ketiga":  "Rabu",
		"hari_keempat": "Kamis",
		"hari_kelima":  "Jum'at",
		"hari_keenam":  "Sabtu",
		"hari_ketujuh": "Minggu",
	}
	fmt.Println(map_hari["hari_pertama"])

	arr_angka := [...]int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
	}

	map_angka := map[string]int{
		"satu":     1,
		"dua":      2,
		"tiga":     3,
		"empat":    4,
		"lima":     5,
		"enam":     6,
		"tujuh":    7,
		"delapan":  8,
		"sembilan": 9,
		"sepuluh":  0,
	}

	fmt.Println(arr_angka[5])
	fmt.Println(map_angka["enam"])

	/*
		Keunggulan dari penggunaan map adalah
		dapat menambahkan data
	*/

	fmt.Println("sebelum ditambah", map_angka)
	map_angka["sebelas"] = 11
	fmt.Println("sesudah ditambah",map_angka)

	/*
		Di map juga dapat melakukan penghapusan data
	*/

	fmt.Println("sebelum dihapus", map_angka)
	delete(map_angka, "sepuluh")
	fmt.Println("sesudah dihapus", map_angka)

	map_huruf := make(map[string]rune) 
	map_huruf["A"] = 'A'
	map_huruf["B"] = 'B'
	map_huruf["C"] = 'C'
	map_huruf["D"] = 'D'
	map_huruf["E"] = 'E'
	fmt.Println(map_huruf["C"])
	fmt.Printf("%c\n", map_huruf["C"])
	fmt.Println(string(map_huruf["C"]))
	fmt.Println(len(map_huruf))
	// fmt.Println(cap(map_huruf)) -> map tidak punya capacity

	/*
		Rune adalah alias dari int32, jadi dalam golang 
		tidak ada tipe data char. Untuk memanipulasi
		char, dapat menggunakan rune. Kemudian saat
		pemanggilan data gunakan printf "%c\n" untuk
		menampilkan manipulasi data atau dapat dikonversi
		ke string()
	*/

}
