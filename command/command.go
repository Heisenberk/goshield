// Package command contenant les fonctions de parsing pour interpréter les commandes de l'utilisateur.
package command

import "fmt"
import "errors"
import "github.com/Heisenberk/goshield/structure"
import "github.com/Heisenberk/goshield/App"

// Parse représente la fonction qui interpréte les commandes de l'utilisateur. 
func Parse(arg []string) (*structure.Documents, error) {

	// si l'utilisateur ne met pas d'arguments.
	if len(arg)==0 {
		return nil, errors.New("Aucun argument. ")

	// si l'utilisateur choisit des paramètres en ligne de commande.
	}else {
		// -e/-d -p password [file1].
		if len(arg)>=4 {
			var d structure.Documents

			// s'il veut chiffrer.
			if arg[0]=="-e" || arg[0]=="--encrypt"{
				fmt.Println("chiffrement")
				d.Mode=structure.ENCRYPT
				
			// s'il veut déchiffrer.
			}else if arg[0]=="-d" || arg[0]=="--decrypt"{
				fmt.Println("dechiffrement")
				d.Mode=structure.DECRYPT

			// si le mode choisi n'est pas reconnu. 
			}else {
				return nil, errors.New("Mode invalide. ")
			}

			// Détection du mot de passe.
			if arg[1]=="-p" || arg[1]=="--password"{
				d.Password=arg[2]
				fmt.Println(d.Password)
			}else {
				return nil, errors.New("Aucun mot de passe détecté. ")
			}

			// Enregistrement des fichiers/dossiers à (dé)chiffrer.
			d.Doc = make([]string, len(arg)-3)
			for i:=3; i<len(arg); i++{
				fmt.Println(arg[i])
				d.Doc[i-3]=arg[i]
				
			}
			return &d, nil
		}

		return nil, errors.New("Commande non reconnue. ")
	}
}
func Interpret( d  *structure.Documents ,err error ) {

	if (err==nil){
		if(d.Mode == 1){
			App.Lister(d.Doc[0])
		}

		
	}else if(err.Error()=="Aucun argument. "){
	fmt.Println("Commande de l'application")
	fmt.Println("-e/-d")
	fmt.Println("--encrypt : permet de choisir de chiffrer ")
	fmt.Println("--decrypt : permet de choisir de  déchiffrer")
	fmt.Println("-p[password] : permet de taper le mot de passe " )
	fmt.Println("[Liste des fichiers/ dossiers : on liste les fichiers que l'on va chiffrer déchiffrer]")

}
}