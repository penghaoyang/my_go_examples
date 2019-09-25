package main

import (
	"fmt"

	"./simi"
)

func main() {
	data1 := []map[string]string{
		{"a": "1", "b": "2", "c": "3"},    // 1
		{"a": "5", "b": "12", "c": "21"},  // 2
		{"a": "3", "b": "4", "c": "5"},    // 1
		{"a": "5", "b": "6", "c": "5"},    // 2
		{"a": "7", "b": "8", "c": "6"},    // 3
		{"a": "3", "b": "2", "c": "6"},    // 1
		{"a": "15", "b": "13", "c": "11"}, // 4
		{"a": "8", "b": "8", "c": "8"},    // 2
		{"a": "9", "b": "6", "c": "10"},   // 3
		{"a": "3", "b": "11", "c": "12"},  // 1
		// {"a": "11", "b": "16", "c": "24"}, // 5 , 这个是不成功的会被归到 1 中，需要优化，
	}
	data1Result := simi.Aggr(data1, []string{"a", "b"})
	fmt.Println(data1Result)
}
