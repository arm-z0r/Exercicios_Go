package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Letexto() string {
	var result string
	for result == "" {
		fmt.Scanf("%s", &result)
	}
	return result
}

func aleatorio() int {
	n := rand.Intn(10)
	return n
}
func regiaocpf(sigla string) int {
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
func random() int {
	n := rand.Intn(10)
	return n
}

func random2() int {
	n := rand.Intn(10)
	return n
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
func main() {
	var prog string
	fmt.Printf("Qual dos seguintes documentos deseja gerar: CPF, RENAVAM ou TITULO de eleitor?\n")
	prog = Letexto()
	switch prog {
	case "CPF":
		geraCPF()
	case "RENAVAM":
		geraRENAVAM()
	case "TITULO":
		geraTitulo()
	default:
		fmt.Printf("Não sei gerar esse documento.\n")
	}
}
func geraCPF() {
	fmt.Println("Digite a sigla do Estado dessejada: ")
	estado := Letexto()
	digito := regiaocpf(estado)
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
func geraRENAVAM() {
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
func geraTitulo() {
	fmt.Println("Digite a sigla do Estado dessejada: ")
	lugar := Letexto()
	digitosreg := regiao(lugar)
	rand.Seed(time.Now().Unix())
	var digitos []int
	for i := 0; i < 8; i++ {
		digitos = append(digitos, random2())
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
