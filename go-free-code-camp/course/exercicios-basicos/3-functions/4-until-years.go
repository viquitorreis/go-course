package main

import "fmt"

func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {

	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}

	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}

	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}

	return

}

func teste(age int) {
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(age)
	fmt.Println("Você vai ser aulto em", yearsUntilAdult)
	fmt.Println("Você vai beber em", yearsUntilDrinking)
	fmt.Println("Você vai poder fazer empréstimo de um carro em", yearsUntilCarRental)
}

func main() {
	teste(21)
}