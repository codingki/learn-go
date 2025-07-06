package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount float64 = 1000
	var years float64 = 10.0
	var expectedReturnRate float64 = 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate (%): ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future value: %.1f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future real value (adjusted for inflation): %.1f\n", futureRealValue)
	// fmt.Println("Future value: ", futureValue)
	// fmt.Printf("Future value: %.1f \nFuture real value (adjusted inflation): %.1f\n", futureValue, futureRealValue)
	// fmt.Println("Future real value: ", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}
