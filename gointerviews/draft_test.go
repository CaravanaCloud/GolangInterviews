package main

import (
	"fmt"
	"testing"
)

func TestDraft(t *testing.T) {
	items := [][]string{
		{"item2", "3", "4"},
		{"item3", "17", "8"},
		{"item1", "10", "15"}}
	sortParameter := int32(1)
	sortOrder := int32(0)
	itemsPerPage := int32(2)
	pageNumber := int32(1)
	got := fetchItemsToDisplay(items, sortParameter, sortOrder, itemsPerPage, pageNumber)
	fmt.Println(got)
}
