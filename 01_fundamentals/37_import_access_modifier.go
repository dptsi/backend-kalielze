package main

import (
	"01_fundamentals/helper" // import
	"fmt"
)

func main() {
	helper.SayHello("Karuniawan") // access modifier
	// helper.sayGoodbye() // error

	// fmt.Println(helper.version) // error
	fmt.Println(helper.App)
}
