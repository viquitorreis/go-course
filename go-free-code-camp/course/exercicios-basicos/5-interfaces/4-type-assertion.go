package main

import (
	"fmt"
)

func getExpenseReport(e expense) (string, float64) {
	emailVariable, ok := e.(email)

	if ok {
		return emailVariable.toAddress, emailVariable.cost()
	}

	smsVariable, ok := e.(sms)
	if ok {
		return smsVariable.toPhoneNumber, smsVariable.cost()
	}

	return "", 0.0

}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func estimateYearlyCost(e expense, averageMessagesPerYear int) float64 {
	return e.cost() * float64(averageMessagesPerYear)
}

func test(e expense) {
	address, cost := getExpenseReport(e)
	switch e.(type) {
	case email:
		fmt.Printf("Report: O email%s vai custar: %.2f\n", address, cost)
		fmt.Println("====================================")
	case sms:
		fmt.Printf("Report: The sms going to %s vai custar: %.2f\n", address, cost)
		fmt.Println("====================================")
	default:
		fmt.Println("Report: custo invalido")
		fmt.Println("====================================")
	}
}

func main() {
	test(email{
		isSubscribed: true,
		body:         "Eai pessoal",
		toAddress:    "john@does.com",
	})
	test(email{
		isSubscribed: false,
		body:         "Esse meeting poderia ter sido um email :)",
		toAddress:    "jane@doe.com",
	})
	test(email{
		isSubscribed: false,
		body:         "Esse meeting poderia ter sido um email :)",
		toAddress:    "elon@doe.com",
	})
	test(sms{
		isSubscribed:  false,
		body:          "Esse meeting poderia ter sido um email :)",
		toPhoneNumber: "+155555509832",
	})
	test(sms{
		isSubscribed:  false,
		body:          "Esse meeting poderia ter sido um email :)",
		toPhoneNumber: "+155555504444",
	})
	test(invalid{})
}
