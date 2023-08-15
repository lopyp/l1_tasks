package main

import (
	"fmt"
	"log"
	"math/big"
)

func decCalc(a, b, op string) (string, error) {
	na, ok := big.NewInt(0).SetString(a, 10)
	if !ok {
		return "", fmt.Errorf("ошибка преобразования строки в число: %s", a)
	}

	nb, ok := big.NewInt(0).SetString(b, 10)
	if !ok {
		return "", fmt.Errorf("ошибка преобразования строки в число: %s", b)
	}

	switch op {
	case "+":
		return na.Add(na, nb).String(), nil
	case "-":
		return na.Sub(na, nb).String(), nil
	case "*":
		return na.Mul(na, nb).String(), nil
	case "/":
		return na.Div(na, nb).String(), nil
	default:
		return "", fmt.Errorf("неизвестная операция: %s", op)
	}
}

func main() {
	log.Println(decCalc("15453458987908723452345", "2342755233452345", "+"))
}
