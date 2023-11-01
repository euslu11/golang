package error_handling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string,error) {
	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasi sayi giriniz")
	}
	return "Bildiniz" , nil 
}

func Demo2() {
	mesaj , _ := TahminEt(55)
	fmt.Println(mesaj)
	fmt.Println(TahminEt(-159))
	fmt.Println(TahminEt(-99))
}
