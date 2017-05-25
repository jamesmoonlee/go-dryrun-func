package main

import "fmt"

func main(){
	var i int = 10
	var j int
	fmt.Printf("Factorial of %d is: %d\n", i, factorial(i))
	fmt.Printf("Fibonaci serial of %d is: ", i)

	for j = 0; j <= i ; j++ {
		fmt.Printf("%d ", fibonacci(j))
	}
}

func factorial(i int) int {
	if (i <= 1) {
		return 1
	} else {
		return i * factorial(i-1)
	}
}


func fibonacci(i int) int {
	if (i == 0) {
		return 0
	} else if (i == 1) {
		return 1
	} else {
		return (fibonacci(i - 1) + fibonacci(i - 2))
	}
}
