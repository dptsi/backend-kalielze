package main

import (
	"flag"
	"fmt"
)

func main() {
	var hostname *string = flag.String("hostname", "localhost", "put your database host")
	var username *string = flag.String("username", "root", "put your database user")
	var password *string = flag.String("password", "root", "put your database password")
	var port *int = flag.Int("port", 1443, "put your database port")

	flag.Parse()

	fmt.Println(*hostname)
	fmt.Println(*username)
	fmt.Println(*password)
	fmt.Println(*port)

	/*
		masih banyak documentation mengenai package flag

		http://golang.org/pkg/flag/
	*/
}
}