package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	if position == "up" {
		return append([]int{newData}, data...)
	} else if position == "down" {
		return append(data, newData)
	}
	return data
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	newData := 6
	position := "up"
	fmt.Println(AddElement(data, newData, position)) // [6 1 2 3 4 5]

	data = []int{1, 2, 3, 4, 5}
	newData = 6
	position = "down"
	fmt.Println(AddElement(data, newData, position)) // [1 2 3 4 5 6]
}
