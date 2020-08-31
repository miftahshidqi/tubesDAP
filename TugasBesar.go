package main

import (
	"fmt"
	"strings"
)

type Hp struct {
	merk, tipe, tahunkel string
	harga, RAM, ROM, jum int
}

type cus struct {
	nama, tglbeli, merkhp, jenishp string
	jmlhbrg, hargabeli             int
}

const N int = 100

type tab [N]Hp

var data tab

type trans [N]cus

var tran trans

var index int

var icus int

func main() {
	
	index = -1
	icus = -1
	menu()
	
}

func menu() {
	var pil string
	var udahan bool
	
	udahan = false
	for !udahan && index <= N {
		fmt.Println("  -----Selamat datang di counter kami!-----  ")
		fmt.Println("  -----       1. Penambahan           -----  ")
		fmt.Println("  -----       2. Edit                 -----  ")
		fmt.Println("  -----       3. Delete               -----  ")
		fmt.Println("  -----       4. Search and Show      -----  ")
		fmt.Println("  -----       5. Show All Data        -----  ")
		fmt.Println("  -----       6. Sort                 -----  ")
		fmt.Println("  -----       7. Transaksi            -----  ")
		fmt.Println("  -----       8. Show All Transaksi   -----  ")
		fmt.Println("  -----       9. Selesai              -----  ")
		fmt.Println("Masukkan Pilihan anda : ")
		fmt.Scan(&pil)

		switch pil {
		case "1":
			penambahan(&data)
		case "2":
			edit(&data)
		case "3":
			hapus(&data)
		case "4":
			SearchAndShow()
		case "5":
			ShowAllData()
		case "6":
			sort(&data)
		case "7":
			transaksi(&tran, &data)
		case "8":
			ShowTransaksi(tran)
		case "9":
			udahan = true
			fmt.Println	("Terima Kasih")
		}
	}
}

func SearchAndShow() {
	var (
		pil  string
		x    string
		spek string
		s    int
		i    int
	)

	fmt.Println("Tampilkan berdasarkan: ")
	fmt.Println("1. Merk")
	fmt.Println("2. Tipe")
	fmt.Println("3. Spesifikasi")
	fmt.Println("Pilih: ")
	fmt.Scan(&pil)

	switch pil {
	case "1":
		fmt.Println("Masukan Merk: ")
		fmt.Scan(&x)

		for i = 0; i <= index; i++ {
			if strings.ToLower(data[i].merk) == strings.ToLower(x) {
				fmt.Println(i+1,".")
				fmt.Println("Merk: ", data[i].merk)
				fmt.Println("Tipe: ", data[i].tipe)
				fmt.Println("Tahun Keluar: ", data[i].tahunkel)
				fmt.Println("Harga: ", data[i].harga)
				fmt.Println("RAM: ", data[i].RAM)
				fmt.Println("ROM: ", data[i].ROM)
				fmt.Println("Jumlah: ", data[i].jum)
			}
		}
	case "2":
		fmt.Println("Masukan Tipe: ")
		fmt.Scan(&x)

		for i = 0; i <= index; i++ {
			if strings.ToLower(data[i].tipe) == strings.ToLower(x) {
				fmt.Println(i+1,".")
				fmt.Println("Merk: ", data[i].merk)
				fmt.Println("Tipe: ", data[i].tipe)
				fmt.Println("Tahun Keluar: ", data[i].tahunkel)
				fmt.Println("Harga: ", data[i].harga)
				fmt.Println("RAM: ", data[i].RAM)
				fmt.Println("ROM: ", data[i].ROM)
				fmt.Println("Jumlah: ", data[i].jum)
			}
		}
	case "3":
		fmt.Println("1. RAM")
		fmt.Println("2. ROM")
		fmt.Println("Pilih: ")
		fmt.Scan(&spek)

		if spek == "1" {
			fmt.Println("Masukan RAM: ")
			fmt.Scan(&s)

			for i = 0; i <= index; i++ {
				if data[i].RAM == s {
					fmt.Println(i+1,".")
					fmt.Println("Merk: ", data[i].merk)
					fmt.Println("Tipe: ", data[i].tipe)
					fmt.Println("Tahun Keluar: ", data[i].tahunkel)
					fmt.Println("Harga: ", data[i].harga)
					fmt.Println("RAM: ", data[i].RAM)
					fmt.Println("ROM: ", data[i].ROM)
					fmt.Println("Jumlah: ", data[i].jum)
				}
			}
		}
		if spek == "2" {
			fmt.Println("Masukan ROM: ")
			fmt.Scan(&s)

			for i = 0; i <= index; i++ {
				if data[i].ROM == s {
					fmt.Println(i+1,".")
					fmt.Println("Merk: ", data[i].merk)
					fmt.Println("Tipe: ", data[i].tipe)
					fmt.Println("Tahun Keluar: ", data[i].tahunkel)
					fmt.Println("Harga: ", data[i].harga)
					fmt.Println("RAM: ", data[i].RAM)
					fmt.Println("ROM: ", data[i].ROM)
					fmt.Println("Jumlah: ", data[i].jum)
				}
			}
		}
	}
}

