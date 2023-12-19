package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	// Using the custom functions
	v1 := get(m, "k1")
	fmt.Println("Value for key 'k1':", v1)

	lister(m)

	m = remove(m, "k2")
	lister(m)
}

// hadi katroutrner b key
func get(m map[string]int, key string) int {
	return m[key]
}

// lister les valeurs dyal key
func lister(m map[string]int) {
	fmt.Println("Listing key-value pairs:")
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}

// supprimer avec key
func remove(m map[string]int, key string) map[string]int {
	delete(m, key)
	fmt.Printf("Removed key '%s' from the map.\n", key)
	return m
}
