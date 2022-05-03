package pairsum_test

import (
	"testing"

	pairsum "github.com/CaravanaCloud/GolangInterviews/gointerviews/pair_sum"
	"github.com/stretchr/testify/assert"
)

func TestPairSum(t *testing.T) {
	inputArr := []int{3, 5, -4, 8, 11, 1, -1, 6}
	targetSum := 10
	returnArr := pairsum.PairSumFunc(inputArr, targetSum)

	assert.Equal(t, (returnArr[0] + returnArr[1]), targetSum)
}

func TestPairSumHash(t *testing.T) {
	inputArr := []int{3, 5, -4, 8, 11, 1, -1, 6}
	targetSum := 10
	returnArr := pairsum.PairSumHashFunc(inputArr, targetSum)

	assert.Equal(t, (returnArr[0] + returnArr[1]), targetSum)
}

// array = [3, 5, -4, 8, 11, 1, -1, 6]
