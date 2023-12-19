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

	fmt.Println("Obtenir 'estiam':", dict.Obtenir("estiam"))

	dict.Lister()

	dict.Supprimer("ismail")
	dict.Lister()
}
