package conditionals

import "fmt"

func Demo1() {

	var hesap float64 = 1000
	var cekpara float64 = 900

	if cekpara > hesap {
		fmt.Println("Yetersiz Bakiye")
	}

	if cekpara <= hesap {
		fmt.Println("Para Hazırlanıyor")
		hesap = hesap - cekpara
	}

	fmt.Println("Kalan Bakiye : "+ fmt.Sprintf("%v",hesap))
}	