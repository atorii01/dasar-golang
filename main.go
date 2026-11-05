package main

import "fmt"

func main() {
	HitungPBB(20000000)
}

func HitungPBB(njop float64) {

	kenapajak := njop - 15000000

	if kenapajak <=0 {
		fmt.Printf("PBB Rp = 0")
		return
	}
	
	var hasil float64

	if njop < 1000000000 {
		hasil = njop * 0.0015
	} else if njop >= 1000000000 {
		hasil = njop * 0.0025
	}
fmt.Printf("Hasil HitungPBB = Rp%.2f\n",hasil)
}	