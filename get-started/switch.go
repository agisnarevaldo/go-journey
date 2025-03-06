package main

import "fmt"

func main() {
	nilai := 10
	keterangan := ""
	predikat := ""

	switch nilai {
	case 10:
		keterangan = "Anda mendapat nilai sempurna"
	case 8:
		keterangan = "Anda Mendapat nilai yang baik"
	case 6:
		keterangan = "Coba lagi tahun depan"
	default:
		fmt.Println("Anda harus menyelesaikan dulu studi")
	}
	fmt.Printf("Keterangan\t: %v\n", keterangan)

	switch nilai != 0 {
	case true:
		fmt.Println("Nilai sudah muncul")
	case false:
		fmt.Println("Nilai belum muncul")
	}

	switch {
	case nilai > 9:
		predikat = "A"
	case nilai > 8:
		predikat = "B"
	case nilai > 7:
		predikat = "C"
	case nilai > 6:
		predikat = "D"
	case nilai > 0:
		predikat = "E"
	default:
		fmt.Println("Belum ada Predikat")
	}
	fmt.Printf("Predikat\t: %v\n", predikat)
}
