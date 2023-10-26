package loops

import "fmt"

//aklımda tahmin ettigim sayiyi bul.
func Demo3() {
	aksay := 18
	tahsay := 0
	tahmin := 0 

	fmt.Println("bir sayi tahmin et")
	fmt.Scanln(&tahsay)
	tahmin = tahmin + 1

	for aksay != tahsay {
		if tahsay > aksay {
			fmt.Println("Daha küçük bir sayi giriniz")
			fmt.Scanln(&tahsay)
			tahmin = tahmin + 1
		} else if tahsay < aksay {
			fmt.Println("Daha büyük bir sayi giriniz")
			fmt.Scanln(&tahsay)
			tahmin = tahmin + 1
		}

	}
	basari:=""
	if (tahmin>0 && tahmin <=3){
		basari ="süper"
	}else if tahmin <= 6 {
		basari = "güzel"
	}else {
		basari = "fena degil"
	}



	fmt.Printf("Helal Olsun Dayioglu. %v . Tahminde Bildin ve Tahmin Durumu = %v",tahmin,basari)	 
}