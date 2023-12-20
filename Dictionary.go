package dictionary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Entry struct {
	Mot        string `json:"mot"`
	Definition string `json:"definition"`
}

type Dictionnaire struct {
	filePath string
	entries  []Entry
}

func NewDictionnaire(filePath string) *Dictionnaire {
	return &Dictionnaire{
		filePath: filePath,
		entries:  nil,
	}
}

func (d *Dictionnaire) loadFromFile() error {
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

func (d *Dictionnaire) Add(mot, definition string) error {
	err := d.loadFromFile()
	if err != nil {
		return err
	}

	newEntry := Entry{Mot: mot, Definition: definition}
	d.entries = append(d.entries, newEntry)

	err = d.saveToFile()
	if err != nil {
		return err
	}

	return nil
}

func (d *Dictionnaire) Get(mot string) (string, error) {
	err := d.loadFromFile()
	if err != nil {
		return "", err
	}

	for _, entry := range d.entries {
		if entry.Mot == mot {
			return entry.Definition, nil
		}
	}

	return "", fmt.Errorf("Mot non trouvé : %s", mot)
}

func (d *Dictionnaire) Remove(mot string) error {
	err := d.loadFromFile()
	if err != nil {
		return err
	}

	var updatedEntries []Entry
	found := false
	for _, entry := range d.entries {
		if entry.Mot == mot {
			found = true
		} else {
			updatedEntries = append(updatedEntries, entry)
		}
	}

	if !found {
		return fmt.Errorf("Mot non trouvé : %s", mot)
	}

	d.entries = updatedEntries

	err = d.saveToFile()
	if err != nil {
		return err
	}

	return nil
}

func (d *Dictionnaire) List() error {
	err := d.loadFromFile()
	if err != nil {
		return err
	}

	fmt.Println("Listing key-value pairs:")
	for _, entry := range d.entries {
		fmt.Printf("Mot: %s, Definition: %s\n", entry.Mot, entry.Definition)
	}

	return nil
}
