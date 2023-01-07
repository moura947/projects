package main

import "fmt"

func main() {
	fmt.Println("Bem-vindo ao meu quiz!")
	
	fmt.Printf("Digite seu nome: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Olá, %v, bem-vindo ao jogo!\n", name)
	
	fmt.Printf("Digite a sua idade: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10{
		fmt.Println("Eba, você pode jogar :)")
	} else {
		fmt.Println("Você não pode jogar :(")
		return
	}
	
	score := 0
	num_questions :=3

	fmt.Printf("Qual o mais doce: mel ou sal? ")
	var answer string
	fmt.Scan(&answer)

	if answer == "mel" || answer == "Mel" ||answer == "MEL" {
		fmt.Println("Certa resposta!")
		score++
	} else {
		fmt.Println("Resposta incorreta.")
	}

	fmt.Printf("Quem tem o pêlo cinza é o Tom ou o Jerry? ")
	var answer2 string
	fmt.Scan(&answer2)

	if answer2 == "Tom" || answer2 == "tom" || answer2 == "TOM" {
		fmt.Println("Certa resposta!")
		score++
	} else {
		fmt.Println("Resposta incorreta.")
	}

	fmt.Printf("Em que ano estamos? ")
	var year uint
	fmt.Scan(&year)

	if year == 2023 {
		fmt.Println("Certa resposta!")
		score++
	} else {
		fmt.Println("Resposta incorreta.")
	}

	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("Você acertou %v de %v questões. Seu percentual de respostas corretas é %v%%.", score, num_questions, percent)

}