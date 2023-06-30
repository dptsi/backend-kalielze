package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(pembagian(5,3))

	hasil, err := pembagian(5,0)
	if err == nil {
		fmt.Println("hasil", hasil)
	} else {
		fmt.Println("error", err)
	}
}

func pembagian(nilai float64, pembagi float64) (float64, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh 0")
	} else {
		return nilai/pembagi, nil
	}
}