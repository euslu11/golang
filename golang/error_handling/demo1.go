package error_handling

import (
	"fmt"
	"os"
)

//type assertion
func Demo1() {
	f,err := os.Open("demo11.txt")
	//nil

	if err != nil {
		if pErr,ok := err.(*os.PathError); ok{
			fmt.Println("Dosya Bulunamadı",pErr.Path)
			return
		}else{
			fmt.Println("Dosya Bulunamadı")
			return
		}
	}else {
		fmt.Println(f.Name())
	}
}
