package main

import (
	"fmt"
	"sort"
)

type User struct {
	nama string
	usia int
}

type UserSlice []User

func main() {

	users := []User{
		{"Wawan", 30},
		{"El", 5},
		{"Lita", 28},
		{"Zea", 4},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)

	/*
		masih banyak documentation mengenai package sort

		http://golang.org/pkg/sort/
	*/
}

func (userslice UserSlice) Len() int {
	return len(userslice)
}

func (userslice UserSlice) Less(i, j int) bool {
	return userslice[i].nama < userslice[j].nama
}

func (userslice UserSlice) Swap(i, j int) {
	userslice[i], userslice[j] = userslice[j], userslice[i]
}
