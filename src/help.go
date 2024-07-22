package main

import (
	"flag"
	"fmt"
	"os"
)

// cette fonction permets d'afficher les aides essentiels pour le jeu
func help() {
	helpFlag := flag.Bool("h", false, "Affiche les explications")

	flag.Parse()

	// Vérification de l'option -h
	if *helpFlag {
		printHelp()
		os.Exit(0)
	}
}

// fonction qui permets d'afficher les aides
func printHelp() {
	ClearConsole()
	fmt.Println("C LE TIRET H MEC ! 🎉✨")
}
