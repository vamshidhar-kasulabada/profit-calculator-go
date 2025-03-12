package main

import "fmt"

func main() {
	var revenue float64
	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	var expense float64
	fmt.Print("Enter expense: ")
	fmt.Scan(&expense)
	var taxRate float64
	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)
	calculateAndDisplayProfit(revenue, expense, taxRate)
}

func calculateAndDisplayProfit(revenue float64, expense float64, taxRate float64) {
	var ebt float64 = revenue - expense
	fmt.Println("expense before tax", ebt)
	var profit float64 = ebt * (1 - taxRate/100)
	fmt.Println("profit:", profit)
	var ratio = ebt / profit
	fmt.Println("ratio of ebt to profit:", ratio)
}
