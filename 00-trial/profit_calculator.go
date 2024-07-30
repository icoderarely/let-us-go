package main

import "fmt"

func main() {
	var revenue, expenses float64
	var taxRate float64 = 3.3

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("TaxRate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	fmt.Print(ebt, profit, ratio)
}
