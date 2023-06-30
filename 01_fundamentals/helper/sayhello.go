package helper

import "fmt"

var version = "1.0.0" // tidak dapat diakses dari luar
var App = "Golang"

func SayHello(name string) {
	fmt.Println("Name:", name)
}

// tidak dapat diakses di luar package
func SayGoodbye() {
	fmt.Println("Selamat Tinggal")
}
