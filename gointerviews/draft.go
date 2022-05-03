package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

/*
 * Complete the 'fetchItemsToDisplay' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts following parameters:
 *  1. 2D_STRING_ARRAY items
 *  2. INTEGER sortParameter
 *  3. INTEGER sortOrder
 *  4. INTEGER itemsPerPage
 *  5. INTEGER pageNumber
 */

func fetchItemsToDisplay(items [][]string,
	sortParameter int32,
	sortOrder int32,
	itemsPerPage int32,
	pageNumber int32) []string {
	// Write your code here
	// sort the array
	// cut the pages
	// get the name from selected page
	fmt.Println(items)
	sortItems(items, sortParameter, sortOrder)
	fmt.Println(items)

	start := pageNumber * itemsPerPage
	end := (start + itemsPerPage) - 1
	max := int32(len(items) - 1)
	if end > max {
		end = max
	}
	page := items[start : end+1]
	names := make([]string, len(page))
	for i, _ := range page {
		names[i] = page[i][0]
	}
	return names
}

func sortItems(items [][]string,
	sortParameter int32,
	sortOrder int32) error {
	if sortOrder > 3 {
		return errors.New("Sort order bigger than 3")
	}
	sort.SliceStable(items, func(a, b int) bool {
		order := false
		valA := items[a][sortParameter]
		valB := items[b][sortParameter]
		ascending := sortOrder == 0

		if sortParameter > 0 {
			intA, _ := strconv.Atoi(valA)
			intB, _ := strconv.Atoi(valB)
			if ascending {
				order = intA < intB
			} else {
				order = intA > intB
			}
		} else {
			if ascending {
				order = valA < valB
			} else {
				order = valA > valB
			}
		}
		return order
	})
	return nil
}
