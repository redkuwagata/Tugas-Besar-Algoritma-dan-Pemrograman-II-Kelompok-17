package main
import "fmt"
type Pemasukan struct {
	nama     string
	jumlah   float64
	frekuensi string
	kategori string
}
const NMAKS int = 100
type tabPemasukan [NMAKS]Pemasukan
func inputPemasukan(data *tabPemasukan, jumlah *int) {
	var lanjut string 
		lanjut = "y"
	for lanjut == "y" && *jumlah < NMAKS {
		var p Pemasukan
		fmt.Println("\nNama Pemasukan (gaji/freelance/usaha/loyalti): ")
		fmt.Scan(&p.nama)
		fmt.Println("\nJumlah Pemasukan (Rp): ")
		fmt.Scan(&p.jumlah)
		fmt.Println("\nJangka Waktu (bulanan/tahunan): ")
		fmt.Scan(&p.frekuensi)
		fmt.Println("\nKategori (main/side/passive): ")
		fmt.Scan(&p.kategori)
		data[*jumlah] = p
		*jumlah = *jumlah + 1
		fmt.Println("\nTambah Pemasukan Lagi? (y/n): ")
		fmt.Scan(&lanjut)
	}
}
func konversiKeBulanan(jumlah float64, frekuensi string) float64 {
	if frekuensi == "tahunan" {
		return jumlah / 12
	}
	return jumlah
}
func hitungTotal(data tabPemasukan, jumlah int) (float64, float64, float64) {
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
func Ringkasan(main, side, pasif float64) {
	totalBulanan := main + side + pasif
	totalTahunan := totalBulanan * 12

	fmt.Println("\n======= RINGKASAN =======")
	fmt.Printf("Main    : Rp%.2f/bulan\n", main)
	fmt.Printf("Side    : Rp%.2f/bulan\n", side)
	fmt.Printf("Pasif   : Rp%.2f/bulan\n", pasif)
	fmt.Printf("TOTAL   : Rp%.2f/bulan | Rp%.2f/tahun\n", totalBulanan, totalTahunan)

	fmt.Println("\n======= REKOMENDASI =======")
	fmt.Printf("Tabungan 10%% : Rp%.2f\n", totalBulanan*0.10)
	fmt.Printf("Investasi 30%%: Rp%.2f\n", totalBulanan*0.30)
	fmt.Printf("Kebutuhan 60%%: Rp%.2f\n", totalBulanan*0.60)
}
func Investasi(investasi float64) {
	var pilihan int
	fmt.Println("\nPilih Jenis Investasi:")
	fmt.Println("1. Tabungan Biasa")
	fmt.Println("2. Deposito")
	fmt.Println("3. Reksadana")
	fmt.Println("4. Saham")
	fmt.Println("Pilihan (1-4): ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		fmt.Printf("Dana Rp%.2f dimasukkan ke: Tabungan Biasa\n", investasi)
	case 2:
		fmt.Printf("Dana Rp%.2f dimasukkan ke: Deposito\n", investasi)
	case 3:
		fmt.Printf("Dana Rp%.2f dimasukkan ke: Reksadana\n", investasi)
	case 4:
		fmt.Printf("Dana Rp%.2f dimasukkan ke: Saham\n", investasi)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}
func cariKategori(data tabPemasukan, jumlah int, kategori string) {
	fmt.Println("\nHasil Pencarian (Kategori):")
	ditemukan := false
	for i := 0; i < jumlah; i++ {
		if data[i].kategori == kategori {
			fmt.Printf(" %s | Rp%.2f | %s | %s\n", data[i].nama, data[i].jumlah, data[i].frekuensi, data[i].kategori)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Data tidak ditemukan.")
	}
}
func urutMenaik(data tabPemasukan, jumlah int) {
	var A tabPemasukan
	for i := 0; i < jumlah; i++ {
		A[i] = data[i]
	}
	for i := 0; i < jumlah-1; i++ {
		idx := i
		for j := i + 1; j < jumlah; j++ {
			if A[j].jumlah < A-+[idx].jumlah {
				idx = j
			}
		}
		temp := A[i]       
		A[i] = A[idx]      
		A[idx] = temp      

	}
	fmt.Println("\nUrutan Pemasukan (Rendah ke Tinggi):")
	for i := 0; i < jumlah; i++ {
		fmt.Printf(" %s | Rp%.2f | %s | %s\n", A[i].nama, A[i].jumlah, A[i].frekuensi, A[i].kategori)
	}
}
func urutMenurun(data tabPemasukan, jumlah int) {
	var A tabPemasukan
	for i := 0; i < jumlah; i++ {
		A[i] = data[i]
	}
	for i := 0; i < jumlah-1; i++ {
		idx := i
		for j := i + 1; j < jumlah; j++ {
			if A[j].jumlah > A[idx].jumlah {
				idx = j
			}
		}
		temp := A[i]       
		A[i] = A[idx]      
		A[idx] = temp
	}
	fmt.Println("\nUrutan Pemasukan (Tinggi ke Rendah):")
	for i := 0; i < jumlah; i++ {
		fmt.Printf(" %s | Rp%.2f | %s | %s\n", A[i].nama, A[i].jumlah, A[i].frekuensi, A[i].kategori)
	}
}
func tampilkanSemua(data tabPemasukan, jumlah int) {
	fmt.Println("\nDaftar Pemasukan:")
	for i := 0; i < jumlah; i++ {
		fmt.Printf(" %s | Rp%.2f | %s | %s\n", data[i].nama, data[i].jumlah, data[i].frekuensi, data[i].kategori)
	}
}
func tampilanMenu() int {
	var pilihan int
	fmt.Println("\n======= MENU =======")
	fmt.Println("1. Tambah Pemasukan")
	fmt.Println("2. Tampilkan Ringkasan & Rekomendasi")
	fmt.Println("3. Cari Pemasukan per Kategori")
	fmt.Println("4. Urutkan Pemasukan (Naik)")
	fmt.Println("5. Urutkan Pemasukan (Turun)")
	fmt.Println("6. Tampilkan Semua Data")
	fmt.Println("7. Selesai")
	fmt.Print("Pilih menu (1-7): ")
	fmt.Scan(&pilihan)
	return pilihan
}
func main() {
	var data tabPemasukan
	var jumlah int
	var selesai = false
	var kategori string
	for !selesai {
		switch tampilanMenu() {
		case 1:
			inputPemasukan(&data, &jumlah)
		case 2:
			main, side, pasif := hitungTotal(data, jumlah)
			Ringkasan(main, side, pasif)
			Investasi((main + side + pasif) * 0.30)
		case 3:
			fmt.Println("Kategori (main/side/passive): ")
			fmt.Scan(&kategori)
			cariKategori(data, jumlah, kategori)
		case 4:
			urutMenaik(data, jumlah)
		case 5:
			urutMenurun(data, jumlah)
		case 6:
			tampilkanSemua(data, jumlah)
		case 7:
			selesai = true
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

