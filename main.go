package main

import (
	"example/dictionnaire/dictionary"
	"fmt"
)

func main() {
	// Spécifiez le chemin du fichier pour le dictionnaire
	filePath := "C:/Users/Lenovo/Desktop/workspace go/dictionnaire/entries.json"

	// Créez une nouvelle instance de Dictionnaire avec le chemin du fichier
	dict := dictionary.NewDictionnaire(filePath)

	err := dict.List()
	if err != nil {
		fmt.Println("Erreur lors de l'affichage initial:", err)
		return
	}

	// Obtenez la définition d'un mot
	mot := "golang"
	definition, err := dict.Get(mot)
	if err != nil {
		fmt.Println("Erreur lors de la récupération:", err)
	} else {
		fmt.Printf("La définition de '%s' est '%s'\n", mot, definition)
	}

	// Supprimez un mot
	motASupprimer := "chat"
	err = dict.Remove(motASupprimer)
	if err != nil {
		fmt.Println("Erreur lors de la suppression:", err)
		return
	}

	// Affichez à nouveau la liste après la suppression
	err = dict.List()
	if err != nil {
		fmt.Println("Erreur lors de l'affichage après suppression:", err)
		return
	}
}
