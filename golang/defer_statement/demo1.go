package deferstatement

import "fmt"

func A() {
	fmt.Println("A fonskiyonu çalıştı")
}

func B() {
	fmt.Println("B fonskiyonu çalıştı")
	defer A()
	defer C()
}
func C() {
	fmt.Println("C fonskiyonu çalıştı")
}
