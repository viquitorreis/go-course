package main

import "fmt"


func main() {

	const name = "Cícero"
	const openRate = 30.5

	msg := fmt.Sprintf("Oi %s, seu oopen rate é %.1f porcento", name, openRate)

	fmt.Println(msg)
}