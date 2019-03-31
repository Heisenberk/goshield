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

func Interpret( d  *structure.Documents ,err error ) {
	if (err==nil){
		
	}else if(err.Error()=="Aucun argument. "){
	fmt.Println("Commande de l'application")
	fmt.Println("-e/-d")
	fmt.Println("--encrypt : permet de choisir de chiffrer ")
	fmt.Println("--decrypt : permet de choisir de  déchiffrer")
	fmt.Println("-p[password] : permet de taper le mot de passe " )
	fmt.Println("[Liste des fichiers/ dossiers : on liste les fichiers que l'on va chiffrer déchiffrer]")

}
}

// main représente la fonction principale de GoShield.
func main() {
	var d structure.Documents
	d.Password = "password"
	crypto.EncryptFileAES("test5.txt", &d)
	/*var d *structure.Documents
	d, err := command.Parse(os.Args[1:])
	fmt.Println(d)
<<<<<<< HEAD
	fmt.Println(err)*/
	//fmt.Println(err)
	//Interpret(d,err)
}
