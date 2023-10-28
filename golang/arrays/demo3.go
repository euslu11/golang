package arrays

import "fmt"

func Demo3() {
	sayi := [5]int{06, 18, 74, 14, 37}
	fmt.Println(sayi)

	enbuyuk := sayi[0]

	//length = uzunluk
	for i:=0;i<len(sayi);i++{
		if sayi[i]>enbuyuk {
			enbuyuk = sayi[i]
		}
		fmt.Println("en büyük : ",enbuyuk)
	}

}