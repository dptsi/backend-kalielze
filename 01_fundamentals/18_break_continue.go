package main

import "fmt"

func main() {

	arr_angka := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
	}

	for i := 0; i < len(arr_angka); i++ {
		if arr_angka[i] > 7 {
			break
		}
		fmt.Println(arr_angka[i])
	}

	for i := 0; i < len(arr_angka); i++ {
		if arr_angka[i]%3 == 0 {
			continue
		}
		fmt.Println(arr_angka[i])
	}
}
