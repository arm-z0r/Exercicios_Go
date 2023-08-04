package main

import (
	"fmt"
	"math/rand"
	"time"
)

func aleatorio() int {
	n := rand.Intn(10)
	return n
}
func main() {
	fmt.Println("Digite a região escolhida: ")
	estado := Letexto()
	digito := regiao(estado)
	rand.Seed(time.Now().Unix())
	var digitos []int
	for i := 0; i < 8; i++ {
		digitos = append(digitos, aleatorio())
	}
	digitos = append(digitos, digito)
	var multiplicadores = []int{
		10, 9, 8, 7, 6, 5, 4, 3, 2,
	}
	soma := 0
	for i := 0; i < len(digitos); i++ {
		soma += digitos[i] * multiplicadores[i]
	}
	d1 := 0
	resto := soma % 11
	if resto == 0 || resto == 1 {
		d1 = 0
	} else {
		d1 = 11 - resto
	}
	digitos = append(digitos, d1)
	soma2 := 0
	for i := 1; i < len(digitos); i++ {
		soma2 += digitos[i] * multiplicadores[i-1]
	}
	d2 := 0
	resto2 := soma2 % 11
	if resto2 == 0 || resto2 == 1 {
		d2 = 0
	} else {
		d2 = 11 - resto2
	}
	digitos = append(digitos, d2)
	for i := 0; i < len(digitos); i++ {
		fmt.Print(digitos[i])
	}
}
func regiao(sigla string) int {
	switch sigla {
	case "DF", "GO", "MS", "MT", "TO":
		return 1
	case "AC", "AM", "AP", "PA", "RO", "RR":
		return 2
	case "CE", "MA", "PI":
		return 3
	case "AL", "PB", "PE", "RN":
		return 4
	case "BA", "SE":
		return 5
	case "MG":
		return 6
	case "ES", "RJ":
		return 7
	case "SP":
		return 8
	case "PR", "SC":
		return 9
	case "RS":
		return 0
	default:
		return -1
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