func penambahan(data *tab) {
	var (
		x,y string
		i int
	)
	
	fmt.Println("Masukkan merk HP: ")
	fmt.Scan(&x)
	fmt.Println("Masukkan tipe HP: ")
	fmt.Scan(&y)
	
	i = search (*data,x,y)

	if i != -1 {
		fmt.Println	("Hp tersebut sudah ada maka akan ditambahkan jumlahnya")
		data[i].jum++
		fmt.Println("Jumlah berhasil ditambahkan")
	} else {
		index++
		data[index].merk = x
		data[index].tipe = y
		fmt.Println("Masukkan Tahun keluar HP: ")
		fmt.Scan(&data[index].tahunkel)
		fmt.Println("Masukkan Harga HP: ")
		fmt.Scan(&data[index].harga)
		fmt.Println("Masukkan RAM HP: ")
		fmt.Scan(&data[index].RAM)
		fmt.Println("Masukkan ROM HP: ")
		fmt.Scan(&data[index].ROM)
		fmt.Println("Masukkan jumlah HP: ")
		fmt.Scan(&data[index].jum)
	}

}

func search(data tab, m3rk, t1p3 string) int {
	var i int

	i = 0
	for (i <= index) && (data[i].merk != m3rk && data[i].tipe != t1p3) {
		i++
	}

	if data[i].merk == m3rk && data[i].tipe == t1p3 {
		return i
	} else {
		return -1
	}
}
func edit(data *tab) {
	var (
		i                  int
		editmerk, edittipe string
	)

	fmt.Println("Masukan merk hp yang mau di edit : ")
	fmt.Scan(&editmerk)
	fmt.Println("Masukan tipe hp yang mau di edit : ")
	fmt.Scan(&edittipe)

	i = search(*data , editmerk, edittipe)

	if i != -1 {
		fmt.Println("Masukkan Merk HP: ")
		fmt.Scan(&data[i].merk)
		fmt.Println("Masukkan Tipe Hp: ")
		fmt.Scan(&data[i].tipe)
		fmt.Println("Masukkan Tahun keluar HP: ")
		fmt.Scan(&data[i].tahunkel)
		fmt.Println("Masukkan Harga HP: ")
		fmt.Scan(&data[i].harga)
		fmt.Println("Masukkan RAM HP: ")
		fmt.Scan(&data[i].RAM)
		fmt.Println("Masukkan ROM HP: ")
		fmt.Scan(&data[i].ROM)
		fmt.Println("Masukkan jumlah HP: ")
		fmt.Scan(&data[i].jum)
	} else {
		fmt.Println("Maaf data tidak ditemukan")
	}
}

func hapus(data *tab) {
	var (
		hapusmerk, hapustipe string
		i                    int
	)
	fmt.Println("Masukan merk hp yang mau di hapus : ")
	fmt.Scan(&hapusmerk)
	fmt.Println("Masukan tipe hp yang mau di hapus : ")
	fmt.Scan(&hapustipe)

	i = search(*data, hapusmerk, hapustipe)

	if i != -1 {
		data[i].merk = ""
		data[i].tipe = ""
		data[i].tahunkel = ""
		data[i].harga = 0
		data[i].RAM = 0
		data[i].ROM = 0
		data[i].jum = 0
		fmt.Println	("Data ke-",i+1," berhasil dihapus")
		refreshDelete (&*data, i)
	} else {
		fmt.Print("maaf data tidak ditemukan")
	}
}

