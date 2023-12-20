package dictionary

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Dictionary map[string]string

// New crée et renvoie une nouvelle instance de Dictionary.
func New() Dictionary {
	return make(Dictionary)
}

// Fonction pour charger le dictionnaire depuis le fichier dictionary.txt
func (d Dictionary) LoadFromFile(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			word := strings.TrimSpace(parts[0])
			definition := strings.TrimSpace(parts[1])
			d[word] = definition
		}
	}

	return nil
}

// Fonction pour enregistrer le dictionnaire dans le fichier dictionary.txt
func (d Dictionary) SaveToFile(filename string) error {
	var lines []string
	for word, definition := range d {
		lines = append(lines, fmt.Sprintf("%s : %s", word, definition))
	}

	data := []byte(strings.Join(lines, "\n"))
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
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
