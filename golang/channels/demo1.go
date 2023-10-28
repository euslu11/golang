package channels

func CiftSayilar(CiftSayiCn chan int) {
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		toplam = toplam + i
	}
	CiftSayiCn <- toplam
}

func TekSayilar(TekSayiCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam = toplam + i
	}
	TekSayiCn <- toplam
}