package main

import (
	"fmt"
	"math"
)

func Lefloat() float64 {
	var i float64
	_, err := fmt.Scanf("%f", &i)
	for err != nil {
		_, err = fmt.Scanf("%f", &i)
	}
	return i
}
func Letexto() string {
	var i string
	_, err := fmt.Scanf("%s", &i)
	for err != nil {
		_, err = fmt.Scanf("%s", &i)
	}
	return i
}

func soma(a, b float64) float64 {
	s := a + b
	return s
}

func menos(a, b float64) float64 {
	m := a - b
	return m
}

func vezes(a, b float64) float64 {
	v := a * b
	return v
}

func divide(a, b float64) float64 {
	d := a / b
	return d
}
func expoente(a, b float64) float64 {
	a = math.Pow(a, b)
	return a
}
func expo(a, b float64) float64 {
	//a ^ b = a * a repetido b vezes
	expo := 1.0
	for i := 0; i < int(b); i++ {
		expo = expo * (a)
	}
	return expo
}
func fatorial(a float64) float64 {
	if a == 0 {
		return 1
	} else {
		return a * fatorial(a-1)
	}
}
func expo2(a, b float64) float64 {
	if b == 0 {
		return 1
	} else {
		return a * expo2(a, (b-1))
	}
}
func fibo(a float64) float64 {
	if a == 0 {
		return 0
	} else if a == 1 {
		return 1
	} else {
		return fibo(a-1) + fibo(a-2)
	}
}
func somat(a float64) float64 {
	if a == 0 {
		return 0
	} else {
		return a + somat(a-1)
	}
}

func main() {
	fmt.Printf("Digite o primeiro número: ")
	pd := Lefloat()
	fmt.Printf("Digite a operação: ")
	op := Letexto()
	fmt.Printf("Digite o segundo número: ")
	sd := Lefloat()
	var resultado float64
	switch op {
	case "+":
		resultado = soma(pd, sd)
	case "-":
		resultado = menos(pd, sd)
	case "*":
		resultado = vezes(pd, sd)
	case "/":
		resultado = divide(pd, sd)
	case "^":
		resultado = expo2(pd, sd)
	case "!":
		resultado = fatorial(pd)
	case "f":
		resultado = fibo(pd)
	case "s":
		resultado = somat(pd)
	}
	fmt.Println(resultado)
}

//Escreva uma função recursiva que receba um numero e retorne o somatario de todos os numeros até ele
//Por exemplo
//Somat(9) deveria retornar 45
