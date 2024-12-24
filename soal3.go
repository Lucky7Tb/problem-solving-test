package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(BalancedBracket("(({{[[]]}}))"))
	fmt.Println(BalancedBracket("(({})"))
	fmt.Println(BalancedBracket("( ) { } [ ]"))
}

func BalancedBracket(s string) string {
	s = strings.Replace(s, " ", "", -1)
	var listOfBrackets = strings.Split(s, "")
	var stacks []string = []string{}
	for i := 0; i <= len(listOfBrackets)-1; i++ {
		if listOfBrackets[i] != ")" && listOfBrackets[i] != "]" && listOfBrackets[i] != "}" && listOfBrackets[i] != "(" && listOfBrackets[i] != "[" && listOfBrackets[i] != "{" {
			return "NO"
		}

		if listOfBrackets[i] == ")" && stacks[0] == "(" {
			stacks = append([]string{}, stacks[1:]...)
		} else if listOfBrackets[i] == "]" && stacks[0] == "[" {
			stacks = append([]string{}, stacks[1:]...)
		} else if listOfBrackets[i] == "}" && stacks[0] == "{" {
			stacks = append([]string{}, stacks[1:]...)
		} else {
			stacks = append([]string{listOfBrackets[i]}, stacks...)
		}
	}

	if len(stacks) == 0 {
		return "YES"
	}

	return "NO"
}
