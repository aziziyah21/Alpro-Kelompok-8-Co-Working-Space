package main

import "fmt"

type CoWorkingSpace struct {
	Nama      string
	Lokasi    string
	Fasilitas []string
	Harga     int
	Rating    float64
	Ulasan    string
}

var spaces []CoWorkingSpace

func main() {
	for {
		fmt.Println("\nMENU COWORKING SPACE")
		fmt.Println("1. Menambahkan Data")
		fmt.Println("2. Mengubah Data")
		fmt.Println("3. Menghapus Data")
		fmt.Println("4. Mengubah Fasilitas")
		fmt.Println("5. Mengubah Rating")
		fmt.Println("6. Tambah/Ubah Ulasan")
		fmt.Println("7. Menghapus Ulasan")
		fmt.Println("8. Tampilkan Semua")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")

		var pilih int
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahData()
		case 2:
			ubahData()
		case 3:
			hapusData()
		case 4:
			ubahFasilitas()
		case 5:
			ubahRating()
		case 6:
			tambahUlasan()
		case 7:
			hapusUlasan()
		case 8:
			tampilkanSemua()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tampilkanSemua() {
	if len(spaces) == 0 {
		fmt.Println("Belum ada data.")
		return
	}
	for i, s := range spaces {
		fmt.Printf("Data #%d\n", i+1)
		fmt.Printf("Nama    : %s\n", s.Nama)
		fmt.Printf("Lokasi  : %s\n", s.Lokasi)
		fmt.Printf("Harga   : Rp %d\n", s.Harga)
		fmt.Printf("Rating  : ⭐ %.1f\n", s.Rating)
		fmt.Printf("Fasilitas:\n")
		for _, f := range s.Fasilitas {
			fmt.Printf(" - %s\n", f)
		}
		fmt.Printf("Ulasan  : %s\n", s.Ulasan)
		fmt.Println("------------------------------")
	}
}

func tambahData() {
	var nama, lokasi, fasilitas1, fasilitas2, fasilitas3, ulasan string
	var harga int
	var rating float64

	fmt.Print("Nama: ")
	fmt.Scan(&nama)
	fmt.Print("Lokasi: ")
	fmt.Scan(&lokasi)
	fmt.Print("Fasilitas 1: ")
	fmt.Scan(&fasilitas1)
	fmt.Print("Fasilitas 2: ")
	fmt.Scan(&fasilitas2)
	fmt.Print("Fasilitas 3: ")
	fmt.Scan(&fasilitas3)
	fmt.Print("Harga: ")
	fmt.Scan(&harga)
	fmt.Print("Rating: ")
	fmt.Scan(&rating)
	fmt.Print("Ulasan (1 kata): ")
	fmt.Scan(&ulasan)

	spaces = append(spaces, CoWorkingSpace{
		Nama:      nama,
		Lokasi:    lokasi,
		Fasilitas: []string{fasilitas1, fasilitas2, fasilitas3},
		Harga:     harga,
		Rating:    rating,
		Ulasan:    ulasan,
	})

	fmt.Println("Data berhasil ditambahkan.")
	tampilkanSemua()
}

func ubahData() {
	tampilkanSemua()
	fmt.Print("Pilih nomor yang ingin diubah: ")
	var n int
	fmt.Scan(&n)
	if n < 1 || n > len(spaces) {
		fmt.Println("Data tidak ditemukan.")
		return
	}
	n--

	var nama, lokasi, fasilitas1, fasilitas2, fasilitas3 string
	var harga int
	var rating float64

	fmt.Print("Nama baru: ")
	fmt.Scan(&nama)
	fmt.Print("Lokasi baru: ")
	fmt.Scan(&lokasi)
	fmt.Print("Fasilitas 1: ")
	fmt.Scan(&fasilitas1)
	fmt.Print("Fasilitas 2: ")
	fmt.Scan(&fasilitas2)
	fmt.Print("Fasilitas 3: ")
	fmt.Scan(&fasilitas3)
	fmt.Print("Harga baru: ")
	fmt.Scan(&harga)
	fmt.Print("Rating baru: ")
	fmt.Scan(&rating)

	spaces[n].Nama = nama
	spaces[n].Lokasi = lokasi
	spaces[n].Fasilitas = []string{fasilitas1, fasilitas2, fasilitas3}
	spaces[n].Harga = harga
	spaces[n].Rating = rating

	fmt.Println("Data berhasil diubah.")
	tampilkanSemua()
}

func ubahFasilitas() {
	tampilkanSemua()
	fmt.Print("Pilih nomor untuk mengubah fasilitas: ")
	var n int
	fmt.Scan(&n)
	if n < 1 || n > len(spaces) {
		fmt.Println("Data tidak ditemukan.")
		return
	}
	n--

	var f1, f2, f3 string
	fmt.Print("Fasilitas 1 baru: ")
	fmt.Scan(&f1)
	fmt.Print("Fasilitas 2 baru: ")
	fmt.Scan(&f2)
	fmt.Print("Fasilitas 3 baru: ")
	fmt.Scan(&f3)

	spaces[n].Fasilitas = []string{f1, f2, f3}
	fmt.Println("Fasilitas berhasil diubah.")
}

func ubahRating() {
	tampilkanSemua()
	fmt.Print("Pilih nomor untuk mengubah rating: ")
	var n int
	fmt.Scan(&n)
	if n < 1 || n > len(spaces) {
		fmt.Println("Data tidak ditemukan.")
		return
	}
	n--

	var rating float64
	fmt.Print("Masukkan rating baru: ")
	fmt.Scan(&rating)

	spaces[n].Rating = rating
	fmt.Println("Rating berhasil diubah.")
}

func hapusData() {
	tampilkanSemua()
	fmt.Print("Pilih nomor yang ingin dihapus: ")
	var n int
	fmt.Scan(&n)
	if n < 1 || n > len(spaces) {
		fmt.Println("Data tidak ditemukan.")
		return
	}
	n--
	spaces = append(spaces[:n], spaces[n+1:]...)
	fmt.Println("Data berhasil dihapus.")
	tampilkanSemua()
}

func tambahUlasan() {
	tampilkanSemua()
	fmt.Print("Pilih nomor untuk tambah/ubah ulasan: ")
	var n int
	fmt.Scan(&n)
	if n < 1 || n > len(spaces) {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	var ulasan string
	fmt.Print("Masukkan ulasan baru (1 kata): ")
	fmt.Scan(&ulasan)
	spaces[n-1].Ulasan = ulasan
	fmt.Println("Ulasan berhasil diperbarui.")
}

func hapusUlasan() {
	tampilkanSemua()
	fmt.Print("Pilih nomor untuk hapus ulasan: ")
	var n int
	fmt.Scan(&n)
	if n < 1 || n > len(spaces) {
		fmt.Println("Data tidak ditemukan.")
		return
	}
	spaces[n-1].Ulasan = ""
	fmt.Println("Ulasan berhasil dihapus.")
}
