package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Fonction qui permets de clear le cmd
func ClearConsole() {
	var cmd *exec.Cmd

	// Choix de la commande en fonction du système d'exploitation
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	case "linux", "darwin":
		cmd = exec.Command("clear")
	default:
		fmt.Println("Système d'exploitation non pris en charge pour le nettoyage de la console.")
		return
	}

	// Exécution de la commande
	cmd.Stdout = os.Stdout
	cmd.Run()
}
