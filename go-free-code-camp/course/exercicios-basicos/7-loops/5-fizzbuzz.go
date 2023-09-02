package main

func fizzbuzz() {
	count := 1
	for count <= 100 {
		if count%3 == 0 {
			println("fizz")
		} else if count%5 == 0 {
			println("buzz")
		} else if count%3 == 0 && count%5 == 0 {
			println("fizzbuzz")
		} else {
			println(count)
		}
		count++
	}

}

func main() {
	fizzbuzz()
}
