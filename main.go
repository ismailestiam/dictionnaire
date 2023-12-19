package main

import (
	"fmt"
)

// Dictionary represents a simple dictionary
type Dictionary struct {
	data map[string]string
}

func NewDictionary() *Dictionary {
	return &Dictionary{data: make(map[string]string)}
}

func (d *Dictionary) Add(key, value string) {
	d.data[key] = value
}

func (d *Dictionary) Get(key string) string {
	return d.data[key]
}

func (d *Dictionary) Remove(key string) {
	delete(d.data, key)
}

// List prints all key-value pairs in the dictionary
func (d *Dictionary) List() {
	fmt.Println("Listing key-value pairs:")
	for key, value := range d.data {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}

func main() {
	// Using the dictionary methods
	dict := NewDictionary()

	dict.Add("word1", "definition1")
	dict.Add("word2", "definition2")

	fmt.Println("Get 'word1':", dict.Get("word1"))

	dict.List()

	dict.Remove("word2")
	dict.List()
}
