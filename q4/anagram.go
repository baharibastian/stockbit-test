package main

import (
	"fmt"
	"sort"
)

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	r := StringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func GetGroupedAnagram(words []string) map[string][]string {
	result := make(map[string][]string, 0)

	for _, word := range words {
		sortedWord := SortStringByCharacter(word)
		result[sortedWord] = append(result[sortedWord], word)
	}

	return result
}

func main() {
	words := []string{"eat","tea","tan","ate","nat","bat"}
	res := GetGroupedAnagram(words)
	for _, g := range res {
		fmt.Println(g)
	}
}
