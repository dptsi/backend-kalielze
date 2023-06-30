package main

import "fmt"

func main() {
	interfaceempty := InterfaceEmpty(2)
	fmt.Println(interfaceempty)
}

func InterfaceEmpty(i int) interface{} {
	if i == 1 {
		return "satu"
	} else if i == 2 {
		return true
	} else {
		return i
	}
}

/*
	Kesimpulan
	Empty interace digunakan jika memang belum tau
	tipe data apa yang akan digunakan. Tapi mungkin
	nanti dalam penggunaan perlu adanya convert data
	agar data hasil empty interface dapat digunakan
*/
