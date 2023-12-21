package main

import (
	"Dictionnaire/dictionary"
	"fmt"
	"time"
)

const filename = "dictionary/dictionary.txt"

func main() {
	// Création d'un dictionnaire
	dic := dictionary.New()

	// Chargement du dictionnaire depuis le fichier "dictionary.txt"
	err := dic.LoadFromFile(filename)
	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier :", err)
		return
	}

	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------")

	// Utilisation de Get pour afficher la définition spécifique d'un mot
	mot_a_afficher := "WIFI"
	definition, mottrouve := dic.Get(mot_a_afficher)
	if mottrouve {
		fmt.Printf("Definition de %s: %s\n", mot_a_afficher, definition)
	} else {
		fmt.Printf("%s ne se trouve pas dans le dictionnaire\n", mot_a_afficher)
	}

	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------")

	// Utilisation de la méthode Remove pour supprimer un mot du dictionnaire en utilisant un channel
	channelSync := make(chan struct{})

	go func() {
		defer close(channelSync)

		// Suppression d'un mot du dictionnaire
		err := dic.Remove("Mangue", channelSync, channelSync, filename)
		if err != nil {
			fmt.Println(err)
		}

		// Ajout d'un nouveau mot au dictionnaire
		dic.Add("Voiture", "Véhicule", channelSync, filename)

		// Chargement du dictionnaire depuis le fichier après les modifications
		err = dic.LoadFromFile(filename)
		if err != nil {
			fmt.Println("Erreur lors du chargement du fichier :", err)
			return
		}
	}()

	// Attente de la fin de toutes les opérations avec un délai supplémentaire
	select {
	case <-channelSync:
		// Toutes les opérations sont terminées
	case <-time.After(5 * time.Second):
		// Si cela prend trop de temps, passe à l'étape suivante
		fmt.Println("Le délai d'attente est écoulé")
	}
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------")

	// Appel de la méthode List pour obtenir la liste triée des mots et de leurs définitions
	wordList := dic.List()
	fmt.Println("\nListe triée des mots et leurs définitions :")
	for _, entry := range wordList {
		fmt.Println(entry)
	}

	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------")

}
