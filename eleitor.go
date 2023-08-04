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
	fmt.Println("Digite a região desejada: ")
	lugar := Letexto()
	digitosreg := regiao(lugar)
	rand.Seed(time.Now().Unix())
	var digitos []int
	for i := 0; i < 8; i++ {
		digitos = append(digitos, random())
	}
	digitos = append(digitos, digitosreg...)
	var multiplicadores = []int{
		2, 3, 4, 5, 6, 7, 8, 9,
	}
	soma := 0
	for i := 0; i < len(multiplicadores); i++ {
		soma += digitos[i] * multiplicadores[i]
	}
	digito1 := 0
	resto := soma % 11
	if resto == 10 {
		digito1 = 0
	} else {
		digito1 = resto
	}
	var mult2 = []int{
		7, 8, 9,
	}
	digitos = append(digitos, digito1)
	soma2 := 0
	for i := 0; i < len(mult2); i++ {
		soma2 += mult2[i] * digitos[8+i]
	}
	digito2 := 0
	resto = soma % 11
	if resto == 10 {
		digito2 = 0
	} else {
		digito2 = resto
	}
	digitos = append(digitos, digito2)
	fmt.Printf("%v", digitos)
}
func regiao(sigla string) []int {
	switch sigla {
	case "SP":
		return []int{0, 1}
	case "MG":
		return []int{0, 2}
	case "RJ":
		return []int{0, 3}
	case "RS":
		return []int{0, 4}
	case "BA":
		return []int{0, 5}
	case "PR":
		return []int{0, 6}
	case "CE":
		return []int{0, 7}
	case "PE":
		return []int{0, 8}
	case "SC":
		return []int{0, 9}
	case "GO":
		return []int{1, 0}
	case "MA":
		return []int{1, 1}
	case "PB":
		return []int{1, 2}
	case "PA":
		return []int{1, 3}
	case "ES":
		return []int{1, 4}
	case "PI":
		return []int{1, 5}
	case "RN":
		return []int{1, 6}
	case "AL":
		return []int{1, 7}
	case "MT":
		return []int{1, 8}
	case "MS":
		return []int{1, 9}
	case "DF":
		return []int{2, 0}
	case "SE":
		return []int{2, 1}
	case "AM":
		return []int{2, 2}
	case "RO":
		return []int{2, 3}
	case "AC":
		return []int{2, 4}
	case "AP":
		return []int{2, 5}
	case "RR":
		return []int{2, 6}
	case "TO":
		return []int{2, 7}
	default:
		return []int{2, 8}
	}
}
func Letexto() string {
	var i string
	_, err := fmt.Scanf("%s", &i)
	for err != nil {
		_, err = fmt.Scanf("%s", &i)
	}
	return i
}
