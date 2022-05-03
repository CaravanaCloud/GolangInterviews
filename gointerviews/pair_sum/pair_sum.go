package pairsum

import "fmt"

func PairSumFunc(arr []int, target int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if (arr[i] + arr[j]) == target {
				fmt.Println(arr[i], arr[j])
				return []int{arr[i], arr[j]}
			}
		}
	}

	return nil
}

func PairSumHashFunc(arr []int, target int) []int {
	memoria := make(map[int]int) //hash table

	for i := 0; i < len(arr)-1; i++ {
		atual := arr[i]
		falta := target - atual

		if _, ok := memoria[falta]; ok {
			return []int{atual, falta}
		}

		memoria[atual] = atual
	}

	return nil
}
