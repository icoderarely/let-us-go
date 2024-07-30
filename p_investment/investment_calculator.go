package main

import (
	"fmt"
	"math"
)

func main() {
	var invAmt = 1000
	var intRate = 5.5
	var years = 10

	var futureValue = float64(invAmt) * math.Pow(1+intRate/100, float64(years))

	fmt.Print("future value is ", futureValue, "\n")

	// fmt.Print("hello")
	// fmt.Println("world")
	// fmt.Print("yo")

	name, age := "Kim", 22
	amt, years := 100.0, 23
}
