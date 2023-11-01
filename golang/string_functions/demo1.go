package string_functions

//alias
import (
	"fmt"
	"strings"
)

//case sensitive
//ascii
func Demo1() {
	isim := "golang"
	fmt.Println(strings.Count(isim,"g"))
	fmt.Println(strings.Contains(isim,"a"))
	sonuc := strings.Index(isim,"l")

	if sonuc != -1 {
		fmt.Println("A harfi var")
	} else {
		fmt.Println("A harfi yok")
	}
}