package main

import (
	"fmt"
	"strconv"
)

func main() {

	// ================= simpel helloworld ================= //
	fmt.Println("================= simpel helloworld =================")
	fmt.Println("Helloworld")
	fmt.Println("======================== *** ========================\n")

	// simple number
	fmt.Println("================= simpel number =================")
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga = ", 3.5)
	fmt.Println("======================== *** ========================\n")

	// ================= simple boolean ================= //
	fmt.Println("================= simpel boolean =================")
	fmt.Println("Benar = ", true)
	fmt.Println("Salah = ", false)
	fmt.Println("======================== *** ========================\n")

	// ================= simple string ================= //
	fmt.Println("================= simpel string =================")
	fmt.Println("Helloworld")
	fmt.Println(len("Helloworld")) // mengetahui panjang string
	fmt.Println("Helloworld"[0])   // mencari byte karakter string
	fmt.Println("======================== *** ========================\n")

	// simple variable
	fmt.Println("================= simpel variable =================\n")
	fmt.Println("~~~~~~~~~~~~~~ contoh 1 - step by step ~~~~~~~~~~~~~~")
	/***contoh 1 - step by step***/
	var stringHello string
	stringHello = "Helloworld" // pendeklarasian terpisah
	fmt.Println(stringHello)
	stringHello = "newHelloworld" // deklarasi dengan string baru
	fmt.Println(stringHello)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~ *** ~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("~~~~~~~~~~~~~~ contoh 2 - direct ~~~~~~~~~~~~~~")
	/***contoh 2 - direct***/
	var stringWorld = "Helloworld" // pendeklarasian langsung, tidak perlu mendefinisikan variable
	fmt.Println(stringWorld)
	stringWorld = "newHelloworld" // deklarasi dengan string baru
	fmt.Println(stringWorld)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~ *** ~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("~~~~~~~~~~~~~~ contoh 3 - without var ~~~~~~~~~~~~~~")
	/***contoh 3 - without var***/
	stringNovar := "Helloworld"
	// jika ditulis stringNovar string := "Helloworld" -> akan mengakibatkan error,
	// karena := tidak perlu pedeklarasian tipe variabel
	// dan := hanya untuk pendeklarasian awal
	fmt.Println(stringNovar)
	stringNovar = "newHelloworld" // deklarasi dengan string baru
	fmt.Println(stringNovar)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~ *** ~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("~~~~~~~~~~~~~~ contoh 4 - multiple var ~~~~~~~~~~~~~~")
	/***contoh 4 - multiple var***/
	var (
		tgl_lahir = "16/11/1992"
		tmp_lahir = "Jombang"
	)
	// multiple var harus dideklarasikan menggunakan var ()
	fmt.Println(tgl_lahir)
	fmt.Println(tmp_lahir)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~ *** ~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("======================== *** ========================\n")

	// simple constant
	fmt.Println("================= simpel constant =================")
	const firstname string = "Karuniawan"
	const lastname = "Putra"
	const age = 31
	/*constant tidap perlu di panggil tidak apa-apa*/
	const (
		awal  string = "Karuniawan"
		akhir        = "Putra"
		nomor        = 31
	)
	/*constant sama seperti variabel dapat dibuat multiple constant*/
	fmt.Println("======================== *** ========================\n")

	// simple convertion
	fmt.Println("================= simpel convertion =================")
	var value32 int32 = 129
	var value64 int64 = int64(value32)
	var value8 int8 = int8(value32)
	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value8)
	var alias = "Wawan"
	var chr = alias[0]         // dideklarasikan sebagai byte
	var chrAlias = string(chr) // untuk mendefinisikan karakter yang telah diambil di variabel e
	fmt.Println(chrAlias)
	fmt.Println(string(chr))
	fmt.Println("======================== *** ========================\n")

	// simple declaration
	fmt.Println("================= simpel declaration =================")
	type noktp string
	var noKTP noktp = "351507" // noktp menjadi sebuah tipe data yang dapat digunakan diberbagai variabel yang bernilai string
	fmt.Println(noKTP)
	fmt.Println("======================== *** ========================\n")

	// simple math operation
	fmt.Println("================= simpel math operation =================")
	var (
		a uint64 = 18446744073709551615 // batas maksimal uint64 0 sampai 18446744073709551615
		b uint64 = 16111992
		c uint64 = a - b
		d uint64 = a + b
	)
	fmt.Println(c)
	fmt.Println(d) // tidak dapat menunjukkan hasil yang sebenarnya karena ada batas maksimal tertinggi
	b += 19071994
	fmt.Println(b)
	b++
	fmt.Println(b)
	var e int64 = -9223372036854775808 // batas int644 -9223372036854775808 sampai 9223372036854775807
	fmt.Println(e)
	fmt.Println("======================== *** ========================\n")

	// simple compare operation
	fmt.Println("================= simpel compare operation =================")
	var (
		g           = 6
		h           = 7
		result bool = g == h
	)
	fmt.Println(result)
	result = g <= 6
	fmt.Println(result)
	fmt.Println("======================== *** ========================\n")

	// simple boolean operation
	fmt.Println("================= simpel boolean operation =================")
	var (
		i          = 100
		j          = 200
		k     bool = i >= 150
		l     bool = j >= 150
		lulus bool = k && l
	)
	fmt.Println(lulus)
	fmt.Println("======================== *** ========================\n")

	// simple array
	fmt.Println("================= simpel array =================")
	var arrayString [2]string
	arrayString[0] = "Karuniawan"
	arrayString[1] = "Putra"
	fmt.Println(arrayString[0], arrayString[1])
	var arrayInteger = [3]int{
		1, 2, 3,
	}
	fmt.Println(arrayInteger[0], arrayInteger[1], arrayInteger[2])
	fmt.Println(len(arrayString))
	fmt.Println(len(arrayInteger))
	fmt.Println("======================== *** ========================\n")

	// simple slice
	fmt.Println("================= simpel slice =================\n")
	fmt.Println("~~~~~~~~~~~~~~ contoh 1 - slice standart ~~~~~~~~~~~~~~")
	/***contoh 1 - slice standart***/
	arrayBulan := [...]string{
		"Jan", "Feb", "Mar", "Apr", "Mei", "Jun",
		"Jul", "Agu", "Sep", "Okt", "Nov", "Des",
	}
	fmt.Println(arrayBulan[4:7])
	fmt.Println(len(arrayBulan[4:7]))
	fmt.Println(cap(arrayBulan[4:7]))
	arrayBulan[5] = "Diubah"
	fmt.Println(arrayBulan[4:7])
	arrayBulan[4:7][0] = "Bukan-Mei-Lagi"
	fmt.Println(arrayBulan[4:7])
	fmt.Println(arrayBulan) // karena data sudah dirubah lewat slice, maka data array ikut berubah
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~ *** ~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("~~~~~~~~~~~~~~ contoh 2 - slice append ~~~~~~~~~~~~~~")
	/***contoh 2 - slice append***/
	arrayHari := [...]string{
		"Senin", "Selasa", "Rabu", "Kamis",
		"Jum'at", "Sabtu", "Minggu",
	}
	sliceHari := arrayHari[5:]
	sliceHari[0] = "sabtuGabut"
	sliceHari[1] = "mingguLembur"
	fmt.Println(arrayHari)
	sliceHari2 := append(sliceHari, "lemburTambahan")
	sliceHari2[0] = "omg!"
	fmt.Println(sliceHari2)
	fmt.Println(arrayHari) // arrayHari tidak berubah karena append, karena apprend membentuk array baru yang melebihi kapasitas arrayHari
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~ *** ~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("~~~~~~~~~~~~~~ contoh 3 - make slice ~~~~~~~~~~~~~~")
	/***contoh 3 - make slice***/
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Karuniawan"
	newSlice[1] = "Putra"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~ *** ~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("~~~~~~~~~~~~~~ contoh 4 - copy slice ~~~~~~~~~~~~~~")
	/***contoh 4 - copy slice***/
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
	copySlice = append(arrayBulan[8:], "test")
	fmt.Println(copySlice)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~ *** ~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Println("======================== *** ========================\n")

	// simple map
	fmt.Println("================= simpel map =================")
	person := map[string]string{
		"name":    "Karuniawan Putra",
		"address": "Ngagel Dadi",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	person["kantor"] = "ITS"
	fmt.Println(person)
	fmt.Println(len(person))
	book := make(map[string]string)
	book["title"] = "Pemrograman Golang"
	book["autror"] = "Karuniawan"
	book["owner"] = "PT Dua Lima"
	fmt.Println(book)
	delete(book, "owner")
	fmt.Println(book)
	fmt.Println("======================== *** ========================\n")

	// simple if expression
	fmt.Println("================= simpel if expression =================")
	if person["name"] == "karuniawan" {
		fmt.Println("benar, ini karuniawan")
	} else {
		fmt.Println("salah, ini bukan karuniawan")
	}
	/***short statement if expression***/
	if length := len(person["name"]); length <= 5 {
		fmt.Println(length)
		fmt.Println("benar")
	} else {
		fmt.Println(length)
		fmt.Println("salah")
	}
	fmt.Println("======================== *** ========================\n")

	// simple switch expression
	// simple if expression
	fmt.Println("================= simpel switch expression =================")
	switch person["name"] {
	case "Karuniawan":
		fmt.Println("Karuniawan")
	case "Putra":
		fmt.Println("Putra")
	default:
		fmt.Println("Anda Karuniawan Putra")
	}
	/***short statement switch expression***/
	switch length := len(person["name"]); length <= 5 {
	case true:
		fmt.Println(length)
		fmt.Println("benar")
	case false:
		fmt.Println(length)
		fmt.Println("salah")
	}
	/***short without statement switch expression***/
	switch {
	case len(person["name"]) >= 5:
		fmt.Println(person["name"])
		fmt.Println("Nama ", person["name"], " sudah panjang")
	default:
		fmt.Println(person["name"])
		fmt.Println("Nama sudah cukup")
	}
	fmt.Println("======================== *** ========================\n")

	// simple for loop
	fmt.Println("================= simpel for loop =================")
	for cnt := 0; cnt < len(arrayHari); cnt++ {
		fmt.Println(arrayHari[cnt])
		//cnt++
	}
	for i, newcnt := range arrayHari {
		fmt.Println("Index", i, "=", newcnt)
	}
	for _, newcnt := range arrayHari {
		fmt.Println(newcnt)
	}
	for key, newcnt := range person {
		fmt.Println(key, newcnt)
	}
	fmt.Println("======================== *** ========================\n")

	// simple break & continue
	fmt.Println("================= simpel break & continue =================")
	for br_cn := 0; br_cn < 10; br_cn++ {
		if br_cn == 5 {
			break
		}
		fmt.Println("break", br_cn)
	}
	for br_cn := 0; br_cn < 10; br_cn++ {
		if br_cn%2 == 0 {
			continue
		}
		fmt.Println("continue", br_cn)
	}
	fmt.Println("======================== *** ========================\n")

	// simple function
	fmt.Println("================= simpel function =================")
	funcStandart()
	funcParameter("parameter")
	fmt.Println("function return", funcReturn(75))
	tinggi, satuan := funcReturnMultiple(75)
	fmt.Println("function return multiple", tinggi, satuan)
	newTinggi, _ := funcReturnMultiple(75)
	fmt.Println("function return multiple", newTinggi)
	namedTinggi, namedSatuan := funcReturnNamed(75)
	fmt.Println("function return named", namedTinggi, namedSatuan)
	fmt.Println("======================== *** ========================\n")

	// simple varags/variadic function
	fmt.Println("================= simpel varags/variadic function =================")
	total := sumAll("data", 75, 65, 89)
	fmt.Println("total", total)

	slice := []int{4, 5, 6, 7, 8}
	total = sumAll("data", slice...)
	fmt.Println("total", total)
	fmt.Println("======================== *** ========================\n")

	// simple value function
	fmt.Println("================= simpel value function =================")
	funcValue := funcGetValue
	fmt.Println("variable func", funcValue("Karuniawan"))
	fmt.Println("get func", funcGetValue("Karuniawan"))
	fmt.Println("======================== *** ========================\n")

	// simple paremeter function
	fmt.Println("================= simpel parameter function =================")
	wordData := wordGetData
	wordFilter("peh peh peh", wordData)
	wordFilter("Jancuk", wordGetData)
	paymentFilter("bagaimana gaji saya?", 6500000, paymentGetData)
	fmt.Println("======================== *** ========================\n")

	// simple anonymous function
	fmt.Println("================= simpel anonymous function =================")
	blacklist := func(name string) bool {
		return name == "Putra"
	}
	guesFilter("Putra", blacklist)
	guesFilter("karunaiwan", func(name string) bool {
		return name == "Putra"
	})
	fmt.Println("======================== *** ========================\n")

	// simple recursive function
	fmt.Println("================= simpel recursive function =================")
	fmt.Println(factorialLoop(5))
	fmt.Println(5 * 4 * 3 * 2 * 1)
	fmt.Println(factorialRecursive(5))
	fmt.Println("======================== *** ========================\n")

	// simple closure
	fmt.Println("================= simpel closure =================")
	nm := "wawan"
	ctr := 0
	inc := func() {
		nm = "putra"
		fmt.Println("Increment")
		ctr++
	}
	inc()
	fmt.Println(ctr)
	inc()
	fmt.Println(ctr)
	fmt.Println(nm)
	fmt.Println("======================== *** ========================\n")

	// simple dever, panic & recover
	fmt.Println("================= simpel dever, panic & recover =================")
	runFunction(0)
	fmt.Println("======================== *** ========================\n")

	// simple struct
	fmt.Println("================= simpel struct =================")
	var propertyHuman People
	propertyHuman.Name = "Karuniawan Putra"
	propertyHuman.Address = "Ngagel Dadi"
	propertyHuman.Age = 30
	fmt.Println(propertyHuman.Name)
	fmt.Println(propertyHuman.Address)
	fmt.Println(propertyHuman.Age)
	human := People{
		Name:    "Karunaiwan Putra",
		Address: "Ngagel Dadi",
		Age:     30,
	}
	fmt.Println(human)
	human2 := People{"karuniawan Putra", "Ngagel Dadi", 30}
	fmt.Println(human2)
	fmt.Println("======================== *** ========================\n")
}

