package string_functions

import (
	"fmt"
	"strings"
)

func Demo2() {
	isim := "golang"
	fmt.Println(strings.HasPrefix(isim,"go"))
	fmt.Println(strings.HasSuffix(isim,"ng"))
	harf := []string{"g","o","l","a","n","g"}
	sonuc :=(strings.Join(harf,"/"))
	fmt.Println(sonuc)

	fmt.Println(strings.Replace(sonuc,"/","|",-1))
	fmt.Println(strings.Split(sonuc,"+"))
}