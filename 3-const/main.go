package main

import "fmt"

func main() {
	const pi float64 = 3.1415
	const (
		pi float64 = 3.1415
		e float64 = 2.7182
	)
	const pi, e = 3.1415, 2.7182
	const n = 5     //  тип int

	// Если определяется последовательность констант, то инициализацию значением можно опустить для всех констант, кроме первой. В этом случае константа без значения полчит значение предыдущей константы:
	const (
		a = 1
		b
		c
		d = 3
		f
	)
	fmt.Println(a, b, c, d, f)      // 1, 1, 1, 3, 3

}