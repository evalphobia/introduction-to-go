// +build ignore

package main

func main() {
	// START OMIT
	var π = 3.14159
	var radius = 6371.0 // 地球の半径（km）
	var circumference = 2 * π * radius

	println("地球の円周は", circumference, "kmです")
	// END OMIT
}
