package main

import "fmt"

func main() {
	sendsSoFar := 430
	const sendsToAdd = 25

	result := incrementSends(sendsSoFar, sendsToAdd)

	fmt.Println("Você enviou", result, "mensages")
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	return sendsSoFar + sendsToAdd
}