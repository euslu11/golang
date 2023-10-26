package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekpara float64 = 900

	//if-else 
	if cekpara > hesap {
		fmt.Println("Yetersiz Bakiye")
	}  else if cekpara==hesap{
		fmt.Println("Paranız Hazırlanıyor")
	}else {
		fmt.Println("Paranız Hazırlanıyor")
		}

	
}