// Package command contenant les fonctions de parsing pour interpréter les commandes de l'utilisateur.
package command

import (

    "fmt"
    "io/ioutil"
    "strings"
    "os"


)
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
//Met à jour le chemin à chaque fois que l'on rencontre un dossier,
//Pour cela, on ajoute le nom du dossier dans le chemin (= dans le cas du chiffrement d'un dossier)
func Maj_Chemin_Crypt(path string,name string,isdir bool, d  *structure.Documents){
        //si l'objet spécifié par le chemin est un dossier
        if(isdir ){

        path =path+name+"/"
        Chiffre_DossierOuFichier(path,d)
        path=strings.TrimRight(path,"/")
        path=strings.TrimRight(path,name)

    }
}
//Met à jour le chemin à chaque fois que l'on rencontre un dossier,
//Pour cela, on ajoute le nom du dossier dans le chemin (= dans le cas du déchiffrement d'un dossier)
func Maj_Chemin_Decypt(path string,name string,isdir bool, d  *structure.Documents){
        //si l'objet spécifié par le chemin est un dossier
        if(isdir ){

        path =path+name+"/"
        Dechiffre_DossierOuFichier(path,d)
        path=strings.TrimRight(path,"/")
        path=strings.TrimRight(path,name)

    }
}

//Chiffre le contenu d'un document si l'objet spécifié par le chemin est un dossier
//sinon
//l'objet spécifié par le chemin est un fichier alors on chiffre ce fichier 
func Chiffre_DossierOuFichier(path string,d *structure.Documents){
    
    fi, err := os.Stat(path)
    
    if err != nil {
        fmt.Println(err)
        return
    }
    mode := fi.Mode();
    //si l'objet spécifié par le chemin est un dossier
    if(mode.IsDir()==true){
        //Si l'utilisateur a oublié le "/" à la fin du chemin du fichier
    	if(strings.LastIndexAny(path, "/") != len(path) - 1){

            path=path+ string(os.PathSeparator)
    }
    //On lit dans le dossier visée par le chemin
   entries, err := ioutil.ReadDir(path)

    if err != nil {

        fmt.Println(err)
    }
    for _, entry := range entries {

    	p:=path + entry.Name()
        // si l'extension du fichier est différent de .gsh on peut chiffrer le fichier
    	if(p[len(p)-4:]!=".gsh"){

    	   crypto.EncryptFileAES(path+entry.Name(),d)
    	   fmt.Println(path+entry.Name())
    	}
        //Si on tombe sur un dossier on met à jour le chemin
       Maj_Chemin_Crypt(path,entry.Name(),entry.IsDir(),d)

    }
    //si l'objet spécifié par le chemin est un fichier
    }else if(mode.IsRegular()==true){

     crypto.EncryptFileAES(getname_file(path),d)
    }

}
//Déchiffre le contenu d'un document si l'objet spécifié par le chemin est un dossier
//sinon
//l'objet spécifié par le chemin est un fichier alors on déchiffre ce fichier 
func Dechiffre_DossierOuFichier(path string,d *structure.Documents){
    
    fi, err := os.Stat(path)
    
    if err != nil {
        fmt.Println(err)
        return
    }
    mode := fi.Mode();
    //si l'objet spécifié par le chemin est un dossier
    if(mode.IsDir()==true){
        //Si l'utilisateur a oublié le "/" à la fin du chemin du fichier
    	if(strings.LastIndexAny(path, "/") != len(path) - 1){

            path=path+ string(os.PathSeparator)
    }
    
   entries, err := ioutil.ReadDir(path)

    if err != nil {

        fmt.Println(err)
    }
        for _, entry := range entries {
            //on déchiffrer le fichier spécifié par le chemin
    	   crypto.DecryptFileAES(path+entry.Name(),d)
    	   fmt.Println(path+entry.Name())
           //Si on tombe sur un dossier on met à jour le chemin
            Maj_Chemin_Decypt(path,entry.Name(),entry.IsDir(),d)

        }
    //si l'objet spécifié par le chemin est un fichier
    }else if(mode.IsRegular()==true){

        crypto.DecryptFileAES(getname_file(path),d)

    }
      
}
//Si l'utilisateur a oublié le "/" à la fin du chemin du fichier
func getname_file(path string) string {

    path =strings.Trim(path,"/")
    path=path[strings.LastIndexAny(path, "/")+1:]

    return path
}

func Interpret( d  *structure.Documents ,err error ) {


	if (err==nil){
		if(d.Mode == 1){

			for i := 1; i < len(d.Doc); i++ {
			     Chiffre_DossierOuFichier(d.Doc[i],d)
			}
			
		}
		if(d.Mode == 2){
			for i := 0; i < len(d.Doc); i++ {
				Dechiffre_DossierOuFichier(d.Doc[i],d)
			}
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