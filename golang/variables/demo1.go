package variables

import "fmt"

func Demo1() {

	//string
	var metin string = "MARABA"
	fmt.Println(metin)
	fmt.Println(metin)
	fmt.Println(metin)

	//integer
	var kdv int = 10 
	fmt.Println(kdv)
	fmt.Println(5*kdv)
	
	//boolean
	var durum bool
	var metin1 string = "Euslu"
	var metin2 string = "Euslu11"

	//true-false
	durum = metin1==metin2
	fmt.Println(durum)
}