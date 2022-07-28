package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
	}

	fmt.Println(colors)
	delete(colors, "red")
	fmt.Println(colors)

	var m1 map[string]string
	//m1["black"] = "#000000"
	fmt.Println(m1)

	m2 := make(map[string]string)
	m2["green"] = "#4bf557"
	fmt.Println(m2)


	colors["green"] = "#4bf557"
	colors["blue"] = "#4bf7"
	printMap(colors)

}

func printMap(m map[string]string)  {
	for key, value := range m {
		fmt.Println(key, value)
	}
}