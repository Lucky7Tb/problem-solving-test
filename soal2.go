package main

import (
	"fmt"
)

func main() {
	fmt.Println(DenseRanking([]int{150, 130, 110, 100, 95}, []int{105, 135, 155, 80}))

	fmt.Println(DenseRanking([]int{95, 85, 85, 75, 65}, []int{80, 90, 100}))

	fmt.Println(DenseRanking([]int{120, 110, 90, 90, 70}, []int{80, 100, 115, 110}))
}

func DenseRanking(score1 []int, score2 []int) []int {
	var result []int
	var ranking map[int]int = make(map[int]int)

	if len(score1) == 0 {
		return result
	}

	var lastRanking = 0
	ranking[lastRanking] = score1[0]
	for i := 1; i <= len(score1)-1; i++ {
		if ranking[lastRanking] != score1[i] {
			lastRanking++
			ranking[lastRanking] = score1[i]
		}
	}

	for i := 0; i <= len(score2)-1; i++ {
		rank := -1
		for k := 0; k <= len(ranking)-1; k++ {
			if score2[i] == ranking[k] {
				rank = k + 1
				break
			}

			if score2[i] > ranking[k] {
				if k == 0 {
					rank = k + 1
					break
				}

				if score2[i] < ranking[k-1] {
					rank = k + 1
					break
				}
			}
		}
		if rank == -1 {
			rank = len(ranking) + 1
		}
		result = append(result, rank)
	}

	return result
}
