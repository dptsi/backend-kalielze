package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "     Learn golang, programming is amazing to my experience     "
	fmt.Println(str)

	trim := strings.Trim(str, " ") // menghilangkankelebihan karakter diawal dan akhir string
	fmt.Println(trim) 

	tolowwer := strings.ToLower(trim) // membuat string menjadi huruf kecil semua
	fmt.Println(tolowwer) 

	toupper := strings.ToUpper(trim) // membuat string menjadi huruf besar semua
	fmt.Println(toupper) 

	contains := strings.Contains(trim, "ama") // melihat apakah ada potongan string yang serupa
	fmt.Println(contains) 

	split := strings.Split(trim, ",") // memotong dan menjadikan array berdasarkan separator
	fmt.Println(split) 
	fmt.Println(len(split))

	replaceall := strings.ReplaceAll(trim, "r", "R") // merubah semua string yang akan direplace
	fmt.Println(replaceall)

	/*
		masih banyak documentation mengenai package strings

		http://golang.org/pkg/strings/
	*/
}