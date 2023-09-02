package main

import (
	"fmt"
)

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {

	if mToSend.sender.name == "" {
		return false
	}

	if mToSend.recipient.name == "" {
		return false
	}

	if mToSend.sender.number == 0 {
		return false
	}

	if mToSend.recipient.number == 0 {
		return false
	}

	return true
}

func test(mToSend messageToSend) {
	fmt.Printf(`enviando "%s" from %s (%v) to %s (%v)...`,
		mToSend.message,
		mToSend.sender.name,
		mToSend.sender.number,
		mToSend.recipient.name,
		mToSend.recipient.number,
	)
	fmt.Println()
	if canSendMessage(mToSend) {
		fmt.Println("...Enviado!")
	} else {
		fmt.Println("...não pode enviar a mensagem")
	}
	fmt.Println("====================================")
}

func main() {
	test(messageToSend{
		message: "Vocẽ tem reunião amanhã",
		sender: user{
			name:   "Victor Reis",
			number: 16545550987,
		},
		recipient: user{
			name:   "Cícero",
			number: 19035558973,
		},
	})
	test(messageToSend{
		message: "Você tem um evento amanha",
		sender: user{
			number: 16545550987,
		},
		recipient: user{
			name:   "César",
			number: 0,
		},
	})
	test(messageToSend{
		message: "Você tem uma festa amanha",
		sender: user{
			name:   "Pompeu",
			number: 16545550987,
		},
		recipient: user{
			name:   "Otaviano",
			number: 19035558973,
		},
	})
	test(messageToSend{
		message: "Você tem um aniversário amanhã",
		sender: user{
			name:   "Tiro",
			number: 0,
		},
		recipient: user{
			name:   "Marco Aurelio",
			number: 19035558973,
		},
	})
}