func sort(data *tab) {
	var (
		i, j, inmin int
		tempA, tempB, tempC           Hp
		pilih                         string
	)
	if (index >= 0) {
		fmt.Println("Sorting Berdasarkan: ")
		fmt.Println("1. Jumlah handphone ")
		fmt.Println("2. Harga Hp ")
		fmt.Println("3. Tahun Keluar Hp ")
		fmt.Println("Pilih: ")
		fmt.Scan(&pilih)

		switch pilih {
		case "1":
			i = 0
			for i < index {
				inmin = i
				j = i + 1 
				for (j <= index)  {
					if (data[j].jum < data[inmin].jum){
						inmin = j
					} 
					j = j + 1
				}
				if i != inmin {
					fmt.Println (data[inmin].jum)
					tempA = data[inmin]
					data[inmin] = data[i]
					data[i] = tempA
				}
				i = i + 1
			}
		case "2":
			i = 0
			for i < index {
				inmin = i
				j = i + 1 
				for (j <= index)  {
					if (data[j].harga < data[inmin].harga){
						inmin = j
					} 
					j = j + 1
				}
				if i != inmin {
					tempB = data[inmin]
					data[inmin] = data[i]
					data[i] = tempB
				}
				i = i + 1
			}
		case "3":
			i = 0
			for i < index {
				inmin = i
				j = i + 1 
				for (j <= index)  {
					if (data[j].tahunkel < data[inmin].tahunkel){
						inmin = j
					} 
					j = j + 1
				}
				if i != inmin {
					tempC = data[inmin]
					data[inmin] = data[i]
					data[i] = tempC
				}
				i = i + 1
			}

		}

		ShowAllData()
	} else {
		fmt.Println	("Tidak ada data")
	}
}

func ShowAllData() {
	var (
		i int
	)
	if index >= 0 {
		for i = 0; i <= index; i++ {
			fmt.Println("Data ke- ", i+1)
			fmt.Println("Merk: ", data[i].merk)
			fmt.Println("Tipe: ", data[i].tipe)
			fmt.Println("Tahun Keluar: ", data[i].tahunkel)
			fmt.Println("Harga: ", data[i].harga)
			fmt.Println("RAM: ", data[i].RAM)
			fmt.Println("ROM: ", data[i].ROM)
			fmt.Println("Jumlah: ", data[i].jum)
			fmt.Println()
		}
	} else {
		fmt.Println("Tidak ada data apapun")
		fmt.Println()
	}
}

func transaksi(dataT *trans, data *tab) {
	var (
		a, b    string
		c, jmlh int
		d, done string
		udahan  bool
		m       int
	)

	udahan = false
	for !udahan {
		SearchAndShow()
		fmt.Println("udah siap beli (y/n) ?  ")
		fmt.Scan(&done)
		if done == "y" {
			udahan = true
		}
	}
	jmlh = 0
	d = "y"

	for d == "y" {
		fmt.Println("merk hp: ")
		fmt.Scan(&a)
		fmt.Println("tipe HP: ")
		fmt.Scan(&b)

		m = 0
		for m <= index && data[m].merk != a && data[m].tipe != b {
			m++
		}
		if data[m].merk == a && data[m].tipe == b {
			icus++
			fmt.Println("Nama Pembeli : ")
			fmt.Scan(&dataT[icus].nama)
			fmt.Println("Tanggal Beli (DD/MM/YYYY): ")
			fmt.Scan(&dataT[icus].tglbeli)
			dataT[icus].merkhp = a
			dataT[icus].jenishp = b
			fmt.Println("Jumlah yang ingin dibeli : ")
			fmt.Scan(&c)
			for c > data[m].jum {
				fmt.Println("maaf barang tidak mencukupi. Ulangi lagi")
				fmt.Println("Jumlah yang ingin dibeli : ")
				fmt.Scan(&c)
			}
			data[m].jum = data[m].jum - c
			if data[m].jum == 0 {
				data[m].merk = ""
				data[m].tipe = ""
				data[m].tahunkel = ""
				data[m].harga = 0
				data[m].RAM = 0
				data[m].ROM = 0
				data[m].jum = 0
				refreshDelete (&*data, m)
			}
			dataT[icus].jmlhbrg = c
			dataT[icus].hargabeli = c * data[m].harga
			jmlh = jmlh + (c * data[m].harga)
		} else {
			fmt.Println("Maaf barang tidak ditemukan")
		}
		fmt.Println("Masih mau beli? (y/n)")
		fmt.Scan(&d)
	}

	fmt.Println("Jumlah harga yang harus dibayar: ", jmlh)
}

func ShowTransaksi (tran trans) {
	var i int
	
	if icus >= 0 {
		for i = 0; i <= icus; i++ {
			fmt.Println("Data ke- ", i+1)
			fmt.Println("Nama Pembeli: ", tran[i].nama)
			fmt.Println("Tanggal Beli: ", tran[i].tglbeli)
			fmt.Println("Merk HP: ", tran[i].merkhp)
			fmt.Println("Jenis HP: ", tran[i].jenishp)
			fmt.Println("Jumlah Barang: ", tran[i].jmlhbrg)
			fmt.Println("Harga: ", tran[i].hargabeli)
			fmt.Println()
		}
	} else {
		fmt.Println	("Tidak ada data transaksi")
	}
}

func refreshDelete (data *tab, i int) {
	var j int
	
	for i < index {
		j = i + 1
		data[i] = data[j]
		i++
	}
	index--
}