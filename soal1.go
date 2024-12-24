package main

import (
	"fmt"
)

func main() {
	fmt.Println(SloanesOEIS(5))
	fmt.Println(SloanesOEIS(20))
	fmt.Println(SloanesOEIS(10))
}

func SloanesOEIS(n int) string {
	var result string

	if n == 1 {
		return "1"
	}

	if n == 2 {
		return "1-2"
	}

	var pascalTriangle [][]int = [][]int{{1}, {1, 1}}

	for i := 3; i <= n; i++ {
		var tmp []int = []int{1}

		for k := 1; k <= i-2; k++ {
			tmp = append(tmp, pascalTriangle[i-2][k]+pascalTriangle[i-2][k-1])
			if len(tmp) == 3 {
				break
			}
		}

		tmp = append(tmp, 1)
		pascalTriangle = append(pascalTriangle, tmp)
	}

	for i := 0; i <= len(pascalTriangle)-1; i++ {
		var sum = 0
		if len(pascalTriangle[i]) >= 3 {
			for k := 0; k <= 2; k++ {
				sum += pascalTriangle[i][k]
			}
		} else {
			for k := 0; k <= len(pascalTriangle[i])-1; k++ {
				sum += pascalTriangle[i][k]
			}
		}

		if result == "" {
			result = fmt.Sprintf("%d", sum)
		} else {
			result = fmt.Sprintf("%s-%d", result, sum)
		}
	}

	return result
}
