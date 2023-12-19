package dictionary

import (
	"fmt"
	"sort"
)

type Dictionary map[string]string

// Add ajoute un mot et sa définition au dictionnaire.
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

// Get pour afficher la définition spécifique d'un mot
func (d Dictionary) Get(word string) (string, bool) {
	definition, mottrouve := d[word]
	return definition, mottrouve
}

// Remove pour supprimer un mot du dictionnaire.
func (d Dictionary) Remove(word string) {
	delete(d, word)
}

// List renvoie une liste triée des mots et de leurs définitions.
func (d Dictionary) List() []string {
	var result []string
	for word, definition := range d {
		result = append(result, fmt.Sprintf("%s: %s", word, definition))
	}
	sort.Strings(result)
	return result
}
