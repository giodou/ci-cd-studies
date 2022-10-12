package main

import "fmt"

func main()  {
	fmt.Println(Soma(10,30))
}

func Soma(a int, b int) int {
	return a + b
}

func Subtracao(a int, b int) int {
	return a - b
}

func Divisao(a int, b int) int {
	return a / b
}