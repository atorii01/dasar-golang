package main

import "fmt"

/*
==============================================================
BAGIAN 4: PERULANGAN (FOR LOOP)
==============================================================
Uniknya, Go HANYA memiliki satu kata kunci untuk perulangan,
yaitu `for`. Tidak ada `while` atau `do-while`.
Semua variasi perulangan diselesaikan menggunakan `for`.

Variasi for di Go:
1. Standard loop: for init; kondisi; post { }
2. While-style loop: for kondisi { }
3. Infinite loop: for { } (dihentikan dengan `break`)

CARA MENJALANKAN FILE INI:
    go run 04_perulangan_for.go
==============================================================
*/

func main() {
	// 1. Perulangan Standar (dari 1 sampai 5)
	fmt.Println("--- 1. Loop Standar ---")
	for i := 1; i <= 5; i++ {
		fmt.Printf("Putaran ke-%d\n", i)
	}

	// 2. Perulangan seperti 'while' (hanya ada kondisi)
	fmt.Println("\n--- 2. Loop Gaya 'While' ---")
	hitungMundur := 3
	for hitungMundur > 0 {
		fmt.Println("Hitung:", hitungMundur)
		hitungMundur--
	}
	fmt.Println("Mulai!")

	// 3. Menggunakan 'break' dan 'continue'
	fmt.Println("\n--- 3. Menggunakan Continue & Break ---")
	for angka := 1; angka <= 10; angka++ {
		if angka%2 == 0 {
			// Lewati angka genap (hanya tampilkan angka ganjil)
			continue
		}
		if angka > 7 {
			// Berhenti jika sudah lebih dari 7
			break
		}
		fmt.Println("Angka ganjil:", angka)
	}
}

