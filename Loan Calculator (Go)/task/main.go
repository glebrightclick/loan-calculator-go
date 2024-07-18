package main

import (
	"flag"
	"fmt"
	"log"
	"math"
)

type calculator struct {
	principal        float64
	numberOfPayments float64
	payment          float64
	interest         float64
}

// Convert annual interest rate to monthly and to a decimal
func (c *calculator) getMonthlyInterestRate() float64 {
	return c.interest / (12 * 100)
}

func (c *calculator) calculateNumberOfPayments() {
	i := c.getMonthlyInterestRate()
	c.numberOfPayments = math.Ceil(math.Log(c.payment/(c.payment-i*c.principal)) / math.Log(1+i))
}

func (c *calculator) calculatePayment() {
	i := c.getMonthlyInterestRate()
	c.payment = math.Ceil(c.principal * (i * math.Pow(1+i, c.numberOfPayments)) / (math.Pow(1+i, c.numberOfPayments) - 1))
}

func (c *calculator) getLastPayment() float64 {
	return c.principal - c.payment*(c.numberOfPayments-1.00)
}

func (c *calculator) calculatePrincipal() {
	i := c.getMonthlyInterestRate()
	c.principal = c.payment * (math.Pow(1+i, c.numberOfPayments) - 1) / (i * math.Pow(1+i, c.numberOfPayments))
}

func main() {
	initValue := 0.00
	calc := calculator{}

	flag.Float64Var(&calc.principal, "principal", initValue, "loan principal")
	flag.Float64Var(&calc.payment, "payment", initValue, "loan principal")
	flag.Float64Var(&calc.numberOfPayments, "periods", initValue, "loan principal")
	flag.Float64Var(&calc.interest, "interest", initValue, "loan principal")
	flag.Parse()

	switch initValue {
	case calc.payment:
		processMonthlyPayment(calc)
	case calc.principal:
		processPrincipal(calc)
	case calc.numberOfPayments:
		processNumberOfMonthlyPayments(calc)
	case calc.interest:
		log.Fatal("Not yet implemented")
	}
}

func processNumberOfMonthlyPayments(calc calculator) {
	// perform calculation
	calc.calculateNumberOfPayments()
	// display result

	// It will take 8 years and 2 months to repay this loan!
	years := math.Floor(calc.numberOfPayments / 12)
	months := calc.numberOfPayments - 12*years

	if years == 0 {
		if months > 1 {
			fmt.Printf("It will take %0.f months to repay the loan\n", months)
		} else {
			fmt.Printf("It will take 1 month to repay the loan\n")
		}
	} else if months == 0 {
		if years > 1 {
			fmt.Printf("It will take %0.f years to repay the loan\n", years)
		} else {
			fmt.Printf("It will take 1 year to repay the loan\n")
		}
	} else {
		yearsPart, monthsPart := fmt.Sprintf("%0.f years", years), fmt.Sprintf("%0.f months", months)
		if years == 1 {
			yearsPart = "1 year"
		}
		if months == 1 {
			monthsPart = "1 month"
		}

		fmt.Printf("It will take %s and %s to repay the loan\n", yearsPart, monthsPart)
	}
}

func processMonthlyPayment(calc calculator) {
	// perform calculation
	calc.calculatePayment()
	// get last month payment
	lastMonthPayment := calc.payment // calc.getLastPayment()
	// display result
	if lastMonthPayment != calc.payment {
		fmt.Printf("Your monthly payment = %0.f and the last payment = %0.f.\n", calc.payment, lastMonthPayment)
	} else {
		fmt.Printf("Your monthly payment = %0.f!\n", calc.payment)
	}
}

func processPrincipal(calc calculator) {
	// perform calculation
	calc.calculatePrincipal()
	// display result
	fmt.Printf("Your loan principal = %0.f!\n", calc.principal)
}