func funcStandart() {
	fmt.Println("function standart")
}

func funcParameter(a string) {
	fmt.Println("function", a)
}

func funcReturn(b int64) int64 {
	return 100 + b
}

func funcReturnMultiple(b int64) (int64, string) {
	return 100 + b, "cm"
}

func funcReturnNamed(b int64) (tinggi int64, satuan string) {
	tinggi = 100 + b
	satuan = "cm"
	return
}

func sumAll(name string, number ...int) int {
	total := 0
	for _, value := range number {
		total += value
	}
	return total
}

func funcGetValue(name string) int {
	return 100
}

func wordFilter(word string, filter func(string) string) {
	wordFiltered := filter(word)
	fmt.Println(wordFiltered)
}

func wordGetData(word string) string {
	if word == "Jancuk" {
		return "Sampeyan wong surabaya cak!"
	} else {
		return "guduk surabaya ta?"
	}
}

type paidFilter func(int64, int64) string

func paymentFilter(question string, fee int64, filter paidFilter) {
	var umk int64 = 4500000
	wordFiltered := filter(fee, umk)
	fmt.Println(question, wordFiltered)
}

func paymentGetData(fee int64, umk int64) string {
	if fee >= umk {
		return "bayaran anda tinggi " + strconv.FormatInt(fee, 10)
	} else {
		return "bayaran anda rendah " + strconv.FormatInt(fee, 10)
	}
}

type blacklist func(string) bool

func guesFilter(nameGuest string, cekBlacklist blacklist) {
	if cekBlacklist(nameGuest) {
		fmt.Println("blocked", nameGuest)
	} else {
		fmt.Println("welcome", nameGuest)
	}
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func defferLogging() {
	fmt.Println("selesai memanggil function")
	msg := recover() // recover sebaiknya disimpan dalam deffer
	if msg != nil {
		fmt.Println(msg)
	}
}

func runFunction(value int) {
	defer defferLogging()
	a := 10 / value
	if value == 0 {
		panic("interger cann't divide by zero")
	} else {
		fmt.Println("Hasil 10 /", value, "adalah", a)
	}
}

type People struct {
	Name, Address string
	Age           int
}
