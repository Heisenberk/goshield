package command

import "fmt"
import "errors"
import "github.com/Heisenberk/goshield/structure"

const ENCRYPT int = 1
const DECRYPT int = 2


func Parse(arg []string) (*structure.Documents, error) {

	// si l'utilisateur ne met pas d'arguments
	if len(arg)==0 {
		return nil, errors.New("Aucun argument. ")

	// si l'utilisateur choisit des paramètres en ligne de commande
	}else {
		// -c/-d -p password [file1]
		if len(arg)>=4 {
			var d structure.Documents

			// s'il veut chiffrer
			if arg[0]=="-e" || arg[0]=="--encrypt"{
				fmt.Println("chiffrement")
				d.Mode=ENCRYPT
				
			// s'il veut déchiffrer
			}else if arg[0]=="-d" || arg[0]=="--decrypt"{
				fmt.Println("dechiffrement")
				d.Mode=DECRYPT

			// si le mode choisi n'est pas reconnu. 
			}else {
				return nil, errors.New("Mode invalide. ")
			}

			// Détection du mot de passe
			if arg[1]=="-p" || arg[1]=="--password"{
				d.Password=arg[2]
				fmt.Println(d.Password)
			}else {
				return nil, errors.New("Aucun mot de passe détecté. ")
			}

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

func Test() string {
	return "coucou"
}