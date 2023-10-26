package loops

import "fmt"

//asal sayi mi degil mi
func Demo4() {
	sayi := 0
	fmt.Println("Bir Sayi Giriniz")
	fmt.Scanln(&sayi)
	
	asal:=true

	for i:=2;i<sayi;i++ {
		if sayi % i == 0{
			asal = false
		}
	}

	if  asal == true {
		fmt.Println(sayi,"Degil")
	}else {
		fmt.Println(sayi,"Asal")
	}
}