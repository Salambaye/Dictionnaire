package main

import (
	"fmt"

	"Dictionnaire/dictionary"
	"time"
)

const filename = "dictionary/dictionary.txt"

func main() {

	//Création d'un dictionnaire
	dic := dictionary.New()

	//Chargement du dictionnaire depuisle fichier "dictionary.txt"
	err := dic.LoadFromFile(filename)
	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier :", err)
		return
	}

	//Utilisation de Get pour afficher la définition spécifique d'un mot
	mot_a_afficher := "WIFI"
	definition, mottrouve := dic.Get(mot_a_afficher)
	if mottrouve {
		fmt.Printf("Definition de %s: %s\n", mot_a_afficher, definition)
	} else {
		fmt.Printf("%s ne se trouve pas dans le dictionnaire\n", mot_a_afficher)
	}

	// Utilisation de la méthode Remove pour supprimer un mot du dictionnaire en utilisant un channel
	channelRemove := make(chan struct{})
	go func() {
		err := dic.Remove("Bras", channelRemove, filename)
		if err != nil {
			fmt.Println(err)
		}
	}()

	// Utilisation de la méthode Add pour ajouter un mot au dictionnaire en utilisant un channel
	channelAdd := make(chan struct{})
	go func() {
		dic.Add("Bras", "Définition du nouveau mot", channelAdd)
	}()
	// Attente des signaux des channels pour garantir que les opérations sont terminées
	select {
	case <-channelRemove:
		// La suppression est terminée
	case <-time.After(2 * time.Second):
		// Si la suppression prend trop de temps, passe à l'étape suivante
		fmt.Println("Le délai d'attente pour la suppression est écoulé")
	}

	select {
	case <-channelAdd:
		// L'ajout est terminé
	case <-time.After(2 * time.Second):
		// Si l'ajout prend trop de temps, passe à l'étape suivante
		fmt.Println("Le délai d'attente pour l'ajout est écoulé")
	}

	// Utilisation de  la méthode Remove pour supprimer un mot du dictionnaire
	// motASupprimer := "Bras"
	// if _, mottrouve := dic.Get(motASupprimer); mottrouve {
	// 	dic.Remove(motASupprimer)
	// 	fmt.Printf("%s est supprimé du dictionnaire\n", motASupprimer)

	// 	// Sauvegarde des modifications dans le fichier texte
	// 	err := dic.SaveToFile(filename)
	// 	if err != nil {
	// 		fmt.Println("Erreur lors de la sauvegarde du fichier :", err)
	// 	}
	// } else {
	// 	fmt.Printf("%s n'est pas dans le dictionnaire, impossible de le supprimer\n", motASupprimer)
	// }

	// Appel de la méthode List pour obtenir la liste triée des mots et de leurs définitions
	wordList := dic.List()
	fmt.Println("\nListe triée des mots et leurs définitions :")
	for _, entry := range wordList {
		fmt.Println(entry)
	}
}
