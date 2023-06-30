package main

import "fmt"

type Pegawai struct{
	nama string
}

func main() {
	pegawai := Pegawai{"Karuniawan Putra"}
	pegawai.chg_nama()
	fmt.Println(pegawai)
	
}

func (pegawai *Pegawai) chg_nama() {
	pegawai.nama = "Mr." + pegawai.nama
}