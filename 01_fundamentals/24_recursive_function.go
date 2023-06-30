package main

import "fmt"

func main() {
	fmt.Println("Menggunakan func", factorial(5))
	fmt.Println("Hitung biasa", (5 * 4 * 3 * 2 * 1))

	fmt.Println("Recursive func", factorial_recursive(5))
}

func factorial(nilai int64) (f int64) {
	f = 1
	for i := nilai; i > 0; i-- {
		f *= i
	}
	return
}

func factorial_recursive(nilai int64) int64 {
	if nilai == 1 {
		return 1
	} else {
		return nilai * factorial_recursive(nilai-1)
	}
}
