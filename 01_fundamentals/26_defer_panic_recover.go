package main

import "fmt"

func main() {
	// contoh penggunaan defer
	// run_app()
	/*
		Defer dijalankan secara LIFO "Last in, First out",
		sehingga defer yang dideklarasikan terakhir akan
		dirunning terebih dahulu
	*/

	run_app1(true)
	/*
		Panic akan menghentikan sebuah proses, sehingga proses
		setelah panic tidak akan dijalankan.

		Defer sebaiknya diletakkan sebelum panic, atau lebih
		disarankan diletakkan di awal func.

		Recover diletakkan paling akhir, dalam kasus defer panic recover,
		maka recover diletakkan pada func defer saja. Recover berfungsi
		untuk tetap melanjutkan proses apabila ada panic yang terjadi.
	*/
	fmt.Println("tetap jalan")
}

func run_app() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			defer logging(i)
		}
		// fmt.Println("index", i)
	}
	defer fmt.Println("selesai looping")
}

// ======================================================== //

func logging(i int) {
	fmt.Println("cek index", i)
}

func run_app1(b bool) {
	defer logging1()
	if b {
		panic("app has error")
	}
	fmt.Println("app is execute")

}

func logging1() {
	msg := recover()
	if msg != nil {
		fmt.Println("err", msg)
	}
	fmt.Println("app has done execute")
}
