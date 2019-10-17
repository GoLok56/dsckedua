package main

import "fmt"

func main() {
	umurOrang := make(map[string] int)
	umurOrang["James"] = 17
	umurOrang["Catalunya"] = 18

	fmt.Println("Umurnya James adalah", umurOrang["James"])
}