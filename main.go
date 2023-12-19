package main

import (
	"fmt"
)

type dictionnaire struct {
	data map[string]string
}

func NewDictionary() *dictionnaire {
	return &dictionnaire{data: make(map[string]string)}
}

func (d *dictionnaire) Add(key, value string) {
	d.data[key] = value
}

func (d *dictionnaire) Get(key string) string {
	return d.data[key]
}

func (d *dictionnaire) Remove(key string) {
	delete(d.data, key)
}

func (d *dictionnaire) List() {
	fmt.Println("Listing key-value pairs:")
	for key, value := range d.data {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}

func main() {

	dict := NewDictionary()

	dict.Add("estiam", "ecole")
	dict.Add("ismail", "etudiant")

	fmt.Println("Get 'estiam':", dict.Get("estiam"))

	dict.List()

	dict.Remove("ismail")
	dict.List()
}
