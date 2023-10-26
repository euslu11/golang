package main

import (
	"fmt"
	"math"
)

func main() {
	//variables
	var a = "Emirhan"
	fmt.Println(a)

	var b,c int = 18,74
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	f := "apple"
	fmt.Println(f)


}

//constants

func Demo1()  {
	const s  string = "constant"
	fmt.Println(s)

	const n  = 19050618
	const d = 3e20 / n 
	fmt.Println(d)

	fmt.Println(int64(n))
	fmt.Println(math.Sin(n))
}

//for

func Demo2() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i+1
	}

	for j:=1 ; j<=10 ; j++ {
		fmt.Println(j)
	}

	for n:=0 ; n<=5 ; n++ {
		if n%2  == 0{
			continue
		}
		fmt.Println(n)
	}

}

//if-else

func Demo3() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	
}































