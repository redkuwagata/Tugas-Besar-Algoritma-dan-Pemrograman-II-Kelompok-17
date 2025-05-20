package main

import "fmt"

type Pemasukan struct {
	nama     string
	jumlah   float64
	frekuensi string
	kategori string
}

const MAKS int = 100

type DaftarPemasukan [MAKS]Pemasukan

func inputPemasukan(data *DaftarPemasukan, jumlah *int) {
	var lanjut string = "y"
	for lanjut == "y" && *jumlah < MAKS {
		var p Pemasukan
		fmt.Print("\nNama Pemasukan (gaji/freelance/usaha/loyalti): ")
		fmt.Scanln(&p.nama)
		fmt.Print("Jumlah Pemasukan (Rp): ")
		fmt.Scanln(&p.jumlah)
		fmt.Print("Jangka Waktu (bulanan/tahunan): ")
		fmt.Scanln(&p.frekuensi)
		fmt.Print("Kategori (main/side/passive): ")
		fmt.Scanln(&p.kategori)

		data[*jumlah] = p
		*jumlah = *jumlah + 1

		fmt.Print("Tambah Pemasukan Lagi? (yes/no): ")
		fmt.Scanln(&lanjut)
	}
}

func konversiKeBulanan(jumlah float64, frekuensi string) float64 {
	if frekuensi == "tahunan" {
		return jumlah / 12
	}
	return jumlah
}

func hitungTotal(data DaftarPemasukan, jumlah int) (float64, float64, float64) {
	var totalMain, totalSide, totalPasif float64
	for i := 0; i < jumlah; i++ {
		bulanan := konversiKeBulanan(data[i].jumlah, data[i].frekuensi)
		if data[i].kategori == "main" {
			totalMain += bulanan
		} else if data[i].kategori == "side" {
			totalSide += bulanan
		} else if data[i].kategori == "passive" {
			totalPasif += bulanan
		}
	}
	return totalMain, totalSide, totalPasif
}