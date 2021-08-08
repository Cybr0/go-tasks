package main

import "fmt"

func main() {
	fizz := "Fizz"
	bazz := "Bazz"
	fizzbazz := fizz + bazz

	for i := 1; i < 101; i++ {
		if i%15 == 0 {
			fmt.Println(fizzbazz)
			continue
		}

		if i%3 == 0 {
			fmt.Println(fizz)
			continue
		}

		if i%5 == 0 {
			fmt.Println(bazz)
			continue
		}
		fmt.Println(i)
	}
}
