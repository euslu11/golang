package structs

import "fmt"

type customer struct {
	name    string
	surname string
	age     int
}

func(c customer) save() {
	fmt.Println("Eklendi : ",c.name,c.surname,c.age)
}

func Demo2() {
	c := customer{name: "Emirhan",surname: "Uslu",age: 21}
	c.save()
}