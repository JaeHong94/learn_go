package main

import "fmt"

func canIDrink(age int) bool {
	koreanAge := age + 2
	if koreanAge > 19 {
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink(18))
}
