package main

import "fmt"

/*
==============================================================
BAGIAN 3: PERCABANGAN (IF - ELSE & SWITCH)
==============================================================
Aturan pada `if` di Go:
- Kondisi TIDAK membutuhkan tanda kurung kurawal `( ... )`.
- Kurung kurawal pembuka `{` HARUS berada di baris yang sama
  dengan pernyataan `if` atau `else`.

Fitur Spesial Go:
- If dengan pernyataan singkat (short statement):
  if nilai := hitung(); nilai > 70 { ... }
- Switch di Go tidak memerlukan kata kunci 'break' di setiap case.

CARA MENJALANKAN FILE INI:
    go run 03_percabangan_if_else.go
==============================================================
*/

func main() {
	nilaiUjian := 82

	fmt.Println("Nilai Ujian Anda:", nilaiUjian)

	// 1. Percabangan if - else if - else
	if nilaiUjian >= 85 {
		fmt.Println("Predikat: A (Sangat Baik)")
	} else if nilaiUjian >= 75 {
		fmt.Println("Predikat: B (Baik)")
	} else if nilaiUjian >= 60 {
		fmt.Println("Predikat: C (Cukup)")
	} else {
		fmt.Println("Predikat: D (Perlu Belajar Lagi)")
	}

	// 2. Percabangan menggunakan switch-case
	hari := "Senin"
	fmt.Print("Status Hari: ")
	switch hari {
	case "Sabtu", "Minggu":
		fmt.Println("Libur akhir pekan! 🎉")
	default:
		fmt.Println("Hari kerja/sekolah! 💼")
	}
}

