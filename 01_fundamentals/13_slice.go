package main

import "fmt"

func main() {

	str_hello := "Helloworld"
	fmt.Println(str_hello)

	arr_angka := [...]int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
	}
	fmt.Println(arr_angka)

	arr_hari := [...]string{
		"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu",
	}
	arr_hari1 := [...]string{
		"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu",
	}
	fmt.Println(arr_hari)

	/*
		Slice hampir sama seperti substring atau sejenis trim, jika
		digunakan pada string
	*/

	slc_hello1 := str_hello[3:]  // data slice adalah semua karakter pada str_hello, dengan slice dimulai dari index 3
	slc_hello2 := str_hello[:3]  // data slice adalah 3 karakter pertama pada str_hello, dengan slice dimulai dari index awal 0
	slc_hello3 := str_hello[1:5] // data slice adalah 5 karakter pertama pada str_hello, dengan slice dimulai dari index awal 1
	fmt.Println(slc_hello1)
	fmt.Println(slc_hello2)
	fmt.Println(slc_hello3)

	/*
		Jika slice pada array, maka slice dapat berfungsi sebagai
		pengubah data juga
	*/

	slc_angka1 := arr_angka[3:8]
	fmt.Println(slc_angka1)
	fmt.Println(len(arr_angka))
	fmt.Println(len(slc_angka1))

	/*
		Khusus untuk array, ada fungsi untuk melihat kapasitas slice.
		Nilai kapasitas slice adalah jumlah data dari keseluruhan
		data yang ada pada arr_angka dimulai dari pemotongan slice.

		Dalam kasus slc_angka1, karena pemotongan slice dimulai dari
		index ke 3, maka hasil kapasitasnya adalah 7. Dimana kapasitas
		arr_angka adalah 10.
	*/

	fmt.Println(cap(slc_angka1))

	/*
		Kesimpulan
		slc_angka1 := arr_angka[3:8]
		adalah memotong arr_angka dari 8 data dimulai dari index 3 dengan
		kapasitas slice adalah 7 dari index 3

		8 data arr_angka -> 1, 2, 3, 4, 5, 6, 7, 8,
		index 3 arr_angka -> 4, 5, 6, 7, 8,
		kapasitas 7 -> 4, 5, 6, 7, 8, 9, 0
	*/

	/*
		Value slice dapat dirubah, namun perubahan value juga akan
		mempegaruhi value array reference
	*/
	slc_angka1[0] = 10
	fmt.Println(slc_angka1)
	fmt.Println(arr_angka)

	/*
		Slice juga memiliki fungsi append, dimana nantinya data pada
		slice dapat ditambahkan. Namun ada aturan, jika penambahannya
		tidak melebih kapasitas slice, maka slice akan tetap. Sebaliknya
		jika dengan append menambah kapasitas slice, maka slice akan
		secara otomatis dibuatkan array baru
	*/

	slc_hari := arr_hari[3:5]
	fmt.Println("Data slc_hari", slc_hari)
	fmt.Println("Length slc_hari", len(slc_hari))
	fmt.Println("Capacity slc_hari", cap(slc_hari))

	slc_append := append(slc_hari, "data")
	fmt.Println("Hasil append slc_hari", slc_append)

	fmt.Println("arr_hari", arr_hari)
	fmt.Println("slc_hari", slc_hari)
	fmt.Println("slc_append", slc_append)

	/*
		Dari contoh slice append di atas, jika kapasitas slice masih ada,
		maka slice masih akan digunakan dan hal itu akan mempengaruhi array
		reference yang digunakan pada slice
	*/

	slc_hari1 := arr_hari1[5:]
	fmt.Println("Data slc_hari1", slc_hari1)
	fmt.Println("Length slc_hari1", len(slc_hari1))
	fmt.Println("Capacity slc_hari1", cap(slc_hari1))

	slc_append1 := append(slc_hari1, "data")
	fmt.Println("Hasil append slc_hari1", slc_append1)

	fmt.Println("arr_hari1", arr_hari1)
	fmt.Println("slc_hari1", slc_hari1)
	fmt.Println("slc_append1", slc_append1)

	/*
		Slice dapat dibuat tanpa terhubung dengan reference
	*/

	mk_slice := make([]string, 4, 7)
	mk_slice[0] = arr_hari1[0]
	mk_slice[1] = arr_hari1[1]
	mk_slice[2] = arr_hari1[2]
	mk_slice[3] = arr_hari1[3]
	// mk_slice[4] = arr_hari1[4]
	// mk_slice[5] = arr_hari1[5]
	// mk_slice[6] = arr_hari1[6]
	mk_slice1 := append(mk_slice, "Jum'at", "Sabtu", "Minggu")
	// mk_slice1[0] = "Kerja"
	fmt.Println(mk_slice)
	fmt.Println(mk_slice1)

	/*
		Jika ingin membuat slice mengambil semua data reference arr_hari1,
		maka bisa menggunakan fungsi copy
	*/

	mk_slice2 := make([]string, len(arr_hari1), cap(arr_hari1))
	copy(mk_slice2, arr_hari1[:])
	fmt.Println(mk_slice2)

	/*
		Pendefinisian array dan slice hampir mirip, perbedaannya array harus
		didefinisikan jumlah array, baik itu secara langsung atau [...]

		Jika kita membuat array tanpa didefinisikan jumlahnya [], maka akan
		dianggap sebagai slice 
	*/

	arr := [...]int{1, 2,3,4,5,}
	slc := []int{1, 2,3,4,5,}
	fmt.Println("ini array", arr)
	fmt.Println("ini slice", slc)

}
