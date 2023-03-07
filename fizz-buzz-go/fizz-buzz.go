package fizzBuzz

import "fmt"

func fizzBuzz(n int) []string {
	var result []string = make([]string, 0, n) // キャパシティを指定
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			result = append(result, "FizzBuzz")
		case i%3 == 0:
			result = append(result, "Fizz")
		case i%5 == 0:
			result = append(result, "Buzz")
		default:
			result = append(result, fmt.Sprint(i))
		}
	}
	return result
}
