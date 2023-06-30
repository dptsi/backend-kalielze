package main

import "fmt"

func main() {

	/*
		Contoh penggunaan for loop untuk map
	*/

	map_nilai_peserta := map[string]int{
		"Adam":    75,
		"Lila":    87,
		"Samir":   78,
		"Zela":    83,
		"Richard": 80,
	}

	for key, value := range map_nilai_peserta {
		// fmt.Println("Menampilkan semua data map", key, value)
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// menseleksi kelulusan
	passing_grade := 80

	for key, value := range map_nilai_peserta {
		if value >= passing_grade {
			fmt.Println("Peserta yang lulus", key, value)
		} else {
			fmt.Println("Peserta yang mengulang", key, value)
		}
	}

	// melihat hanya nama tertentu
	map_selection_peserta := map[string]bool{
		"Adam": true,
		"Zela": true,
	}

	for key, value := range map_nilai_peserta {
		if map_selection_peserta[key] {
			if value >= passing_grade {
				fmt.Println("Peserta yang lulus", key, value)
			} else {
				fmt.Println("Peserta yang mengulang", key, value)
			}
		}
	}

	/*
		Contoh copy map ke slice
	*/

	slc_nilai_peserta := make([]int, 0, len(map_nilai_peserta))
	slc_nama_peserta := make([]string, 0, len(map_nilai_peserta))
	for key, value := range map_nilai_peserta {
		slc_nilai_peserta = append(slc_nilai_peserta, value)
		slc_nama_peserta = append(slc_nama_peserta, key)
	}

	fmt.Println(slc_nilai_peserta)
	fmt.Println(slc_nama_peserta)

	// slc_nama_peserta := make([]string, 0, len(map_nilai_peserta))
	// for key := range map_nilai_peserta {
	// 	slc_nama_peserta = append(slc_nama_peserta, key)
	// }

	// fmt.Println(slc_nama_peserta)

	/*
		Kesimpulan

		Slice tidak dapat langsung  copy data dari map
		menggunakan fungsi copy, karena struktur slice
		berbeda dengan map. Untuk melaukan copy data map
		ke slice, perlu menggunakan for
	*/

	for i := 0; i < len(slc_nama_peserta); i++ {
		fmt.Println(slc_nama_peserta[i], slc_nilai_peserta[i])
	}

	counter := 1

	for counter <= 10 {
		fmt.Println("Counter", counter)
		counter++
	}

}
