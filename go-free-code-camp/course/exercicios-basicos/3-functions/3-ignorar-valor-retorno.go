package main

import "fmt"

func main() {

	firstName, _ := getNames()
	fmt.Println("Bem vindo ao curso de Golang :D", firstName)

}

func getNames() (string, string) {
	return "Victor", "Cícero"
}