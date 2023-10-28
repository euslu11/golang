package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"Laptop",1000,"Lenovo"})
	fmt.Println(product{"Laptop",750,""})
}

type product struct {
	name  string
	price float64
	brand string
}