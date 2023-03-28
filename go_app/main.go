// Created By: Lamees Hemdan
// Created On: March 2023
//
// This file contains the main function for the go_app application.

package main

import (
	"fmt"
)

func main() {
	const rate float64 = 0.18
	var hours float64
	var pay float64
	var governmentPay float64
	var governmentPayRounded float64
	var salary float64
	var roundedPay float64

	fmt.Println("Enter Hours:")
	fmt.Scan(&hours)
	fmt.Println("Enter Salary:")
	fmt.Scan(&salary)

	pay = (hours * salary) * (1 - rate)
	governmentPay = (hours * salary) * rate
	roundedPay:= fmt.Sprintf("%.2f", pay)
	governmentPayRounded:= fmt.Sprintf("%.2f", governmentPay)

	fmt.Println("Pay:", pay)
	fmt.Println("Government Pay:", governmentPay)

	fmt.Println("\nDone.")
}
