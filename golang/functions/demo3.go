package functions

func ToplaVariadic(sayi ...int) int {
	toplam := 0

	for i := 0; i < len(sayi); i++ {
		toplam = toplam + sayi[i]
	}
	return toplam

}