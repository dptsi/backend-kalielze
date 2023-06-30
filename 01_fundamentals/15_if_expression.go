package main

import "fmt"

func main() {

	map_ipk := map[rune]int{
		'A' :100,
		'B' :90,
		'C' :80,
		'E' :50,
		'D' :0,
	}

	nilai_ujian := 85

	if nilai_ujian <= map_ipk['A'] && nilai_ujian >= map_ipk['B'] {
		fmt.Println("Lulus Terbaik")
	} else if nilai_ujian <= map_ipk['B'] && nilai_ujian >= map_ipk['C'] {
		fmt.Println("Cukup Lulus")
	} else {
		fmt.Println("Mengulang")
	}

}