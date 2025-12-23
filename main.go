package main

import "fmt"

func main() {
	for {
		from, userInput, in := reading()

		result := calculateImp(from, userInput, in)
		
		fmt.Println(result)
		isRepeatCalculation := checkRepeatCalculate()
		if !isRepeatCalculation {
			break
		}
	}

}
func reading() (string, int, string) errors {

	var userInput int
	var from string
	var in string
	fmt.Println("=Калькулятор курса валют=")
	for {for {
		fmt.Print("Введите валюту из которой будете конвертировать (usd, euro): ")
		fmt.Scan(&from)
		if from == "usd" || from == "eur" {
			break
		} else {
			return 0 , errors.New("Ошибка, введите usd или eur: ")
		}
	}
	for {
		fmt.Print("Введите количество: ")
		fmt.Scan(&userInput)
		if userInput >= 0 {
			break
		} else {
			return 0, errors.New("Ошибка, введите число больше 0: ")
		}
	} 
	fmt.Print("Введите валюту в которую будете конвертировать (usd, euro): ")
	fmt.Scan(&in)
	return from, userInput, in
}
}
	

func calculateImp(from string, userInput int, in string) float64 {

	var result float64
	courseUsd := 76.1
	courseEur := 86.3

	if from == "usd" && in == "eur" {
		result = float64(userInput) * courseUsd / courseEur
	} else if from == "eur" && in == "usd" {
		result = float64(userInput) * courseEur / courseUsd
	} else {
		return 0
	}
	return result
}
func checkRepeatCalculate() bool {
	var userChoise string
	fmt.Print("Вы хотите сделать еще расчет (y/n): ")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
func checkReadingFrom() bool
