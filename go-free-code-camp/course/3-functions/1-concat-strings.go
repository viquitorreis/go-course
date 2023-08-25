package main

import "fmt"

func concatString(firstWord, secondWord string) string {

	return firstWord + secondWord

}

func main() {
	fmt.Println(concatString("Hello", "World"))
}