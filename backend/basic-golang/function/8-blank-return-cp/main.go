package main

import "fmt"

//fungsi square sama seperti sebelumnya
//yang membedakan adalah fungsi ini
//menggunakan blank return
func main() {
	result1, result2 := square(4, 5)
	fmt.Printf("%d dan %d\n", result1, result2)
	fmt.Println(square(9, 8))
}
func square(number1, number2 int) (result1, result2 int) {
	result1 = number1 * number1
	result2 = number2 * number2
	return
}

// TODO: answer here
