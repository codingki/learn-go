package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investmentAmount float64 = 1000
	var years float64 = 10.0
	var expectedReturnRate float64 = 5.5

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected Return Rate (%): ")
	outputText("Expected Return Rate (%): ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future value: %.1f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future real value (adjusted for inflation): %.1f\n", futureRealValue)
	// fmt.Println("Future value: ", futureValue)
	// fmt.Printf(`Future value: %.1f
	//
	// Future real value (adjusted inflation): %.1f`, futureValue, futureRealValue)
	// fmt.Println("Future real value: ", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
