package main

import "fmt"

func main() {

	var username string = "victor"
	var password string = "26021999"

	fmt.Println("Authorization: Basic", username+":"+password)

}