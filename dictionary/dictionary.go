package dictionary

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"sync"
)

type Dictionary struct {
	entries map[string]string
	mu      sync.RWMutex
}

// type Dictionary map[string]string

// New crée et renvoie une nouvelle instance de Dictionary.
func New() *Dictionary {
	return &Dictionary{
		entries: make(map[string]string),
	}
}

// Fonction pour charger le dictionnaire depuis le fichier dictionary.txt
func (d *Dictionary) LoadFromFile(filename string) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	//On divise le fichier en ligne
	lines := strings.Split(string(file), "\n")
	for _, line := range lines { //On parcourt chaque ligne et on divise chaque ligne en deux parties séparées par le délimitateur ":"
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			word := strings.TrimSpace(parts[0])
			definition := strings.TrimSpace(parts[1])
			d.entries[word] = definition
		}
	}

	return nil
}

// Fonction pour enregistrer le dictionnaire dans le fichier dictionary.txt
func (d *Dictionary) SaveToFile(filename string) error {
	var lines []string
	for word, definition := range d.entries {
		lines = append(lines, fmt.Sprintf("%s: %s", word, definition))
	}

	data := []byte(strings.Join(lines, "\n"))
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Ajoute une entrée au dictionnaire en utilisant un channel.
func (d *Dictionary) Add(word, definition string, ch chan struct{}) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.entries[word] = definition
	ch <- struct{}{}
}

// Supprime une entrée du dictionnaire en utilisant un channel et enregistre les modifications dans le fichier.
func (d *Dictionary) Remove(word string, ch chan struct{}, filename string) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	// Vérifie si le mot est présent dans le dictionnaire.
	if _, exists := d.entries[word]; !exists {
		return fmt.Errorf("le mot '%s' n'existe pas dans le dictionnaire et ne peut pas être supprimé", word)
	}

	// Supprime le mot du dictionnaire.
	delete(d.entries, word)

	// Enregistre les modifications dans le fichier.
	err := d.SaveToFile(filename)
	if err != nil {
		return err
	}

	// Signale que l'opération est terminée.
	close(ch)

	// Affiche un message pour indiquer que le mot a été supprimé.
	fmt.Printf("Le mot '%s' a été supprimé du dictionnaire\n", word)
	return nil
}

// Get pour afficher la définition spécifique d'un mot
func (d *Dictionary) Get(word string) (string, bool) {
	definition, mottrouve := d.entries[word]
	return definition, mottrouve
}

// Remove pour supprimer un mot du dictionnaire.
// func (d Dictionary) Remove(word string) {
// 	delete(d, word)
// }

// List renvoie une liste triée des mots et de leurs définitions.
func (d *Dictionary) List() []string {
	var result []string
	for word, definition := range d.entries {
		result = append(result, fmt.Sprintf("%s: %s", word, definition))
	}
	sort.Strings(result)
	return result
}
