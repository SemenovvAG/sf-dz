package main

import "fmt"

func input() (int, string, string) {
	var userInput int
	var usd float64 = 76.09
	var euro float64 = 88.70
	fmt.Print("=Калькулятор курса валют=")
	fmt.Println("Введите колличество рублей")
	fmt.Scan(&userInput)
	fmt.Print("Введите валюту из которой будете конвертровать(usd, euro)")
	fmt.Scan(&usd)
	fmt.Print("Введите валюту в которую будете конвертировать(usd, euro)")
	fmt.Scan(&euro)
	return userInput, usd, euro
}

func main() {

	reading := input(userInput)
}

func calc(int, string, string) float64 {
}
