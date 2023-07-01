package main

import (
	"fmt"
	"regexp"
)

func main() {

	var regexp *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regexp.MatchString("edo"))
	fmt.Println(regexp.MatchString("ello"))
	fmt.Println(regexp.MatchString("eMmo"))

	// src := regexp.FindAllString("eko eti eki elo eMo eka", -1) // mengambil semua
	src := regexp.FindAllString("eko eti eki elo eMo eka", 3) // mengambil dengan batasan
	fmt.Println(src)

	/*
		masih banyak documentation mengenai package regexp

		https://github.com/google/re2/wiki/Syntax
		https://golang.org/pkg/regexp/
	*/
}
