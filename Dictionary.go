// dictionary.go
package dictionary

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
)

type Entry struct {
	Mot        string `json:"mot"`
	Definition string `json:"definition"`
}

type Dictionnaire struct {
	filePath string
	entries  []Entry
	mutex    sync.Mutex
}

func NewDictionnaire(filePath string) *Dictionnaire {
	return &Dictionnaire{
		filePath: filePath,
		entries:  nil,
	}
}

func (d *Dictionnaire) loadFromFile() error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	fileData, err := ioutil.ReadFile(d.filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileData, &d.entries)
	if err != nil {
		return err
	}

	return nil
}

func (d *Dictionnaire) saveToFile() error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	fileData, err := json.MarshalIndent(d.entries, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(d.filePath, fileData, 0644)
	if err != nil {
		return err
	}

	return nil
}

// AddEntryHandler handles the HTTP POST request to add an entry to the dictionary.
func (d *Dictionnaire) AddEntryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newEntry Entry
	err := json.NewDecoder(r.Body).Decode(&newEntry)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.entries = append(d.entries, newEntry)

	err = d.saveToFile()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// RemoveEntryHandler handles the HTTP DELETE request to remove an entry from the dictionary by word.
func (d *Dictionnaire) RemoveEntryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	word := r.URL.Query().Get("mot")
	if word == "" {
		http.Error(w, "Missing 'mot' parameter", http.StatusBadRequest)
		return
	}

	d.mutex.Lock()
	defer d.mutex.Unlock()

	var updatedEntries []Entry
	found := false
	for _, entry := range d.entries {
		if entry.Mot == word {
			found = true
		} else {
			updatedEntries = append(updatedEntries, entry)
		}
	}

	if !found {
		http.Error(w, "Mot non trouvé", http.StatusNotFound)
		return
	}

	d.entries = updatedEntries

	err := d.saveToFile()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ListEntriesHandler handles the HTTP GET request to list all entries in the dictionary.
func (d *Dictionnaire) ListEntriesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := d.loadFromFile()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(d.entries)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
