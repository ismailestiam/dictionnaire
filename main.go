// main.go
package main

import (
	"example/dictionnaire/dictionary"
	"net/http"
)

func main() {
	filePath := "C:/Users/Lenovo/Desktop/workspace go/dictionnaire/entries.json"

	dict := dictionary.NewDictionnaire(filePath)

	// Configure HTTP routes
	http.HandleFunc("/add", dict.AddEntryHandler)
	http.HandleFunc("/remove", dict.RemoveEntryHandler)
	http.HandleFunc("/list", dict.ListEntriesHandler)

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
