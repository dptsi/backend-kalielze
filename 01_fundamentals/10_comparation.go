package main

import "fmt"

func main() {
	const nilai_lulus = 90

	var (
		nilai_budi  = 90
		nilai_ani   = 89
		nilai_iman  = 100
		nilai_totok = 77
	)

	fmt.Println("Apakah Budi lulus", nilai_budi >= nilai_lulus)
	fmt.Println("Apakah Ani lulus", nilai_ani >= nilai_lulus)
	fmt.Println("Apakah Iman lulus", nilai_iman >= nilai_lulus)
	fmt.Println("Apakah Totok lulus", nilai_totok >= nilai_lulus)
}
