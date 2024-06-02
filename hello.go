package main

import "fmt"


type userRegistrasi struct {
	password string
	nama string
}

type dataUserPenjual struct {
	nama string
	password string
	barangJual [99]string
	stokBarangJual [99]int
	hargaBarangJual [99]int
	totalBarangJual int
}

type dataUserPembeli struct {
	nama string
	password string
	barangBeli [99]string
	hargaBarangBeli [99]int
	stokBarangBeli [99]int
	totalBarangBeli int
}

type akun[9999] userRegistrasi
type arrPenjual[9999] dataUserPenjual
type arrPembeli[9999] dataUserPembeli

func main()  {
	var userPenjual arrPenjual
	var userPembeli arrPembeli
	var akunTempPenjual, akunTempPembeli akun
	var x , jumAkunTempPenjual, jumAkunTempPembeli, jumAkunFixPembeli, jumAkunFixPenjual int
	jumAkunTempPembeli = 0
	jumAkunTempPenjual = 0
	jumAkunFixPenjual = 0
	jumAkunFixPembeli = 0

	for x != 6 {
		fmt.Println("==================================")
		fmt.Println("Home:")
		fmt.Printf("1. Registrasi akun penjual\n2. Registrasi akun pembeli\n3. Login Admin\n4. Login akun penjual\n5. Login akun pembeli\n6. Keluar\n")
		fmt.Scan(&x)
		
		if x == 1 {
			registrasiAkun(&akunTempPenjual, &jumAkunTempPenjual)
		} else if x == 2 {
			registrasiAkun(&akunTempPembeli, &jumAkunTempPembeli)
		} else if x == 3 {
			loginAkunAdmin(&akunTempPembeli, &akunTempPenjual,  &userPenjual, &userPembeli,&jumAkunTempPenjual, &jumAkunTempPembeli, &jumAkunFixPenjual, &jumAkunFixPembeli)
		} else if x == 4 {
			loginAkunPenjual(&userPenjual, &jumAkunFixPenjual)
		} else if x == 5 {
            loginAkunPembeli(&userPembeli, &userPenjual, &jumAkunFixPembeli, &jumAkunFixPenjual)
        }
	}
}

// Registrasi akun
func registrasiAkun(akunTemp *akun, jumAkunTemp *int)  {
	fmt.Println("==================================")
	fmt.Println("Masukkan Username dan password!")
	fmt.Scan(&akunTemp[*jumAkunTemp].nama, &akunTemp[*jumAkunTemp].password)
	*jumAkunTemp++
}

