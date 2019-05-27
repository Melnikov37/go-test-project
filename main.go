package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Сумма равна", add(10, 19))
	var bal = []int{1123, 2132, 2134, 2135, 3456, 7242}
	for i := 0; i < len(bal); i++ {
		fmt.Println("Sum -", i, bal[i])
	}
}
