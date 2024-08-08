package main

import "fmt"

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	outputResults(ebt, profit, ratio)
}

func outputResults(ebt, profit, ratio float64) {
	fmt.Println("Earnings Before Tax: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Printf("Ratio: %.2f", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func getUserInput(ask string) (userInput float64) {
	fmt.Print(ask)
	fmt.Scan(&userInput)
	return
}
