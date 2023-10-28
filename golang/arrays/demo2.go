package arrays

import "fmt"

func Demo2() {
	var sehirler [5]string

	sehirler[0] = "Ankara"
	sehirler[1] = "Çankırı"
	sehirler[2] = "Bartın"
	sehirler[3] = "Bolu"
	sehirler[4] = "Zonguldak"
	fmt.Println(sehirler)

	for i:=0 ; i<=5 ; i++{
		fmt.Println(sehirler[i])
	}


}
