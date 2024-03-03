package main

import (
	"fmt"
	"os"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func cariTeman(absen int) Teman {
	daftarTeman := map[int]Teman{
		1: {"Roni", "Jl. Merdeka No. 1", "Mahasiswa", "Ingin belajar bahasa pemrograman Go"},
		2: {"Bujang", "Jl. Damai No. 10", "Designer", "Ingin memperluas skill programming"},
		3: {"Anton", "Jl. Sejahtera No. 5", "Developer", "Membutuhkan bahasa pemrograman baru untuk pekerjaan"},
	}

	teman, found := daftarTeman[absen]
	if !found {
		return Teman{}
	}
	return teman
}

func main() {
	absen := os.Args[1]
	var noAbsen int

	fmt.Sscanf(absen, "%d", &noAbsen)

	teman := cariTeman(noAbsen)

	if teman.Nama == "" {
		fmt.Println("Teman dengan nomor absen tersebut tidak ditemukan")
		return
	}

	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)

}
