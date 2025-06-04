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
func main() {
	var data tabPemasukan
	var nData int
	var selesai = false
	var kategori string
	for !selesai {
		switch tampilanMenu() {
		case 1:
			inputPemasukan(&data, &nData)
		case 2:
			totalMain, totalSide, totalPasif := hitungTotal(data, nData)
			Ringkasan(totalMain, totalSide, totalPasif)
			Investasi((totalMain + totalSide + totalPasif) * 0.30)
		case 3:
			fmt.Println("Kategori (main/side/passive): ")
			fmt.Scan(&kategori)
			cariKategori(data, nData, kategori)
		case 4:
			urutMenaik(data, nData)
		case 5:
			urutMenurun(data, nData)
		case 6:
			tampilkanSemua(data, nData)
		case 7:
			selesai = true
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
func inputPemasukan(A *tabPemasukan, n *int) {
	var lanjut string 
	var p Pemasukan
	for lanjut == "y" && *n < NMAKS {
		fmt.Println("Nama Pemasukan (gaji/freelance/usaha/royalti): ")
		fmt.Scan(&p.nama)
		fmt.Println("Jumlah Pemasukan (Rp): ")
		fmt.Scan(&p.jumlah)
		fmt.Println("Jangka Waktu (bulanan/tahunan): ")
		fmt.Scan(&p.frekuensi)
		fmt.Println("Kategori (main/side/passive): ")
		fmt.Scan(&p.kategori)
		A[*n] = p
		*n++
		fmt.Println("Tambah Pemasukan Lagi? (y/n): ")
		fmt.Scan(&lanjut)
	}
}
func konversiKeBulanan(jumlah float64, frekuensi string) float64 {
	if frekuensi == "tahunan" {
		return jumlah / 12
	}
	return jumlah
}
func hitungTotalPerBulan(A tabPemasukan, n int) (float64, float64, float64) {
	var totalMain, totalSide, totalPasif float64
	for i := 0; i < n; i++ {
		bulanan := konversiKeBulanan(A[i].jumlah, A[i].frekuensi) //SEMUA DI KONVERSIKAN KE BULANAN JADI TOTALNYA BULANAN
		if A[i].kategori == "main" {
			totalMain += bulanan
		} else if A[i].kategori == "side" { //120000000/12 12jtnya output
			totalSide += bulanan
		} else if A[i].kategori == "passive" {
			totalPasif += bulanan
		}
	}
	return totalMain, totalSide, totalPasif
}
func Ringkasan(totalMain, totalSide, totalPasif float64) {
	totalBulanan := totalMain + totalSide + totalPasif
	totalTahunan := totalBulanan * 12
	fmt.Println("-----------------------------------------------------")
	fmt.Println("                   RINGKASAN ")
	fmt.Println("-----------------------------------------------------")
	fmt.Printf("Main    : Rp%.2f/bulan\n", totalMain)
	fmt.Printf("Side    : Rp%.2f/bulan\n", totalSide)
	fmt.Printf("Pasif   : Rp%.2f/bulan\n", totalPasif)
	fmt.Printf("TOTAL   : Rp%.2f/bulan - Rp%.2f/tahun\n", totalBulanan, totalTahunan)

	fmt.Println("-----------------------------------------------------")
	fmt.Println("                   REKOMENDASI ")
	fmt.Println("-----------------------------------------------------")
	fmt.Printf("Tabungan 10%% : Rp%.2f\n", totalBulanan*0.10)
	fmt.Printf("Investasi 30%%: Rp%.2f\n", totalBulanan*0.30)
	fmt.Printf("Kebutuhan 60%%: Rp%.2f\n", totalBulanan*0.60)
}
func Investasi(investasi float64) {
	var pilihan int
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Pilih Jenis Investasi:")
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
func cariKategori(A tabPemasukan, n int, X string) {
	fmt.Println("Hasil Pencarian (Kategori):")
	ditemukan := false
	for i := 0; i < n; i++ {
		if A[i].kategori == X {
			fmt.Printf("%s  Rp%.2f  %s  %s\n", A[i].nama, A[i].jumlah, A[i].frekuensi, A[i].kategori)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Data tidak ditemukan.")
	}
}
func urutMenaik(A *tabPemasukan, n int) {
	var pass, idx, i int
	var temp Pemasukan
	pass = 1 // array data n = 3 ==> 100000000, 20000000, 30000000 //PASS = 2
	for pass <= n-1 { // PASS = 3 ==> FALSE STOP!!!!!
		idx = pass - 1 // idx = 0 // IDX = 1
		i = pass // i = 1 // I = 2
		for i < n { // 1 < 3 true do // 2 < 3 TRUE DO
			if A[i].jumlah < A[idx].jumlah { //true // I = 2 FALSE
				idx = i // idx = 1 SEKARANG // SKIP
			}
			i = i + 1 // 2 SEKARANG // I = 3
		}
		// Tukar elemen
		temp = A[pass-1] //100000000 // 100000000
		A[pass-1] = A[idx] // TUKER 100000000 JADI KE 1 20000000 JADI KE 0 SETELAH DI TUKER // 100000000 TUKER JADI ==> 100000000
		A[idx] = temp // 20000000, 100000000, 300000000 ==> 0,1,2 TEMP = 100000000

		pass = pass + 1 // 2 SEKARANG // 3
	}
}
func urutMenurun(A *tabPemasukan, n int) {
	var pass, idx, i int // 100000000, 20000000, 30000000 n = 3
	var temp Pemasukan
	pass = 1 // PASS = 2
	for pass <= n-1 { // 1 <= 2 true // PASS <= 2 TRUE
		idx = pass - 1 // idx = 0 // IDX = 1
		i = pass // i = 1 // I = 2
		for i < n { // i < 3 true // 2 < 3 TRUE
			if A[i].jumlah > A[idx].jumlah { // 20000000 !< 100000000 jadi false // 300000000 > 20000000 TRUE
				idx = i // pass 1 skip ini // IDX = 2
			}
			i = i + 1 // i = 2 // I = 3
		}
		// Tukar elemen
		temp = A[pass-1] // temp = 100000000 // TEMP = 20000000
		A[pass-1] = A[idx] // 100000000 = idx ke 0 ttp karena false // 20000000 = A[IDX] // IDX = 2 // JADI 20000000 JADI ARRAY KE 2
		A[idx] = temp // 100000000 = 100000000 ==> 100000000, 20000000, 3000000 // 300000000, 100000000, 20000000 MENURUN TRUE!!!!!!!

		pass = pass + 1 // pass = 2 // PASS = 3 ==> STOP KRNA !<= 2
	}
}

func tampilkanSemua(A tabPemasukan, n int) {
	fmt.Println("Daftar Pemasukan:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s  Rp%.2f  %s  %s\n", A[i].nama, A[i].jumlah, A[i].frekuensi, A[i].kategori)
	}
}
func tampilanMenu() int {
	var pilihan int
	fmt.Println("-----------------------------------------------------")
	fmt.Println("                     MENU ")
	fmt.Println("-----------------------------------------------------")
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
