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
func tampilanMenu() int {
    var pilihan int
    fmt.Println("-----------------------------------------------------")
    fmt.Println("                     MENU ")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("1. Tambah Pemasukan") //input
    fmt.Println("2. Tampilkan Ringkasan & Rekomendasi") //sequential
    fmt.Println("3. Cari Pemasukan per Kategori") //sequential search
    fmt.Println("4. Urutkan Pemasukan (Naik)") // selectionsort
    fmt.Println("5. Urutkan Pemasukan (Turun)") // selectionsort
    fmt.Println("6. Tampilkan Semua Data") //sequential
    fmt.Println("7. Hapus Pemasukan Terakhir") // hapus aja
    fmt.Println("8. Selesai")
    fmt.Print("Pilih menu (1-8): ")
    fmt.Scan(&pilihan)
    return pilihan
}

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
            totalMain, totalSide, totalPasif := hitungTotalPerBulan(data, nData)
            Ringkasan(totalMain, totalSide, totalPasif)
            Investasi((totalMain + totalSide + totalPasif) * 0.30)
        case 3:
            fmt.Println("Kategori (main/side/passive): ")
            fmt.Scan(&kategori)
            cariKategori(data, nData, kategori)
        case 4:
                urutMenaik(&data, nData)
                tampilkanSemua(data, nData)
        case 5:
                urutMenurun(&data, nData)
                tampilkanSemua(data, nData)
        case 6:
                tampilkanSemua(data, nData)
            
        case 7:
            if hapusData(&data, &nData) {
                fmt.Println("Data terakhir berhasil dihapus")
            }
        case 8:
            selesai = true
        default:
            fmt.Println("Pilihan tidak valid")
        }
    }
}

func inputPemasukan(A *tabPemasukan, n *int) {
    if *n >= NMAKS {
        fmt.Println("Array penuh, tidak bisa menambah pemasukan lagi.")
        return
    }
    var lanjut string = "y"
    var p Pemasukan
    for lanjut == "y" && *n < NMAKS {
        fmt.Println("Nama Pemasukan (gaji/freelance/usaha/royalti): ")
        fmt.Scan(&p.nama)
        fmt.Scanln()
        fmt.Println("Jumlah Pemasukan (Rp): ")
        fmt.Scan(&p.jumlah)
        fmt.Scanln()
        fmt.Println("Jangka Waktu (bulanan/tahunan): ")
        fmt.Scan(&p.frekuensi)
        fmt.Scanln()
        fmt.Println("Kategori (main/side/passive): ")
        fmt.Scan(&p.kategori)
        fmt.Scanln()
        A[*n] = p
        *n++
        fmt.Println("Tambah Pemasukan Lagi? (y/n): ")
        fmt.Scan(&lanjut)
        fmt.Scanln()
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
        bulanan := konversiKeBulanan(A[i].jumlah, A[i].frekuensi)
        if A[i].kategori == "main" {
            totalMain += bulanan
        } else if A[i].kategori == "side" {
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
    pass = 1
    for pass <= n-1 {
        idx = pass - 1
        i = pass
        for i < n {
            if A[i].jumlah < A[idx].jumlah {
                idx = i
            }
            i = i + 1
        }
        temp = A[pass-1]
        A[pass-1] = A[idx]
        A[idx] = temp
        pass = pass + 1
    }
}

func urutMenurun(A *tabPemasukan, n int) {
    var pass, idx, i int
    var temp Pemasukan
    pass = 1
    for pass <= n-1 {
        idx = pass - 1
        i = pass
        for i < n {
            if A[i].jumlah > A[idx].jumlah {
                idx = i
            }
            i = i + 1
        }
        temp = A[pass-1]
        A[pass-1] = A[idx]
        A[idx] = temp
        pass = pass + 1
    }
}

func tampilkanSemua(A tabPemasukan, n int) {
    if n == 0 {
        fmt.Println("Tidak ada data untuk ditampilkan.")
        return
    }
    fmt.Println("Daftar Pemasukan:")
    for i := 0; i < n; i++ {
        fmt.Printf("%s  Rp%.2f  %s  %s\n", A[i].nama, A[i].jumlah, A[i].frekuensi, A[i].kategori)
    }
}

func hapusData(A *tabPemasukan, n *int) bool {
    if *n > 0 {
        *n -= 1
        return true
    }
    fmt.Println("Data tidak bisa dihapus")
    return false
}