// Masuk ke akun admin
func loginAkunAdmin(akunTempPembeli, akunTempPenjual *akun, userPenjual *arrPenjual, userPembeli *arrPembeli, jumAkunTempPenjual, jumAkunTempPembeli, jumAkunFixPenjual, jumAkunFixPembeli *int)  {
	var indeksTolak, indeksTerima int
	var n int
	fmt.Println("==================================")
	fmt.Println("Pilih Opsi:")
	fmt.Printf("1. Melihat akun registrasi penjual\n2. Melihat akun registrasi pembeli\n3. Melihat akun yang sudah terdaftar\n4. Logout Admin\n")
	fmt.Scan(&n)

	for n != 4 {
		if n == 1 {
			fmt.Println("==================================")
			fmt.Println("List akun registrasi Penjual:")
			for i := 0; i < *jumAkunTempPenjual; i++ {
				fmt.Println(i + 1, (*akunTempPenjual)[i].nama, (*akunTempPenjual)[i].password)
			}
			fmt.Printf("Pilih indeks keberapa akun yang diaccept(kalau tidak ada, input '-1'): ")
			fmt.Scan(&indeksTerima)
			indeksTerima--
			if indeksTerima >=0 {
				(*userPenjual)[*jumAkunFixPenjual].nama = (*akunTempPenjual)[indeksTerima].nama
				(*userPenjual)[*jumAkunFixPenjual].password = (*akunTempPenjual)[indeksTerima].password
				(*userPenjual)[*jumAkunFixPenjual].totalBarangJual = 0
				*jumAkunTempPenjual--
				*jumAkunFixPenjual++
				hapusAkun(akunTempPenjual, jumAkunTempPenjual, &indeksTerima)
			}
			fmt.Printf("Pilih indeks keberapa akun yang direject(kalau tidak ada, input '-1'): ")
			fmt.Scan(&indeksTolak)
			indeksTolak--
			if indeksTolak >=0 {
				*jumAkunTempPenjual--
				hapusAkun(akunTempPenjual, jumAkunTempPenjual, &indeksTolak)
			}
		} else if n == 2 {
			fmt.Println("==================================")
			fmt.Println("List akun registrasi Pembeli:")
			for i := 0; i < *jumAkunTempPembeli; i++ {
				fmt.Println(i + 1, (*akunTempPembeli)[i].nama, (*akunTempPembeli)[i].password)
			}
			fmt.Printf("Pilih indeks keberapa akun yang diaccept(kalau tidak ada, input '-1'): ")
			fmt.Scan(&indeksTerima)
			indeksTerima--
			if indeksTerima >=0 {
				(*userPembeli)[*jumAkunFixPembeli].nama = (*akunTempPembeli)[indeksTerima].nama
				(*userPembeli)[*jumAkunFixPembeli].password = (*akunTempPembeli)[indeksTerima].password
				*jumAkunTempPembeli--
				*jumAkunFixPembeli++
				hapusAkun(akunTempPembeli, jumAkunTempPembeli, &indeksTerima)
			}
			fmt.Printf("Pilih indeks keberapa akun yang direject(kalau tidak ada, input '-1'): ")
			fmt.Scan(&indeksTolak)
			indeksTolak--
			if indeksTolak >=0 {
				*jumAkunTempPembeli--
				hapusAkun(akunTempPembeli, jumAkunTempPembeli, &indeksTolak)
			}
		} else if n == 3 {
			fmt.Println("==================================")
			fmt.Println("List akun Penjual:")
			for d := 0; d < *jumAkunFixPenjual; d++ {
				fmt.Println(d + 1, (*userPenjual)[d].nama, (*userPenjual)[d].password)
			}
			fmt.Println("==================================")
			fmt.Println("List akun Pembeli:")
			for f := 0; f < *jumAkunFixPembeli; f++ {
				fmt.Println(f + 1, (*userPembeli)[f].nama, (*userPembeli)[f].password)
			}
		}
		fmt.Println("==================================")
		fmt.Println("Pilih Opsi:")
		fmt.Printf("1. Melihat akun registrasi penjual\n2. Melihat akun registrasi pembeli\n3. Melihat akun yang sudah terdaftar\n4. Logout Admin\n")
		fmt.Scan(&n)
	}

}

// Menghapus akun jika tidak disetujui admin
func hapusAkun(akunTemp *akun, jumAkunTemp, indeks *int)  {
	for i := *indeks; i < *jumAkunTemp; i++ {
		(*akunTemp)[i] = (*akunTemp)[i + 1]
	}
}

// Login akun penjual
func loginAkunPenjual(userPenjual *arrPenjual, jumAkunFixPenjual *int)  {
	var username, password string
	var opsi, urutanPenjual int
	fmt.Println("==================================")
	fmt.Println("Masukkan username dan password anda!")
	fmt.Scan(&username, &password)
	urutanPenjual = findUsernamePasswordPenjual(*userPenjual, *jumAkunFixPenjual, username, password)
	if urutanPenjual >= 0 {
		fmt.Println("==================================")
		fmt.Println("Anda berhasil login")
		fmt.Println("Apa yang ingin anda lakukan:")
		fmt.Printf("1. Melihat barang\n2. Menambahkan barang\n3. Mengubah barang\n4. Menghapus barang\n0. Keluar sebagai penjual\n")
		fmt.Scan(&opsi)
		for opsi != 0 {
			if opsi == 1 {
				lihatBarangJualan(userPenjual, urutanPenjual)
			} else if opsi == 2 {
				nambahBarang(userPenjual, urutanPenjual)
			} else if opsi == 3 {
				ubahBarang(userPenjual, urutanPenjual)
			} else if opsi == 4 {
				hapusBarang(userPenjual, urutanPenjual)
			}
			fmt.Println("==================================")
			fmt.Println("Apa yang ingin anda lakukan:")
			fmt.Printf("1. Melihat barang\n2. Menambahkan barang\n3. Mengubah barang\n4. Menghapus barang\n0. Keluar sebagai penjual\n")
			fmt.Scan(&opsi)
		}
		
	} else {
		fmt.Println("==================================")
		fmt.Println("Username dan password yang anda masukkan salah!")
	}
}

