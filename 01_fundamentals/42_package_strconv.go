package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println("Boolean", boolean)
	} else {
		fmt.Println("err msg", err.Error())
	}

	integer, err := strconv.ParseInt("100000", 10, 64)
	if err == nil {
		fmt.Println("Integer", integer)
	} else {
		fmt.Println("err msg", err.Error())
	}

	value := strconv.FormatInt(125, 10) // decimal
	fmt.Println("Decimal", value)
	value = strconv.FormatInt(125, 2) // biner
	fmt.Println("Binary", value)
	value = strconv.FormatInt(125, 8) // okat
	fmt.Println("Octal", value)
	value = strconv.FormatInt(125, 16) // hexa
	fmt.Println("Hexa", value)

	atoi, err_msg := strconv.Atoi("100000")
	if err_msg == nil {
		fmt.Println("Integer", atoi)
	} else {
		fmt.Println("err msg", err_msg.Error())
	}

	value = strconv.Itoa(10000)
	fmt.Println(value)

	/*
		strconv.ParseBool(str) // merubah string menjadi boolean
		strconv.ParseInt(str) // merubah string menjadi int
		strconv.ParseFloat(str) // merubah string menjadi float
		strconv.FormatBool(bool) // merubah bool menjadi string
		strconv.ParseInt(int,..) // merubah int menjadi string
		strconv.ParseFloat(float,..) // merubah float menjadi string

	*/

	/*
		masih banyak documentation mengenai package strconv

		https://golang.org/pkg/strconv/
	*/
}