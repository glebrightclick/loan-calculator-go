package main

import (
	"fmt"
	"math"
)

type calculator struct {
	loanPrincipal  float64
	months         int
	monthlyPayment float64
}

func (c *calculator) calculateMonths() {
	months := c.loanPrincipal / c.monthlyPayment
	c.months = int(months)
	if float64(int(months)) != months {
		c.months++
	}
}

func (c *calculator) calculateMonthlyPayment() {
	c.monthlyPayment = math.Ceil(c.loanPrincipal / float64(c.months))
}

func (c *calculator) getLastMonthPayment() float64 {
	return c.loanPrincipal - c.monthlyPayment*float64(c.months-1)
}

func main() {
	calculator := calculator{}

	fmt.Println("Enter the loan principal:")
	fmt.Scan(&calculator.loanPrincipal)

	fmt.Println(
		"What do you want to calculate?\n" +
			"type \"m\" for number of monthly payments,\n" +
			"type \"p\" for the monthly payment:",
	)
	var option string
	fmt.Scan(&option)

	switch option {
	case "m":
		processNumberOfMonthlyPayments(calculator)
	case "p":
		processMontlyPayment(calculator)
	default:
		fmt.Println("Incorrect option")
	}
}

func processNumberOfMonthlyPayments(calculator calculator) {
	fmt.Println("Enter the monthly payment:")
	fmt.Scan(&calculator.monthlyPayment)
	calculator.calculateMonths()

	if calculator.months > 1 {
		fmt.Printf("It will take %d months to repay the loan\n", calculator.months)
	} else {
		fmt.Printf("It will take 1 month to repay the loan\n")
	}
}

func processMontlyPayment(calculator calculator) {
	fmt.Println("Enter the number of months:")
	fmt.Scan(&calculator.months)
	calculator.calculateMonthlyPayment()

	lastMonthPayment := calculator.getLastMonthPayment()
	if lastMonthPayment != calculator.monthlyPayment {
		fmt.Printf("Your monthly payment = %0.f and the last payment = %0.f.\n", calculator.monthlyPayment, lastMonthPayment)
	} else {
		fmt.Printf("Your monthly payment = %0.f\n", calculator.monthlyPayment)
	}
}
