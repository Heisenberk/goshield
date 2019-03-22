package main

import "os"
import "fmt"

import "github.com/Heisenberk/goshield/command"
import "github.com/Heisenberk/goshield/structure"


func main() {
	var d *structure.Documents
	d, err := command.Parse(os.Args[1:])
	fmt.Println(d)
	fmt.Println(err)
}
