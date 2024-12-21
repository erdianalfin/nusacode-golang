package main

import (
	"errors"
	"fmt"
	"golang/test"
)

func main() {
	test.Edit()
	fmt.Println("hello world")
	cetakAngka(12)
	jumlah := perjumlahan(10, 4)
	fmt.Println(jumlah)

	arr := coba()
	fmt.Println(arr[0])
	str, num := multiReturn("Hallo", 31)
	fmt.Println(str)
	fmt.Println(num)
	count := loop(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(count)

	hasil, error := bagi(4, 2)
	fmt.Println(hasil)
	fmt.Println(error)
}

func cetakAngka(angka int) {
	fmt.Println(angka)
}

func perjumlahan(a int, b int) int {
	return a + b
}

func coba() []string {
	var res []string = []string{"kambing", "kucing", "ayam"}
	return res
}

func multiReturn(a string, b int) (string, int) {
	return a, b
}

func loop(a ...int) int {
	fmt.Println(a)
	hitung := 0

	for i, result := range a {
		hitung += result
		fmt.Printf("Index %d: %v\n", i, result)
	}

	return hitung
}

func bagi(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Ini error")
	}

	return a / b, nil
}
