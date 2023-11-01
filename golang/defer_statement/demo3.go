package deferstatement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() { 
	fmt.Println("Urun Kaydedildi : ",p.productName)
}

func Demo3() {
	p :=product{productName: "Laptop",unitPrice: 7500}
	defer p.save()
	p = product{productName: "Mouse",unitPrice: 50}
	defer p.save()
	fmt.Println("İslem Basarili")
}
