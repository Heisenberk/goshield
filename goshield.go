// Package main contenant la fonction principale.
package main

import "os"

import "github.com/Heisenberk/goshield/command"
import "github.com/Heisenberk/goshield/structure"

// main représente la fonction principale de GoShield.
func main() {

	// Lecture de la demande de l'utilisateur. 
	var d *structure.Documents
	d, err := command.Parse(os.Args[1:])

	// Interprétation de l'application en réponse à la demande de la commande de l'utilisateur. 
	command.Interpret(d,err)

}
