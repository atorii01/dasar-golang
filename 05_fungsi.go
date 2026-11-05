package main

import "fmt"

/*
==============================================================
BAGIAN 5: FUNGSI (FUNCTION)
==============================================================
Fungsi di Go dibuat dengan kata kunci `func`.
Sintaks umum:
    func namaFungsi(param1 tipe, param2 tipe) tipeKembalian {
        return nilai
    }

Fitur Unggulan Go:
- Fungsi dapat mengembalikan lebih dari 1 nilai (Multiple Return Values).
  Fitur ini sangat sering dipakai untuk mengembalikan (hasil, error).

CARA MENJALANKAN FILE INI:
    go run 05_fungsi.go
==============================================================
*/

// Fungsi tanpa return value
func sapa(nama string) {
	fmt.Printf("Halo, %s! Senang belajar Go bersamamu.\n", nama)
}

// Fungsi dengan 1 return value
func hitungLuasPersegi(panjang int, lebar int) int {
	return panjang * lebar
}

// Fungsi dengan multiple return values (mengembalikan hasil bagi & sisa bagi)
func bagiDanSisa(angka int, pembagi int) (int, int) {
	hasilBagi := angka / pembagi
	sisa := angka % pembagi
	return hasilBagi, sisa
}

func main() {
	// 1. Memanggil fungsi biasa
	sapa("Andi")

	// 2. Memanggil fungsi dengan return value
	luas := hitungLuasPersegi(10, 5)
	fmt.Println("Luas Persegi Panjang (10 x 5) =", luas)

	// 3. Memanggil fungsi dengan 2 return values sekaligus
	hasil, sisa := bagiDanSisa(14, 4)
	fmt.Printf("14 dibagi 4 = %d (sisa %d)\n", hasil, sisa)
}

