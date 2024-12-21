package main

import "fmt"

func main() {
	fmt.Println("======= SOAL 1 =========")
	perjumlahan := add(31, 5)
	fmt.Println(31, "+", 5, "=", perjumlahan)
	fmt.Println("======= SOAL 2 =========")
	bagi := bagi(10, 5)
	fmt.Println(10, "+", 5, "=", bagi)
	fmt.Println("======= SOAL 3 =========")
	pengurangan := KurangVariadic(5, 4, 3, 2, 1)
	fmt.Println(5, "-", 4, "-", 3, "-", 2, "-", 1, "=", pengurangan)
}

func add(a int, b int) int {
	c := a + b
	return c
}

func bagi(a int, b int) int {
	c := a / b
	return c
}

func KurangVariadic(numbers ...int) int {
	// Pastikan jumlah parameter tepat 5
	if len(numbers) != 5 {
		panic("Fungsi ini hanya menerima 5 parameter int.")
	}

	// Mulai pengurangan dari elemen pertama
	result := numbers[0]
	for _, num := range numbers[1:] {
		result -= num
	}
	return result
}
