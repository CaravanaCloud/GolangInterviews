package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*
 * Complete the 'counts' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY teamA
 *  2. INTEGER_ARRAY teamB
 */

func counts(teamA []int32, teamB []int32) []int32 {
	counts := make([]int32, len(teamB))
	mem := make(map[int32]int32)
	sort.Slice(teamA,
		func(i, j int) bool { return teamA[i] < teamA[j] })
	lenA := len(teamA)
	for indexB, scoreB := range teamB {
		count, exist := mem[int32(scoreB)]
		if !exist {
			for j := 0; j < lenA; j++ {
				scoreA := teamA[j]
				if scoreA <= scoreB {
					count = count + 1
				} else {
					break
				}
			}
			mem[int32(scoreB)] = count
		}
		counts[indexB] = count
	}
	return counts
}

func main() {
	alen := 100000
	a := make([]int32, alen)
	b := make([]int32, alen)
	for i := 0; i < alen; i++ {
		a[i] = int32(rand.Intn(alen))
		a[i] = int32(rand.Intn(alen))
	}
	start := time.Now()
	counts(a, b)
	end := time.Now()
	fmt.Println(end.Sub(start))
	//fmt.Println(out)
}
