package main

import "fmt"

func main() {
	count := 0
	name := "Wawan"
	fmt.Println("Nilai sebelum", count)
	fmt.Println("Nama sebelum", name)
	incr := func() {
		// name = "Putra" // merubah value name
		// name := "Putra" // tidak merubah value name
		count++
		// fmt.Println(name)
	}
	incr()
	fmt.Println("Nilai sesudah", count)
	fmt.Println("Nama sesudah", name)
}
