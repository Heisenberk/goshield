// Package main contenant la fonction principale.
package main

import "os"
import "fmt"

import "github.com/Heisenberk/goshield/command"
import "github.com/Heisenberk/goshield/structure"

// main représente la fonction principale de GoShield.
func main() {

	var d *structure.Documents
	d, err := command.Parse(os.Args[1:])
	fmt.Println(d)
	command.Interpret(d,err)

}
