package main

import "fmt"

func main() {
	// different ways to create a map

	// 1.
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#0f0a10",
		"blue":  "#82ab71",
	}

	fmt.Printf("colors: %v\n", colors)

	// 2.
	var smells = map[string]string{}

	fmt.Printf("smells: %v\n", smells)

	// 3.
	dished := make(map[string]string)

	fmt.Printf("dished: %v\n", dished)

	// delete an entry
	delete(colors, "red")

	fmt.Printf("colors: %v\n", colors)
}
