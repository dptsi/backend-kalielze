package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {

	// var data *ring.Ring = ring.New(10)

	data := ring.New(10)

	// data.Value = "Inisiasi 1"
	// data = data.Next()
	// data.Value = "Inisiasi 2"

	// set data ring
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value:" + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	// fungsi untuk menampilkan container ring
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})

	/*
		masih banyak documentation mengenai package container ring

		https://golang.org/pkg/container/ring/
	*/
}
