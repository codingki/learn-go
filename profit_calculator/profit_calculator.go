package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err1 := getUserInput("Enter total revenue: ")
	expenses, err2 := getUserInput("Enter total expenses: ")
	taxRate, err3 := getUserInput("Enter tax rate (in %): ")
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Error:", "All inputs must be greater than zero.")
		return
	}
	earningBeforeTax, ratio, earningAfterTax := calculateEarnigs(revenue, expenses, taxRate)

	ebt := fmt.Sprintf("Earnings Before Tax: %.1f\n", earningBeforeTax)
	rat := fmt.Sprintf("Ratio of Earnings Before Tax to Earnings After Tax: %.1f\n", ratio)
	pro := fmt.Sprintf("Earnings After Tax (Profit): %.1f\n", earningAfterTax)

	fmt.Print(ebt, rat, pro)

	content := fmt.Sprintf("%s%s%s", ebt, rat, pro)

	os.WriteFile("profit.txt", []byte(content), 0644)
}

func getUserInput(text string) (input float64, err error) {
	fmt.Print(text)
	fmt.Scan(&input)
	if input <= 0 {
		input = 0
		return input, errors.New("input must be greater than zero")
	}

	return input, nil
}

func calculateEarnigs(revenue, expenses, taxRate float64) (earningBeforeTax, ratio, earningAfterTax float64) {
	earningBeforeTax = revenue - expenses
	earningAfterTax = earningBeforeTax * (1 - taxRate/100)
	ratio = earningBeforeTax / earningAfterTax
	return earningBeforeTax, ratio, earningAfterTax
}
