package conditionals

import "fmt"

func Workshop1() {

	// üç adet değişken (int) en büyük olanı ekrana yazdıran program.

	var sayi1, sayi2, sayi3 int = 15, 20, 25

	if sayi1 > sayi2 && sayi1>sayi3 {
		fmt.Println("En büyük sayı : "+ fmt.Sprintf("%v",sayi1))
	} else if sayi2 > sayi1 && sayi2>sayi3 {
		fmt.Println("En büyük sayı : "+ fmt.Sprintf("%v",sayi2))
	} else {
		fmt.Println("En büyük sayı : "+ fmt.Sprintf("%v",sayi3))
	}
}