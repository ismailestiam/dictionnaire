package main

import "fmt"

func main() {
	// Creating a slice (list) in Go
	myList := []int{1, 2, 3, 4, 5}

	// Accessing elements in the list using get function
	v1 := get(myList, 2)
	fmt.Println("Element at index 2 (using get):", v1)

	// Modifying an element in the list
	myList[2] = 10
	fmt.Println("Updated list:", myList)

	// Removing an element from the list using remove function
	myList = remove(myList, 3)
	fmt.Println("List after removing element at index 3:", myList)

	// Appending elements to the list
	myList = append(myList, 6, 7, 8)
	fmt.Println("List after appending:", myList)

	// Iterating through the list
	fmt.Println("Iterating through the list:")
	for index, value := range myList {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}

// get function retrieves an element from the slice at a specified index
func get(list []int, index int) int {
	if index >= 0 && index < len(list) {
		return list[index]
	}
	// Return a default value or handle the error as needed
	return -1
}

// remove function removes an element from the slice at a specified index
func remove(list []int, index int) []int {
	if index >= 0 && index < len(list) {
		return append(list[:index], list[index+1:]...)
	}
	// Return the original slice if index is out of bounds
	return list
}
