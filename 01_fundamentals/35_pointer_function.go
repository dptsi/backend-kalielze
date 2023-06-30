package main

import "fmt"

type Alamat struct{
	alamat, kota string
	no int
}

func main() {
	alamat := Alamat{
		alamat: "Florencia Regency",
		kota: "Sidoarjo",
	}
	chg_alamat(&alamat)
	fmt.Println(alamat)

	// atau
	var alamatpointer *Alamat = &alamat
	chg_alamat(alamatpointer)
	fmt.Println(alamat)
	
}

func chg_alamat(alamat *Alamat) {
	alamat.no = 25
}