package main

import "fmt"

func main() {
	// var mp_pegawai map[string]string = nil
	mp_pegawai := nil_test(false)
	if mp_pegawai == nil {
		fmt.Println("data map kosong")
	} else {
		fmt.Println(mp_pegawai)
	}
}

/*
	nil hanya bisa digunakan untuk inisiasi pada
	interface, func, map, slice, pointer dan channel
*/
func nil_test(b bool) map[string]string {
	if b {
		return map[string]string{
			"nama":   "Karunaiwan Putra",
			"alamat": "Ngagel Dadi 3 No 6 Surabaya",
		}
	} else {
		return nil
	}
}
