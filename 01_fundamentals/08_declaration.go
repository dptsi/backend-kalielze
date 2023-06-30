package main

import "fmt"

func main() {
	/*
		Declaration digunakan untuk mendeklarasikan
		type data. Fungsinya untuk mengkhususkan
		penggunaan type data
	*/

	type NIP string
	var id_pegawai NIP = "1234567" // penggunaan NIP mendefinisikan, bahwa variabel id_pegawai adalah String
	fmt.Println(id_pegawai)

	var nama NIP = "Karuniawan Putra" // type data NIP otomatis akan merepresentasikan variabel tersebut adalah string
	fmt.Println(nama)
}