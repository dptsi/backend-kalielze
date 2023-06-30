package main

import "fmt"

func main() {
	/*
		Convertion adalah cara untuk mentransformasi
		suatu jenis ke jenis tertentu
	*/

	hello := "Helloworld, I starting learn golang program language"
	fmt.Println(hello)

	/*
		Pada pembahasan 04_string.go, kita telah mengetahui
		proses pengambilan salah satu karakter dalam sebuah string.
		Dimana hasil pengambilan tersebut menhasilkan sebuah byte.

		Untuk mengetahui byte tersebut karakter apa, maka
		perlu dilakukan convertion kedalam bentuk string, sehingga
		kita dapat mengetahui bahwa karakter yang kita inginkan
		adalah sesuai.
	*/

	get_char := hello[9] // mengabil karakter ke sembilan dari awal (0)
	fmt.Println(get_char)

	convert_char := string(get_char)
	fmt.Println(convert_char)

	/*
		Contoh convertion lain bisa untuk angka, yang semula int8,
		menjadi int64. Namum perlu diingat, jika melakukan
		convertion angka harus hati-hati. Karena jika melakukan
		downgrade, maka hasil akan tidak sesuai, hal ini disebabkan
		setiap integer memiliki batasan limit
	*/

	var angka8 int8 = 125
	var angka64 int64 = int64(angka8)
	fmt.Println(angka64)

	var nilai16 int16 = 129
	var nilai8 int8 = int8(nilai16)
	fmt.Println(nilai8)

	/*
		Hasil dari nilai8 tidak dapat terdefinisi dengan baik
	*/
}
