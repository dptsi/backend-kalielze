package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Arguments", args)

	fmt.Println(args[1])
	fmt.Println(args[2])
	fmt.Println(args[3])

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hosname:", hostname)
	} else {
		fmt.Println("err msg:", err.Error())
	}

	/*
		masih banyak documentation mengenai package os

		https://golang.org/pkg/os/
	*/
}
