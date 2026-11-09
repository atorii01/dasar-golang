package lasagna

// 1. Konstanta waktu oven (40 menit)
const OvenTime = 40

// 2. Menghitung sisa waktu di oven
func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

// 3. Menghitung waktu persiapan (2 menit per lapis)
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

// 4. Menghitung total waktu yang sudah berjalan
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}