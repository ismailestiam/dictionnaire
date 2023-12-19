package main

import (
	"example/dictionnaire/dictionary"
	"fmt"
)

func main() {
	// Utilisation de la classe Dictionnaire
	dict := dictionary.NewDictionnaire()

	dict.Ajouter("estiam", "ecole")
	dict.Ajouter("ismail", "etudiant")

	fmt.Println("Get 'estiam':", dict.Get("estiam"))

	dict.Lister()

	dict.Supprimer("ismail")
	dict.Lister()
}
