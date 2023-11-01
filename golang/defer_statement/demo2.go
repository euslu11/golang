package deferstatement

import "fmt"

func Demo2(sayi int32) string {
	defer DeferFunc()
	if sayi%2 == 0 {
		return "Cift sayidir"
	}

	if sayi%2 != 0 {
		return "Tek sayidir"
	}
	return "Belli Degil"
}

func Test() {
	sonuc := Demo2(10)
	fmt.Println(sonuc)
}

func DeferFunc()  {
	fmt.Println("Bitti")
}