package main

import "fmt"

/*
==============================================================
BAGIAN 2: VARIABEL DAN TIPE DATA
==============================================================
Ada 2 cara umum membuat variabel di Go:
1. Menggunakan kata kunci `var`:
   var namaVariable tipeData = nilai
2. Menggunakan operator pendek `:=` (Short Variable Declaration):
   namaVariable := nilai
   (Tipe data akan dideteksi secara otomatis oleh Go / Type Inference)

Tipe data dasar di Go:
- string  : teks (diapit tanda kutip dua "")
- int     : bilangan bulat (contoh: 1, 20, -5)
- float64 : bilangan desimal/koma (contoh: 3.14, 75.5)
- bool    : boolean bernilai true atau false

CARA MENJALANKAN FILE INI:
    go run 02_variabel_dan_tipe_data.go
==============================================================
*/

func main() {
	// Cara 1: Menggunakan 'var' secara eksplisit
	var namaLengkap string = "Budi Santoso"
	var umur int = 25

	// Cara 2: Menggunakan ':=' (Hanya bisa di dalam function)
	pekerjaan := "Software Engineer"
	gaji := 8500000.50 // otomatis float64
	isActive := true   // otomatis bool

	// Menampilkan nilai variabel
	fmt.Println("--- Data Karyawan ---")
	fmt.Println("Nama      :", namaLengkap)
	fmt.Println("Umur      :", umur, "tahun")
	fmt.Println("Pekerjaan :", pekerjaan)
	fmt.Printf("Gaji      : Rp %.2f\n", gaji)
	fmt.Println("Status Aktif:", isActive)

	// Format teks dengan Printf:
	// %s = string, %d = integer, %.2f = float 2 angka di belakang koma, %t = boolean
	fmt.Printf("\nRingkasan: %s (%d th) bekerja sebagai %s.\n", namaLengkap, umur, pekerjaan)
}
