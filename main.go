package main

import (
	"Dictionnaire/dictionary"
	"encoding/json"
	"fmt"
	"net/http"
)

const filename = "dictionary/dictionary.txt"

// main function
func main() {
	// Création d'un dictionnaire
	dic := dictionary.New()

	// Chargement du dictionnaire depuis le fichier "dictionary.txt"
	err := dic.LoadFromFile(filename)
	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier :", err)
		return
	}

	// Gestion des routes
	// Création des gestionnaires pour différentes routes HTTP
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		handleAdd(dic, w, r)
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		handleGet(dic, w, r)
	})

	http.HandleFunc("/remove", func(w http.ResponseWriter, r *http.Request) {
		handleRemove(dic, w, r)
	})

	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		handleList(dic, w, r)
	})

	// Démarrage du serveur HTTP
	http.ListenAndServe(":8080", nil)
}

// handleAdd gère la route /add
func handleAdd(dic *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { //Vérifier si la méthode HHTP de la requête est POST
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	//Attente d'un corps de requête JSON contenant un mot (Word) et une définition (Definition)
	var entry struct {
		Word       string `json:"word"`
		Definition string `json:"definition"`
	}

	decoder := json.NewDecoder(r.Body) // Décodage du JSON de la requête dans la structure entry
	err := decoder.Decode(&entry)
	if err != nil {
		http.Error(w, "Erreur de décodage JSON", http.StatusBadRequest)
		return
	}

	//Création du canal pour synchroniser l'ajout asynchrone au dictionnaire
	channelAdd := make(chan struct{})
	go func() { //Lancement du goroutine
		dic.Add(entry.Word, entry.Definition, channelAdd, filename)
	}()

	<-channelAdd //Envoi d'un signal sur le canal par le goroutine

	w.WriteHeader(http.StatusCreated) //Réponse HTTP avec le code de statut 201 (StatusCreated) pour indiquer que la ressource a été créée avec succès
}

// handleGet gère la route /get
func handleGet(dic *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { //Vérifier si la méthode HHTP de la requête est GET
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	word := r.URL.Query().Get("word")
	definition, found := dic.Get(word)

	if !found {
		http.Error(w, "Mot non trouvé dans le dictionnaire", http.StatusNotFound)
		return
	}

	response := struct {
		Word       string `json:"word"`
		Definition string `json:"definition"`
	}{Word: word, Definition: definition}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handleRemove gère la route /remove
func handleRemove(dic *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete { //Vérifier si la méthode HHTP de la requête est DELETE
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	word := r.URL.Query().Get("word")
	channelRemove := make(chan struct{})

	go func() {
		err := dic.Remove(word, channelRemove, channelRemove, filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}()

	<-channelRemove
}

// handleList gère la route /list
func handleList(dic *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { //Vérifier si la méthode HHTP de la requête est GET
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	wordList := dic.List()
	response := struct {
		Entries []string `json:"entries"`
	}{Entries: wordList}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GET : http://localhost:8080/get?word=WIFI  (GET)

// LIST : http://localhost:8080/list   (GET)

// DELETE : http://localhost:8080/remove?word=NouveauMot (DELETE)

// http://localhost:8080/add (POST)   {"word":"Ordinateur","definition":"Un ordinateur est une machine électronique capable de recevoir, de traiter et de stocker des données"}
