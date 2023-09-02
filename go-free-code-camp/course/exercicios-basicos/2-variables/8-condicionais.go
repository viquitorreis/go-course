package main

import "fmt"

func main() {
	messageLen := 10
	maxMessageLen := 20

	fmt.Println("Tentando enviar a mensagem do tamanho:", messageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Mensagem enviada")
	} else {
		fmt.Println("Mensagem não enviada")
	}

}