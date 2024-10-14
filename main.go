package main

import "fmt"

func main() {
	hargaJual := 150000.0
	hargaBeli := 100000.0
	biayaOperasional := 1000.0
	diskon := 15.0
	jumlahTerjual := 100

	hargaJualAfterDiskon := int(hargaJual) - (int(hargaJual*diskon) / 100)

	fmt.Println("Harga Jual setelah Diskon", hargaJualAfterDiskon)

	totalPendapatan := int(hargaJual) * jumlahTerjual

	fmt.Println("Total Pendapatan", totalPendapatan)

	totalBiaya := int((hargaBeli + biayaOperasional)) * jumlahTerjual

	fmt.Println("Total Biaya", totalBiaya)

	totalKeuntungan := hargaJualAfterDiskon*int(jumlahTerjual) - totalBiaya

	fmt.Println("Total Keuntungan", totalKeuntungan)
}
