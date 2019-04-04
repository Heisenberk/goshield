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

	/*iv:= []byte {36,31,225, 245, 191, 97, 106, 191, 99, 77, 129, 76, 26, 190, 87, 210}
	key:=[]byte{221, 91, 116, 74, 157, 205, 95, 243, 70, 54, 127, 92, 228, 182, 231, 250, 24, 218, 112, 62, 5, 92, 209, 96, 143, 79, 23, 101, 86, 183, 149, 61}
	input:=[]byte{22, 249, 179, 58, 118, 104, 69, 215, 203, 66, 204, 255, 210, 32, 175, 2}
	//var ouput []byte
	output, _:= crypto.DecryptBlocAES(iv , key , input )
	fmt.Println(output)*/


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
<<<<<<< HEAD
	fmt.Println(err)*/
	//fmt.Println(err)
	//Interpret(d,err)
}
