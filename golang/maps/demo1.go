package maps

import "fmt"

func Demo1() {
	//key-value
	sozluk := make(map[string]string)
	sozluk["table"] = "masa"
	sozluk["cyber"] = "siber"
	sozluk["pencil"] = "kalem"
	fmt.Println(sozluk["cyber"])
	fmt.Println("Eleman Sayisi :",len(sozluk))
	fmt.Println(sozluk)
	delete(sozluk,"pencil")
	fmt.Println("Eleman Sayisi :",len(sozluk))

	deger , varMi := sozluk["lamb"]
	fmt.Println(deger)
	fmt.Println("Listede var mi : ",varMi)

}