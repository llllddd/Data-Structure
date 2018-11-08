package main

import (
	"fmt"
)

//Factorial N!=N*(N-1)*...*2*1, Time Complexity O(N)
func Factorial(i int) int {
	if i <= 1 {
		return 1
	} else {
		return i * Factorial(i-1)
	}
}

//TowerofHanoi
func TOHUtil(num int, from string, to string, temp string) {
	if num < 1 {
		return
	}
	TOHUtil(num-1, from, temp, to)
	fmt.Println("Move disk", num, "from peg", from, "to peg", to)
	TOHUtil(num-1, temp, to, from)
}

//Greatest common divisior(GCD),gcd(m,n)=gcd(m,m%n)
func GCD(m int, n int) int {
	if m < n {
		return GCD(n, m)
	}
	if m%n == 0 {
		return n
	}
	return GCD(n, m%n)
}

//Fibonacci number
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	fmt.Println("factorial 5 is ::", Factorial(5))
	TOHUtil(5, "A", "C", "B")
	fmt.Println(GCD(15, 9))
}
