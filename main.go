package main

import (
	"fmt"

	"Dictionnaire/dictionary"
)

func main() {

	//Création d'un dictionnaire
	dic := make(dictionary.Dictionary)

	//Ajout de quelques mots et définitions au dictionnaire
	dic.Add("WIFI", "Le Wi-Fi est une technologie sans fil qui permet aux appareils électroniques de se connecter à Internet et de communiquer entre eux")
	dic.Add("Valise", "Une valise est un type de bagage utilisé pour transporter des vêtements et d'autres articles personnels lors de voyages")
	dic.Add("Informatique", "Domaine des concepts et autres techniques employées pour le traitement automatique de l’information")
	dic.Add("Bras", "Membre du corps")

	//Utilisation de Get pour afficher la définition spécifique d'un mot
	mot_a_afficher := "WIFI"
	definition, mottrouve := dic.Get(mot_a_afficher)
	if mottrouve {
		fmt.Printf("Definition de %s: %s\n", mot_a_afficher, definition)
	} else {
		fmt.Printf("%s ne se trouve pas dans le dictionnaire\n", mot_a_afficher)
	}

	// Utilisation de  la méthode Remove pour supprimer un mot du dictionnaire
	motASupprimer := "Bras"
	if _, mottrouve := dic.Get(motASupprimer); mottrouve {
		dic.Remove(motASupprimer)
		fmt.Printf("%s est supprimé du dictionnaire\n", motASupprimer)
	} else {
		fmt.Printf("%s n'est pas dans le dictionnaire, impossible de le supprimer\n", motASupprimer)
	}

	// Appel de la méthode List pour obtenir la liste triée des mots et de leurs définitions
	wordList := dic.List()
	fmt.Println("\nListe triée des mots et leurs définitions :")
	for _, entry := range wordList {
		fmt.Println(entry)
	}
}
