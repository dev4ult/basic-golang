package exercise

import (
	"fmt"
	"strings"
)

func OddOrEvent(numb int) {
	for i := 1; i <= numb; i++ {
		if i%2 == 0 {
			fmt.Println(i, " adalah bilangan Genap")
		} else {
			fmt.Println(i, " adalah bilangan Ganjil")
		}
	}
}

func FizzBuzz100() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz ")
		} else if i%5 == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Print(i, " ")
		}
	}
}

func AsteriskTriangle(length int) {
	for i := 0; i < length; i++ {
		for j := 0; j < length - i - 1; j++ {
			fmt.Print(" ")
		}

		for k := 0; k <  2 * (i + 1) - 1; k++ {
			if k % 2 == 0 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		} 
		fmt.Println(" ")
	}
}

func ValidPalindrome(word string) {
	var backWord string
	
	for i := len(word) - 1; i >= 0; i-- {
		backWord += string(word[i])
	}

	if strings.TrimLeft(word, " ") == backWord {
		fmt.Println("a Palindrome")
	} else {
		fmt.Println("Not a Palindrome")
	}
}