// Memverifikasi username dan password.
func findUsernamePasswordPenjual(akunFix arrPenjual, jumAkunFix int, username, password string) int {
	for i := 0; i < jumAkunFix; i++ {
		if username == akunFix[i].nama && password == akunFix[i].password {
			return i
		}
	}
	return -1
}

func findUsernamePasswordPembeli(akunFix arrPembeli, jumAkunFix int, username, password string) int {
	for i := 0; i < jumAkunFix; i++ {
		if username == akunFix[i].nama && password == akunFix[i].password {
			return i
		}
	}
	return -1
}

// Menambah barang jualan
func nambahBarang(user *arrPenjual, urutan int)  {
	var namaBarang string
	var hargaBarang, stok int
	fmt.Println("==================================")
	fmt.Println("Masukkan nama, harga dan stok barang yang ingin ditambahkan")
	fmt.Println("(jika merasa cukup, input 'cukup')")
	fmt.Scan(&namaBarang)

	for namaBarang != "cukup" {
		fmt.Scan(&hargaBarang)
		fmt.Scan(&stok)
		(*user)[urutan].barangJual[(*user)[urutan].totalBarangJual] = namaBarang
		(*user)[urutan].hargaBarangJual[(*user)[urutan].totalBarangJual] = hargaBarang
		(*user)[urutan].stokBarangJual[(*user)[urutan].totalBarangJual] = stok
		(*user)[urutan].totalBarangJual++
		fmt.Scan(&namaBarang)
	}
}

// Melihat barang jualan
func lihatBarangJualan(barang *arrPenjual, urutan int)  {
	fmt.Println("==================================")
	fmt.Println("Berikut merupakan list barang yang ada di toko:")
	fmt.Println("Nama Barang", "Harga Barang", "Stok Barang")
		for i := 0; i < (*barang)[urutan].totalBarangJual; i++ {
			fmt.Println(i + 1, (*barang)[urutan].barangJual[i], (*barang)[urutan].hargaBarangJual[i], (*barang)[urutan].stokBarangJual[i])
		}
}

// Mengubah barang jualan
func ubahBarang(user *arrPenjual, urutan int)  {
	var namaBarang, cekNamaBarang string
	var hargaBarang, indeks, stok int
	fmt.Println("Masukkan nama barang yang ingin diubah")
	fmt.Scan(&cekNamaBarang)
	indeks = findIndeksBarangJual(user, cekNamaBarang, urutan)
	if indeks != -1 {
		fmt.Println("Masukkan nama, harga dan stok barang pengganti")
		fmt.Scan(&namaBarang, &hargaBarang, &stok)
		(*user)[urutan].barangJual[indeks] = namaBarang
		(*user)[urutan].hargaBarangJual[indeks] = hargaBarang
		(*user)[urutan].stokBarangJual[indeks] = stok
	} else {
		fmt.Println("Maaf barang yang ingin anda ubah tidak ada")
	}
}

// Menemukan indeks barang tertentu
func findIndeksBarangJual(user *arrPenjual, namaBarang string, urutan int) int {
	var idx int
	idx = -1
	for i := 0; i < (*user)[urutan].totalBarangJual; i++ {
		if (*user)[urutan].barangJual[i] == namaBarang {
			idx = i
		}
	}
	return idx
}

// Menghapus barang jualan
func hapusBarang(user *arrPenjual, urutan int)  {
	var indekss int
	var namaBarang string
	fmt.Println("Masukkan nama barang yang ingin anda hapus")
	fmt.Scan(&namaBarang)
	indekss = findIndeksBarangJual(user, namaBarang, urutan)
	if indekss != -1 {
		deleteBarang(user, namaBarang, urutan, indekss)
	} else {
		fmt.Println("Maaf barang yang ingin anda hapus tidak ada")
	}
}

// delete barang sesuai indeks
func deleteBarang(user *arrPenjual, namaBarang string, urutan, indeks int)  {
	(*user)[urutan].totalBarangJual --
		for d := indeks; d < (*user)[urutan].totalBarangJual; d++ {
			(*user)[urutan].barangJual[d] = (*user)[urutan].barangJual[d + 1]
			(*user)[urutan].hargaBarangJual[d] = (*user)[urutan].hargaBarangJual[d + 1]
			(*user)[urutan].stokBarangJual[d] = (*user)[urutan].stokBarangJual[d + 1]
		}
}

