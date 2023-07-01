package main

import (
	"container/list"
	"fmt"
)

func main() {

	data := list.New()

	data.PushBack("data-pertama")
	data.PushBack("data-kedua")
	data.PushBack("data-ketiga")
	data.PushBack("data-keempat")
	data.PushBack("data-kelima")

	fmt.Println("Front:", data.Front().Value)
	fmt.Println("Back:", data.Back().Value)

	fmt.Println("Front:", data.Front().Prev()) // nil karena data sebelum front tidak ada
	fmt.Println("Back:", data.Back().Next()) // nil karena data sesudah back tidak ada

	// iterasi dari depan
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println("iteration from front:", e.Value)
	}

	//iterasi dari belakang
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println("iteration from back:", e.Value)
	}

	/*
		masih banyak documentation mengenai package container list

		https://golang.org/pkg/container/list/
	*/
}
