package main

import (
	"flag"
	"fmt"
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
	c.principal = math.Floor(c.payment * (math.Pow(1+i, c.numberOfPayments) - 1) / (i * math.Pow(1+i, c.numberOfPayments)))
}

func (c *calculator) getDifferentiatedPayment(month int) float64 {
	i := c.getMonthlyInterestRate()
	return math.Ceil(c.principal/c.numberOfPayments + i*(c.principal-c.principal*(float64(month)-1)/c.numberOfPayments))
}

// The only thing left is to compute the overpayment: the amount of interest paid over the whole term of the loan.

func main() {
	initValue := 0.00
	calc := calculator{}
	var calculationType string

	flag.Float64Var(&calc.principal, "principal", initValue, "loan principal")
	flag.Float64Var(&calc.payment, "payment", initValue, "monthly payment")
	flag.Float64Var(&calc.numberOfPayments, "periods", initValue, "number of payments")
	flag.Float64Var(&calc.interest, "interest", initValue, "interest rate")
	flag.StringVar(&calculationType, "type", "", "type of calculation")
	flag.Parse()

	switch calculationType {
	case "diff":
		emptyValues, incorrectValues := 0, 0
		for _, value := range []float64{calc.principal, calc.numberOfPayments, calc.interest} {
			if value == initValue {
				emptyValues++
			}
			if value < initValue {
				incorrectValues++
			}
		}
		if emptyValues > 1 || incorrectValues > 0 {
			fmt.Println("Incorrect parameters")
			return
		}

		processDifferentiatedPayment(calc)
	case "annuity":
		// calculate empty and incorrect values
		emptyValues, incorrectValues := 0, 0
		for _, value := range []float64{calc.payment, calc.principal, calc.numberOfPayments} {
			if value == initValue {
				emptyValues++
			}
			if value < initValue {
				incorrectValues++
			}
		}
		// if we have more than one empty value, or we have any incorrect values, return error message
		if calc.interest <= initValue {
			incorrectValues++
		}
		if emptyValues > 1 || incorrectValues > 0 {
			fmt.Println("Incorrect parameters")
			return
		}

		switch initValue {
		case calc.payment:
			processMonthlyPayment(calc)
		case calc.principal:
			processPrincipal(calc)
		case calc.numberOfPayments:
			processNumberOfMonthlyPayments(calc)
		}
	default:
		fmt.Println("Incorrect parameters")
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
	fmt.Printf("Overpayment = %0.f\n", calc.payment*calc.numberOfPayments-calc.principal)
}

func processMonthlyPayment(calc calculator) {
	// perform calculation
	calc.calculatePayment()
	// display result
	fmt.Printf("Your annuity payment = %0.f!\n", calc.payment)
	fmt.Printf("Overpayment = %0.f\n", calc.payment*calc.numberOfPayments-calc.principal)
}

func processPrincipal(calc calculator) {
	// perform calculation
	calc.calculatePrincipal()
	// display result
	fmt.Printf("Your loan principal = %0.f!\n", calc.principal)
	fmt.Printf("Overpayment = %0.f\n", calc.payment*calc.numberOfPayments-calc.principal)
}

func processDifferentiatedPayment(calc calculator) {
	// perform calculation
	total := 0.00
	for i := 1; i <= int(calc.numberOfPayments); i++ {
		// calculate differentiated payment
		differentiatedPayment := calc.getDifferentiatedPayment(i)

		fmt.Printf("Month %d: payment is %0.f\n", i, differentiatedPayment)
		total += differentiatedPayment
	}

	fmt.Printf("\nOverpayment = %0.f\n", total-calc.principal)
}
