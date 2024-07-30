package main

import "fmt"

func main() {
	var revenue, expenses, taxes float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	// fmt.Scanf("%f", &revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scanf("%f", &expenses)

	fmt.Print("Enter taxes: ")
	fmt.Scanf("%f", &taxes)

	ebt := revenue - expenses
	profit := ebt * (1 - taxes/100)

	ratio := ebt / profit

	fmt.Print("EBT: ", ebt, "\n")
	fmt.Printf("Profit: %.2f\nRatio: %.2f\n", profit, ratio)
}
