package main

import "fmt"

type AttributPegawai struct {
	nama, nip string
	usia      int
}

func main() {
	supervisor := AttributPegawai{nama:"Zainal"}
	supervisor.supervisor()

	manager := AttributPegawai{nama:"Roni"}
	manager.manager(supervisor)
}

func (supervisor AttributPegawai) supervisor() {
	fmt.Println("Nama supervisor", supervisor.nama)
}

func (manager AttributPegawai) manager(supervisor AttributPegawai) {
	fmt.Println("Hallo supervisor", supervisor.nama, ", saya adalah manager", manager.nama)
}