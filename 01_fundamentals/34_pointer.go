package main

import "fmt"

type Person struct {
	nama, tanggal string
	usia         int
}

func main() {

	// var person1 Person = Person{
	// 	nama:   "Karuniawan Putra",
	// 	alamat: "Ngagel Dadi 3 No 6 Surabaya",
	// 	usia:   30,
	// }

	person1 := Person{
		nama:   "Karuniawan Putra",
		tanggal: "16 Nov 1992",
		usia:   30,
	}

	// var person2 *Person = &person1

	person2 := person1
	person2.nama = "Lita Mandasari"

	/*
		Person2 tidak mereferensi ke person1, value pada person1 hanya
		di duplikasi pada person2. Sehingga ketika terjadi perubahan
		data pada person2 tidak mempengaruhi data pada person1

		Jika ingin mereference pada person1, maka variabel baru atau 
		person1 baru harus menggunakan & sebelum penamaan variabel.
		Sehingga ketika terjadi perubahan data pada person baru akan
		merubah data person1
	*/

	person3 := &person1
	person3.nama = "Zeanissa Aqila Alodie"

	// melakukan perubahan semua data person4
	person4 := &person1
	// person4 = &Person{
	// 	nama:   "El Rafif Zavier Putra",
	// 	tanggal: "23 Mar 2019",
	// 	usia:   5,
	// }

	/*
		Cara diatas pada person4 = &Person{...}, tidak dapat merubah data
		pada person meskipun person4 adalah reference dari person.

		Bagaimana cara merubah semua data person mengikuti data person4
		Hal tersebut dapat dilakukan dengan membuat pointer *, sehingga 
		nantinya data pada person, akan mengikuti data terbaru dari person4
	*/
	
	*person4 = Person{
		nama:   "El Rafif Zavier Putra",
		tanggal: "23 Mar 2019",
		usia:   5,
	}

	/*
		Pada cara di atas, data person1 berubah mengikuti data person4.
		Mengapa data person3 juga ikut berubah, hal itu karena 
		pointer *, pada dasarnya adalah memper barui data yang mereferensi 
		pada  struct Person, sehingga variabel yang terkoneksi pada struct 
		Person akan mengikuti perubahan terbaru, begitupun jika nanti membuat 
		variabel baru yang mereference pada struct Person
	*/

	fmt.Println("Person1", person1)
	fmt.Println("Person2", person2) // tidak merubah value pada person
	fmt.Println("Person3", person3) // merubah value pada person
	fmt.Println("Person4", person4) // merubah semua value yang mereference ke struct person

	// var person5 *Person = &Person{
	// 	nama:   "Zeanissa",
	// 	tanggal: "14 Feb 2020",
	// 	usia:   4,
	// }

	person5 := &Person{
		nama:   "Zeanissa",
		tanggal: "14 Feb 2020",
		usia:   4,
	}

	// var person6 *Person = new(Person)

	person6 := new(Person)
	person6.nama = "Nama Baru"

	fmt.Println("Person5", person5)
	fmt.Println("Person6", person6)
}
