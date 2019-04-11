// Package command contenant les fonctions de parsing pour interpréter les commandes de l'utilisateur.
package command

import "fmt"
import "errors"

import "github.com/Heisenberk/goshield/structure"
import "github.com/Heisenberk/goshield/crypto"


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
				d.Mode=structure.ENCRYPT
				
			// s'il veut déchiffrer.
			}else if arg[0]=="-d" || arg[0]=="--decrypt"{
				d.Mode=structure.DECRYPT

			// si le mode choisi n'est pas reconnu. 
			}else {
				return nil, errors.New("Mode invalide. ")
			}

			// Détection du mot de passe.
			if arg[1]=="-p" || arg[1]=="--password"{
				d.Password=arg[2]
			}else {
				return nil, errors.New("Aucun mot de passe détecté. ")
			}

			// Enregistrement des fichiers/dossiers à (dé)chiffrer.
			d.Doc = make([]string, len(arg)-3)
			for i:=3; i<len(arg); i++{
				d.Doc[i-3]=arg[i]
				
			}
			return &d, nil
		}

		return nil, errors.New("Commande non reconnue. ")
	}
}

// Interpret permet d'associer la commande à l'action de l'application. 
func Interpret(d  *structure.Documents, err error) {

    // si la commande a été correctement interprétée
    if (err==nil){
        // réalisation d'un chiffrement. 
        if(d.Mode == structure.ENCRYPT){
            crypto.EncryptFileFolder(d)
        }

        // réalisation d'un déchiffrement. 
        if(d.Mode == structure.DECRYPT){
            crypto.DecryptFileFolder(d)
        } 

    // si l'utilisateur ne tape aucun argument, on affiche les commandes. 
    }else if(err.Error()=="Aucun argument. "){

        fmt.Println("");
        fmt.Println("Commande de GoShield : ")
        fmt.Println("");
        fmt.Println("-e/--encrypt : permet de choisir de chiffrer ")
        fmt.Println("-d/--decrypt : permet de choisir de  déchiffrer")
        fmt.Println("-p[password] : permet de taper le mot de passe " )
        fmt.Println("- [liste des fichiers/dossiers] : liste les fichiers/dossiers à chiffrer/déchiffrer")

    // si le chiffrement/déchiffrement rencontre un problème. 
    }else {
        fmt.Println(err)
    }
}




