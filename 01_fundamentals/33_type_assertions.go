package main

import "fmt"

func main() {
	/*
		Type assertion hampir sama seperti covertion,
		namun penggunaannya lebih dikhususkan pada
		empty interface
	*/

	emptyinterface := EmptyInterface()
	result_emptyinterface := emptyinterface.(string) // assertion string, meyakini bahwa return yang dihasilkan adalah string
	fmt.Println(result_emptyinterface)

	// defer logging()
	// result_emptyinterface2 := emptyinterface.(bool) // error karena return yang dihasilkan tidak sesuai/panic
	// fmt.Println(result_emptyinterface2)

	/*
		untuk mempermudah penggunaan assertion dapat menggunakan switch
	*/
	switch result_emptyinterface2 := emptyinterface.(type) {
	case int:
		fmt.Println("int", result_emptyinterface2)
	case bool:
		fmt.Println("boolean", result_emptyinterface2)
	default:
		fmt.Println("unkown", result_emptyinterface2)
	}
}

func EmptyInterface() interface{} {
	return "Nama Karuniawan Putra"
}

func logging() {
	msg := recover()
	if msg != nil {
		fmt.Println("err panic:", msg)
	}
}