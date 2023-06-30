package main

import "fmt"

type AttributPegawai struct {
	nama, nip string
	usia      int
}

func main() {

	// cara pertama
	var supervisor AttributPegawai
	supervisor.usia = 30
	supervisor.nama = "Karuniawan Putra"
	supervisor.nip = "19922022410004"

	fmt.Println(supervisor)
	// fmt.Println(supervisor.nama)
	// fmt.Println(supervisor.nip)
	// fmt.Println(supervisor.usia)


	// cara kedua
	manager := AttributPegawai {
		nip: "1992202241004",
		nama: "Karuniawan Putra",
		usia: 30,
	}

	fmt.Println(manager)

	// cara ketiga
	ceo := AttributPegawai {
		"Karuniawan Putra", "1992202241004", 30,
	}

	/*
		Cara ketiga tidak memiliki fleksibilitas jika
		struktur struct berubah posisi atau terjadi
		penambahan/pegurangan data struct

		contoh adalah kode di bawah akan error

		ceo := AttributPegawai {
			"Karuniawan Putra", 30, "1992202241004",
		}

		karena nip harusnya terletak sebelum usia
	*/

	fmt.Println(ceo)

}
