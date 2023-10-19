package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(212341233123)
	b := big.NewInt(2014312311232312331)

	// Перемножение
	multiplyResult := big.NewInt(0)
	multiplyResult.Mul(a, b)
	fmt.Println("Результат умножения:", multiplyResult)

	// Деление
	divideResult := big.NewInt(0)
	divideResult.Div(a, b)
	fmt.Println("Результат деления:", divideResult)

	// Сложение
	addResult := big.NewInt(0)
	addResult.Add(a, b)
	fmt.Println("Результат сложения:", addResult)

	// Вычитание
	subtractResult := big.NewInt(0)
	subtractResult.Sub(a, b)
	fmt.Println("Результат вычитания:", subtractResult)
}