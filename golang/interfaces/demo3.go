package interfaces

import "fmt"

func Dogrula(i interface{}) {
	sayi, ok  := i.(int)
	fmt.Println(sayi,ok)
}

func Demo3() {
	var sayi1 interface{} = "Mauro"
	Dogrula(sayi1)

	var sayi2 interface{} = 1905
	Dogrula(sayi2)
}
