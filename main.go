package main

import "fmt"

func main() {

	userInput, from, in := reading()
	result := calculateImp(userInput, from, in)
	fmt.Println(result)

}
func reading() (int, string, string) {
	var userInput int
	var from string
	var in string
	fmt.Println("=Калькулятор курса валют=")
	fmt.Print("Введите количество: ")
	fmt.Scan(&userInput)

	fmt.Print("Введите валюту из которой будете конвертировать (usd, euro): ")
	fmt.Scan(&from)

	fmt.Print("Введите валюту в которую будете конвертировать (usd, euro): ")
	fmt.Scan(&in)
	return userInput, from, in
}

func calculateImp(userInput int, from string, in string) float64 {
	var result float64
	courseUsd := 76.1
	courseEuro := 86.3

	if from == "usd" && in == "euro" {
		result = float64(userInput) * courseUsd / courseEuro
	} else if from == "euro" && in == "usd" {
		result = float64(userInput) * courseEuro / courseUsd
	} else {
		fmt.Println("Неправильно введена валюта")
	}
	return result
}
