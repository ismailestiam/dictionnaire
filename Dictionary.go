package dictionary

import "fmt"

type Dictionnaire struct {
	data map[string]string
}

func NewDictionnaire() *Dictionnaire {
	return &Dictionnaire{data: make(map[string]string)}
}

func (d *Dictionnaire) Ajouter(mot, definition string) {
	d.data[mot] = definition
}

func (d *Dictionnaire) Get(mot string) string {
	return d.data[mot]
}

func (d *Dictionnaire) Supprimer(mot string) {
	delete(d.data, mot)
}

func (d *Dictionnaire) Lister() {
	fmt.Println("Listing key-value pairs:")
	for key, value := range d.data {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
