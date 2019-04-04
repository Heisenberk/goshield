// Package main contenant la fonction principale.
package main

//import "os"
import "fmt"

//import "github.com/Heisenberk/goshield/command"
import "github.com/Heisenberk/goshield/crypto"
//import "os"
//import "errors"
//import "github.com/Heisenberk/goshield/command"
import "github.com/Heisenberk/goshield/structure"



// main représente la fonction principale de GoShield.
func main() {

	var d structure.Documents
	d.Password = "password"
	err := crypto.DecryptFileAES("./env/test6.md.gsh", &d)
	if err != nil {
		fmt.Println(err)
	}

	/*var d structure.Documents
	d.Password = "password"
	err := crypto.EncryptFileAES("./env/test/test6.md", &d)
	if err != nil {
		fmt.Println(err)
	}*/


	/*var d *structure.Documents
	d, err := command.Parse(os.Args[1:])
	fmt.Println(d)

	fmt.Println(err)
	//fmt.Println(err)
	command.Interpret(d,err)*/
}
