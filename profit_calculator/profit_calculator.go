package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Enter total revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate (%): ")
	fmt.Scan(&taxRate)

	earningBeforeTax := revenue - expenses
	earningAfterTax := earningBeforeTax * (1 - taxRate/100)

	ratio := earningBeforeTax / earningAfterTax

	fmt.Println("Earnings Before Tax: ", earningBeforeTax)
	fmt.Println("Ratio of Earnings Before Tax to Earnings After Tax: ", ratio)
	fmt.Println("Earnings After Tax (Profit): ", earningAfterTax)
}
