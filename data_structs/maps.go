package main

import (
	"fmt"
)

func main() {
	// colors := map[string]string{
	// 	"red":   "#FF0000",
	// 	"green": "#00FF00",
	// }

	colors := make(map[string]string)
	fmt.Println(colors)
	colors["white"] = "#FFFFFF"
	colors["red"] = "#FF0000"
	colors["green"] = "#00FF00"
	fmt.Println(colors)

	empIds := make(map[int]string)
	empIds[1] = "Ananta"
	fmt.Println(empIds)
	delete(colors, "white")
	fmt.Println(colors)

	for i, j := range colors {
		fmt.Printf("Key: %v, Value: %v \n", i, j)
	}
}
