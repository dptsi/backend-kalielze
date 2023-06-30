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

	nilai_ujian := 80
	nama_ujian := "Matematika"
	fmt.Println("Nilai ujian", nilai_ujian)
	fmt.Println("Pengujian", nama_ujian)

	switch nama_ujian {
	case "Bahasa":
		fmt.Println("Standart Kelulusan adalah 80")
	case "Matematika":
		fmt.Println("Standart Kelulusan adalah 85")
	default:
		fmt.Println("Standart Kelulusan adalah 75")
	}

	switch{
	case nilai_ujian <= map_ipk['A'] && nilai_ujian >= map_ipk['B']:
		fmt.Println("Nilai ujian sangat bagus")
	case nilai_ujian <= map_ipk['B'] && nilai_ujian >= map_ipk['C']:
		fmt.Println("Nilai ujian cukup")
	default:
		fmt.Println("Nilai ujian perlu ditingkatkan")
	}

	switch nama_ujian {
	case "Bahasa":
		if nilai_ujian >= 80 {
			fmt.Println("Lulus ujian Bahasa")
		} else {
			fmt.Println("Mengulang ujian Bahasa")
		}
	case "Matematika":
		if nilai_ujian >= 85 {
			fmt.Println("Lulus ujian Matematika")
		} else {
			fmt.Println("Mengulang ujian Matematika")
		}
	default:
		if nilai_ujian >= 75 {
			fmt.Println("Lulus ujian Lainnya")
		} else {
			fmt.Println("Mengulang ujian Lainnya")
		}
	}

	switch nilai_ujian >= 80 {
	case true:
		fmt.Println("Sudah baik")
	case false:
		fmt.Println("Belum cukup")
	}

}