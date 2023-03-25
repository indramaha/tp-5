package main

import "fmt"

func inputTglPinjam(tgl, bln, thn *int) {
	var tglin, blnin, thnin int
	fmt.Scan(&tglin, &blnin, &thnin)
	*tgl = tglin
	*bln = blnin
	*thn = thnin
}

func kabisat(thn int) bool {
	var hasil bool
	if thn % 4 == 0 {
		hasil = true
	} else {
		hasil = false
	}
	return hasil
}

func getJumlahHari(bln, thn int, jmlHari *int){
	if kabisat(thn) == true {
		if bln == 1{
			*jmlHari = 31
		} else if bln == 2 {
			*jmlHari = 29
		} else if bln == 3 {
			*jmlHari = 31
		} else if bln == 4 {
			*jmlHari = 30
		} else if bln == 5 {
			*jmlHari = 31
		} else if bln == 6 {
			*jmlHari = 30
		} else if bln == 7 {
			*jmlHari = 31
		} else if bln == 8 {
			*jmlHari = 31
		} else if bln == 9 {
			*jmlHari = 30
		} else if bln == 10 {
			*jmlHari = 31
		} else if bln == 1 {
			*jmlHari = 30
		} else if bln == 12 {
			*jmlHari = 31
		}
	} else {
		if bln == 1{
			*jmlHari = 31
		} else if bln == 2 {
			*jmlHari = 28
		} else if bln == 3 {
			*jmlHari = 31
		} else if bln == 4 {
			*jmlHari = 30
		} else if bln == 5 {
			*jmlHari = 31
		} else if bln == 6 {
			*jmlHari = 30
		} else if bln == 7 {
			*jmlHari = 31
		} else if bln == 8 {
			*jmlHari = 31
		} else if bln == 9 {
			*jmlHari = 30
		} else if bln == 10 {
			*jmlHari = 31
		} else if bln == 1 {
			*jmlHari = 30
		} else if bln == 12 {
			*jmlHari = 31
		}
	}
}

func valid(tgl, bln, thn int) bool {
	var hasil bool
	var jmlHari int
	getJumlahHari(bln, thn, &jmlHari)
	if thn > 0 && (bln > 0 && bln <= 12) && (tgl > 0 && tgl <= jmlHari){
		hasil = true
	} else {
		hasil = false
	}
	return hasil
}

func hitungTanggalKembali(tgl,bln,thn int, tglHasil, blnHasil, thnHasil *int){
	var jmlHari int
	var tglkembali, blnKembali, thnKembali int

	tglkembali = tgl + 3
	getJumlahHari(bln,thn, &jmlHari)

	if tglkembali > jmlHari {
		*tglHasil = tglkembali % jmlHari
		blnKembali = bln + 1
		if blnKembali > 12 {
			*blnHasil = blnKembali % 12
			thnKembali = thn + 1
			*thnHasil = thnKembali
		} else if blnKembali <= 12 {
			*blnHasil = blnKembali
			*thnHasil = thn
		}
	} else if tglkembali <= jmlHari {
		*tglHasil = tglkembali
		*blnHasil = bln
		*thnHasil = thn
	}
}

func main(){
	var tgl1, bln1, thn1, tgl2, bln2, thn2 int
	inputTglPinjam(&tgl1,&bln1,&thn1)
	for valid(tgl1,bln1,thn1) {
		hitungTanggalKembali(tgl1, bln1, thn1, &tgl2, &bln2, &thn2)
		fmt.Println(tgl2, bln2, thn2)
		inputTglPinjam(&tgl1,&bln1,&thn1)
	}
	fmt.Println("input tidak valid")
}