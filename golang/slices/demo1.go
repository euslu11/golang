package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 4)

	fmt.Println(isimler)

	isimler[0] = "Emirhan"
	isimler[1] = "Basri"
	isimler[2] = "Sakine"
	isimler[3] = "Emre"
	//append = eklemek
	isimler = append(isimler ,"Nazlı")
	isimler = append(isimler ,"Oralet")
	fmt.Println(isimler)
	fmt.Println(len(isimler))

}