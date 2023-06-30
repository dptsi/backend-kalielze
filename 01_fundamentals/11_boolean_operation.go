package main

import "fmt"

func main() {
	const std_lulus = 85

	var (
		nilai_budi  = 90
		nilai_ani   = 89
		nilai_iman  = 100
		nilai_totok = 77
	)

	hasil_budi := nilai_budi >= std_lulus
	hasil_ani := nilai_ani >= std_lulus
	hasil_iman := nilai_iman >= std_lulus
	hasil_totok := nilai_totok >= std_lulus

	fmt.Println("Apakah Budi lulus", hasil_budi)
	fmt.Println("Apakah Ani lulus", hasil_ani)
	fmt.Println("Apakah Iman lulus", hasil_iman)
	fmt.Println("Apakah Totok lulus", hasil_totok)
}
