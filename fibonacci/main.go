package main

import "fmt"

func main() {
	fmt.Print("Введите натуральное число: ")
	var fib int
	fmt.Scan(&fib)

	n0 := 1
	n1 := 1
	var n2 int
	fmt.Println("Последовательность фибоначи:")
	fmt.Printf("%d\n%d\n", n0, n1)
	for i := 3; i <= fib; i++ {
		n2 = n0 + n1
		fmt.Println(n2)
		n0, n1 = n1, n2
	}
}
