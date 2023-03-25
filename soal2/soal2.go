package main

import "fmt"

func membeliKain(uangAwal float64, tUang, tPengeluaran *float64) {
	*tUang = uangAwal - (uangAwal / 4)
	*tPengeluaran += (uangAwal / 4)
}

func membeliBenangJahit(uangAwal float64, tUang, tPengeluaran *float64) {
	*tUang = uangAwal - (uangAwal / 20)
	*tPengeluaran += (uangAwal / 20)
}

func membuatPolaBaju(uangAwal float64, tUang, tPengeluaran *float64) {
	*tUang = uangAwal - (uangAwal / 200)
	*tPengeluaran += (uangAwal / 200)
}

func menjahitBaju(uangAwal float64, tUang, tPengeluaran *float64) {
	*tUang = uangAwal - (uangAwal / 200)
	*tPengeluaran += (uangAwal / 200)
}

func mengemasBaju(uangAwal float64, tUang, tPengeluaran *float64) {
	*tUang = uangAwal - (3*uangAwal / 200)
	*tPengeluaran += (3*uangAwal / 200)
}

func mendistribusikan(uangAwal float64, tUang, tPemasukan, tPengeluaran *float64) {
	*tPemasukan += (uangAwal/2)
	*tPengeluaran += (3*uangAwal / 200)
	*tUang = uangAwal - (3*uangAwal / 200) + *tPemasukan
}

func main() {
	var modal, totalUang, totalPengluaran, totalPemasukan float64
	fmt.Scan(&modal)
	membeliKain(modal, &totalUang, &totalPengluaran)
	membeliBenangJahit(modal, &totalUang, &totalPengluaran)
	for i := 1; i <=2; i++{
		membuatPolaBaju(modal, &totalUang, &totalPengluaran)
	}
	for i := 1; i <=4; i++{
		menjahitBaju(modal, &totalUang, &totalPengluaran)
	}
	for i := 1; i <=2; i++{
		mengemasBaju(modal, &totalUang, &totalPengluaran)
	}
	mendistribusikan(modal, &totalUang, &totalPemasukan, &totalPengluaran)
	fmt.Println(int(totalPengluaran))
	fmt.Println(int(totalPemasukan))
	fmt.Println(int(totalUang))
}