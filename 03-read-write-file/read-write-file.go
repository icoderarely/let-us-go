package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Choose \n1. Write to a file \n2. Read from a file")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Print("Enter the amount: ")
		var amount float64
		fmt.Scan(&amount)
		writeFile(amount)
	case 2:
		amount := readFile()
		fmt.Println("The amount in the file is:", amount)
	}
}

func writeFile(amount float64) {

	amtString := fmt.Sprint(amount)
	content := []byte(amtString)
	// balText := fmt.Sprint(amount)
	filename := "balance.txt"
	err := os.WriteFile(filename, content, 0644)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File written successfully 'balance.txt'")
}

func readFile() float64 {
	fmt.Println("Opening balance.txt")
	filename := "balance.txt"

	data, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("First write to a file")
			panic(err)
			// return 0
		}
		fmt.Println("Error reading the file:", err)
		return 0
	}

	amtText := string(data)
	amt, err := strconv.ParseFloat(amtText, 64)
	if err != nil {
		fmt.Println("Error parsing the file content to float:", err)
		return 0
	}

	fmt.Println("File read successful")
	return amt
}
