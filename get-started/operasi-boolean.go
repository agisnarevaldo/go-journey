package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var (
		absensi    = rand.Intn(100)
		nilai      = rand.Intn(100)
		nilaiAkhir = absensi * nilai / 100
		lulus      bool
		keterangan string
		huruf      string
	)

	switch {
	case nilaiAkhir <= 50:
		lulus = false
		keterangan = "Tidak Lulus"
		huruf = "D"
	case nilaiAkhir > 50:
		lulus = true
		keterangan = "Lulus Bersyarat"
		huruf = "C"
	case nilaiAkhir > 70:
		lulus = true
		keterangan = "Lulus"
		huruf = "B"
	case nilaiAkhir > 80:
		lulus = true
		keterangan = "Lulus"
		huruf = "A"
	}

	fmt.Printf("Nilai Akhir: %d\n", nilaiAkhir)
	fmt.Printf("Lulus: %t\n", lulus)
	fmt.Printf("Keterangan: %s\n", keterangan)
	fmt.Printf("Huruf: %s\n", huruf)

}
