package main

import "fmt"

func main() {
	var buah = []string{"apel", "jeruk", "anggur", "melon"}

	buah = append(buah, "pepaya")

	fmt.Println("Jumlah Element: ", len(buah))
	fmt.Println("Isi Element: ", buah[:])

	for k, v := range buah {
		fmt.Println("Element ke-", k, " : ", v)
	}
}
