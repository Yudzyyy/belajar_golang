package main

import "fmt"

func main() {
	nama := "Andi"
	umur := 23

	fmt.Println("Nama:", nama)

	if umur >= 18 {
		fmt.Println("Status: Dewasa")
	} else {
		fmt.Println("Status: Anak-anak")
	}

	for i := 1; i <= 3; i++ {
		fmt.Println("Perulangan ke-", i)
	}

	hari := "Sabtu"

	switch hari {
	case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
		fmt.Println("Hari kerja")
	case "Sabtu", "Minggu":
		fmt.Println("Hari libur")
	default:
		fmt.Println("Hari tidak valid")
	}
}
