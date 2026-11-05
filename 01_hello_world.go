package main

import "fmt"

/*
==============================================================
BAGIAN 1: STRUKTUR DASAR PROGRAM GO (HELLO WORLD)
==============================================================
- package main:
  Menandakan bahwa file ini adalah program utama yang dapat
  dijalankan langsung (executable), bukan library.

- import "fmt":
  Mengimpor package bawaan Go bernama "fmt" (format)
  yang berguna untuk operasi input/output ke layar/terminal.

- func main():
  Fungsi utama (entry point). Program Go selalu mulai berjalan
  dari dalam fungsi main().

CARA MENJALANKAN FILE INI:
Buka terminal dan ketik:
    go run 01_hello_world.go
==============================================================
*/

func main() {
	fmt.Println("Halo, selamat datang di dunia Go (Golang)!")
	fmt.Println("Ini adalah program Go pertamamu.")
}

