package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random() int {
	n := rand.Intn(10)
	return n
}
func main() {
	rand.Seed(time.Now().Unix())
	var digitos []int
	for i := 0; i < 12; i++ {
		digitos = append(digitos, random())
	}
	var multiplicadores = []int{
		5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2,
	}
	soma := 0
	for i := 0; i < len(multiplicadores); i++ {
		soma += digitos[i] * multiplicadores[i]
	}
	digito1 := 0
	resto := soma % 11
	if resto < 2 {
		digito1 = 0
	} else {
		digito1 = 11 - resto
	}
	digitos = append(digitos, digito1)
	var multiplica2 = []int{
		6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2,
	}
	soma2 := 0
	for i := 0; i < len(multiplica2); i++ {
		soma2 += digitos[i] * multiplica2[i]
	}
	digito2 := 0
	resto = soma2 % 11
	if resto < 2 {
		digito2 = 0
	} else {
		digito2 = 11 - resto
	}
	digitos = append(digitos, digito2)
	fmt.Printf("%v", digitos)
}
