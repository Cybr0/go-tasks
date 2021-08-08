package main

import "fmt"

func main() {
	var number int
	fmt.Print("Введите натуральное число: ")
	fmt.Scan(&number)
	odd, even := 0, 0

	for i := 1; i <= number; i++ {
		if i%2 == 0 {
			even += i
			continue
		}
		odd += i
	}
	fmt.Println("odd:  ", odd)
	fmt.Println("even: ", even)

}
