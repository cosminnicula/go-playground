package main

import "fmt"

func main() {
	// declaration method 1
	colors := map[string]string{
		"red":       "#ff0000",
		"something": "#abcdef",
	}

	fmt.Println(colors)

	// declaration method 2
	var emptyMap map[string]string

	fmt.Println(emptyMap)

	// declaration method 3
	someMap := make(map[string]string) // key of type string, value of type string

	someMap["white"] = "#ffffff"

	fmt.Println(someMap)

	delete(someMap, "white")

	fmt.Println(someMap)

	// iterate over maps
	printMap(colors)


}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for %v is %v\n", color, hex)
	}
}