func loginAkunPembeli(userPembeli *arrPembeli, userPenjual *arrPenjual, jumAkunFixPembeli, jumAkunFixPenjual *int)  {
    var username, password string
	var opsi, urutanPembeli, toko int
	toko = 1
	fmt.Println("==================================")
	fmt.Println("Masukkan username dan password anda!")
	fmt.Scan(&username, &password)
	urutanPembeli = findUsernamePasswordPembeli(*userPembeli, *jumAkunFixPembeli, username, password) 
	if urutanPembeli >= 0{
		for toko >= 0 {
			opsi = 1
			fmt.Println("==================================")
			fmt.Println("Anda berhasil login")
			fmt.Println("Pilih toko yang ingin anda kunjungi, input '0' jika ingin logout)")
			for i := 0; i < *jumAkunFixPenjual; i++ {
				fmt.Println(i + 1, "Toko", userPenjual[i].nama)
			}
			fmt.Scan(&toko)
			toko--
			for opsi > 0 && toko >= 0 {
				fmt.Println("==================================")
				fmt.Println("Apa yang ingin anda lakukan:")
				fmt.Printf("1. Lihat Keranjang\n2. Membeli barang\n3. Cetak transaksi\n0. Keluar toko\n")
				fmt.Scan(&opsi)
				if opsi == 1 {
                	lihatBarangBeli(userPembeli, urutanPembeli)
          	 	 } else if opsi == 2 {
					lihatBarangJualan(userPenjual, toko)
					beliBarang(userPembeli, userPenjual, urutanPembeli, toko)
          	 	 } else if opsi == 3 {
					cetakTransaksi(userPembeli, urutanPembeli)
				 }
			}
		}
    } else {
		fmt.Println("==================================")
		fmt.Println("Anda memasukkan username dan password yang salah")
	}
}

func beliBarang(userPembeli *arrPembeli, userPenjual *arrPenjual, urutanPembeli, toko int)  {
	var namaBarang string
	var index , stok, stokBarangSetelahDibeli int
	fmt.Println("==================================")
	fmt.Println("Masukkan nama dan jumlah barang yang ingin anda beli")
	fmt.Println("(jika merasa cukup, input 'cukup')")
	fmt.Scan(&namaBarang)
	for namaBarang != "cukup" {
		index = -1
		fmt.Scan(&stok)
		index = findIndeksBarangJual(userPenjual, namaBarang, toko)
		if index >= 0 {
			stokBarangSetelahDibeli = (*userPenjual)[toko].stokBarangJual[index] - stok
			if stokBarangSetelahDibeli >= 0 {
				(*userPembeli)[urutanPembeli].barangBeli[(*userPembeli)[urutanPembeli].totalBarangBeli] = (*userPenjual)[toko].barangJual[index]
				(*userPembeli)[urutanPembeli].hargaBarangBeli[(*userPembeli)[urutanPembeli].totalBarangBeli] = (*userPenjual)[toko].hargaBarangJual[index]
				(*userPembeli)[urutanPembeli].stokBarangBeli[(*userPembeli)[urutanPembeli].totalBarangBeli] = stok
				(*userPenjual)[toko].stokBarangJual[index] = stokBarangSetelahDibeli
				(*userPembeli)[urutanPembeli].totalBarangBeli++
			} else {
				fmt.Println("Maaf stok barang yang ingin anda beli tidak cukup")
			}

		} else {
			fmt.Println("Maaf, barang yang ingin anda beli tidak ada.")
		}
		fmt.Scan(&namaBarang)
	}
}

func cetakTransaksi(user *arrPembeli, urutan int)  {
	var total int
	fmt.Println("==================================")
	fmt.Println("Struk pembelian: ")
	for i := 0; i < (*user)[urutan].totalBarangBeli; i++ {
		fmt.Println(i + 1, (*user)[urutan].barangBeli[i], ("berjumlah:"), (*user)[urutan].stokBarangBeli[i], ("harga:"), (*user)[urutan].hargaBarangBeli[i])
		total += (*user)[urutan].hargaBarangBeli[i] * (*user)[urutan].stokBarangBeli[i]
	}
	fmt.Println("Total Belanja: ", total)
}

func lihatBarangBeli(user *arrPembeli, urutan int)  {
	fmt.Println("==================================")
	fmt.Println("Berikut merupakan list barang yang anda beli:")
		for i := 0; i < (*user)[urutan].totalBarangBeli; i++ {
			fmt.Println(i + 1, (*user)[urutan].barangBeli[i], (*user)[urutan].hargaBarangBeli[i], (*user)[urutan].stokBarangBeli[i])
		}
}