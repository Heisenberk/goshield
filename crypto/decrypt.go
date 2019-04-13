// Package crypto contenant les fonctions de chiffrement/déchiffrement.
package crypto

import "crypto/aes"
import "crypto/cipher"
import "errors"
import "os"
import "fmt"
import "io/ioutil"
import "strings"

import "github.com/Heisenberk/goshield/structure"

// DecryptBlocAES déchiffre 1 bloc input avec la clé key et la valeur initiale iv pour donner le bloc déchiffré. 
func DecryptBlocAES(iv []byte, key []byte, input []byte) ([]byte, error){

	// Résultat du chiffrement sera dans output.
	output := make([]byte, aes.BlockSize)

	// Si la taille de l'entrée est invalide on lance une erreur. 
	if len(input)%aes.BlockSize != 0 {
		return output, errors.New("\033[31mFailure Decryption\033[0m : Taille du bloc invalide.")
	}

	// Preparation du bloc qui sera chiffré.
	block, err := aes.NewCipher(key)
	if err != nil {
		return output, errors.New("\033[31mFailure Decryption\033[0m : Erreur lors du déchiffrement d'un bloc.")
	}

	// Chiffrement AES avec le mode opératoire CBC.
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(input, input)

	return input, nil
}

// DecryptFileAES déchiffre un fichier de chemin pathFile avec les données doc. 
func DecryptFileAES(pathFile string, doc *structure.Documents) error{
	
	// ouverture du fichier à déchiffrer
	inputFile, err1 := os.Open(pathFile) 
	if err1 != nil {
		var texteError string = "\033[31mFailure Decryption\033[0m : Impossible d'ouvrir le fichier à déchiffrer "+pathFile+". "
		return errors.New(texteError)
	}

	// renvoie une erreur si l'extension n'est pas la bonne
	if pathFile[(len(pathFile)-4):]!= ".gsh"{
		var texteError string = "\033[31mFailure Decryption\033[0m : L'extension de "+pathFile+" est invalide (doit être \".gsh\"). "
		return errors.New(texteError)
	}

	// renvoie une erreur si la signature n'est pas correcte
	signature := make([]byte, 8)
    _, err2 := inputFile.Read(signature)
    if err2 != nil {
		var texteError string = "\033[31mFailure Decryption\033[0m : Format du fichier à déchiffrer "+pathFile+" invalide. "
		return errors.New(texteError)
	}

    // lecture du salt et déduction de la clé
    salt := make([]byte, 15)
    _, err22 := inputFile.Read(salt)
    if err22 != nil {
		var texteError string = "\033[31mFailure Decryption\033[0m : Impossible de lire le salt du fichier chiffré "+pathFile+". "
		return errors.New(texteError)
	}
	doc.Salt=salt
	DeductHash(doc)

	// lecture de la valeur IV
	IV := make([]byte, 16)
	_, err23 := inputFile.Read(IV)
    if err23 != nil {
		var texteError string = "\033[31mFailure Decryption\033[0m : Impossible de lire la valeur d'initialisation du fichier chiffré "+pathFile+". "
		return errors.New(texteError)
	}

	// lecture de la taille du dernier bloc
	lengthTab := make([]byte, 1)
	_, err24 := inputFile.Read(lengthTab)
    if err24 != nil {
		var texteError string = "\033[31mFailure Decryption\033[0m : Impossible de lire la taille du dernier bloc du fichier chiffré "+pathFile+". "
		return errors.New(texteError)
	}

	stat, err2 := inputFile.Stat()
	if err2 != nil {
  		var texteError string = "\033[31mFailure Decryption\033[0m : Impossible d'interpréter le fichier à déchiffrer "+pathFile+". "
		return errors.New(texteError)
	}

	// on soustrait la taille de la signature (8) + le salt (15) + IV (16) + taille du dernier bloc (1)
	var division int = (int)((stat.Size()-8-15-16-1)/aes.BlockSize) 
	var iterations int = division
	if (int)(stat.Size()-8-15-16-1)%aes.BlockSize != 0 {
		var texteError string = "\033[31mFailure Decryption\033[0m : Fichier" + pathFile +" non conforme pour le déchiffrement AES. "
		return errors.New(texteError)
	}

    // ouverture du fichier résultat
    var nameOutput string=pathFile[:(len(pathFile)-4)]
    outputFile, err3 := os.Create(nameOutput)
    if err3 != nil {
  		var texteError string = "\033[31mFailure Decryption\033[0m : Impossible d'écrire le fichier chiffré "+nameOutput+". "
		return errors.New(texteError)
	}

	input := make([]byte, 16)
	var cipherBlock []byte
	temp := make([]byte, 16)
	
	for i:=0 ; i<iterations ; i++ {

    	// si on est au tour i (i!=0), IV vaut le chiffré du tour i-1
    	if (i) != 0 {
    		IV = temp
    	}

		input =[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

		// lecture de chaque bloc de 16 octets
		_, err8 := inputFile.Read(input)
		if err8 != nil {
  			var texteError string = "\033[31mFailure Decryption\033[0m : Impossible de lire dans le fichier à déchiffrer "+pathFile+". "
			return errors.New(texteError)
		}

    	temp =[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		copy(temp, input)
		
		// déchiffrement de chaque bloc et écriture
		var err10 error
		cipherBlock, err10 = DecryptBlocAES(IV, doc.Hash, input)
		if err10 != nil {
			var texteError string = "\033[31mFailure Decryption\033[0m : Impossible de déchiffrer le fichier "+pathFile+". "
			return errors.New(texteError)
		}
		
		// dans le dernier bloc, il faut enlever les bits de padding qui ne sont pas dans le message initial.
		if i==(iterations-1) {
			if lengthTab[0]!= 0 {
				_, err11 := outputFile.Write(cipherBlock[:lengthTab[0]])
				if err11 != nil {
  					var texteError string = "\033[31mFailure Decryption\033[0m : Impossible d'écrire dans le fichier "+nameOutput+". "
					return errors.New(texteError)
				}
			}else {
				_, err12 := outputFile.Write(cipherBlock)
				if err12 != nil {
  					var texteError string = "\033[31mFailure Decryption\033[0m : Impossible d'écrire dans le fichier "+nameOutput+". "
					return errors.New(texteError)
				}
			}
			
			
		}else {
			_, err13 := outputFile.Write(cipherBlock)
			if err13 != nil {
  				var texteError string = "\033[31mFailure Decryption\033[0m : Impossible d'écrire dans le fichier "+nameOutput+". "
				return errors.New(texteError)
			}
		}
	}

	// fermeture des fichiers. 
	inputFile.Close()
	outputFile.Close()

	var messageSuccess string = "- \033[32mSuccess Decryption\033[0m : "+pathFile+" : resultat dans le fichier "+nameOutput
    fmt.Println(messageSuccess)

	return nil
}

// DecryptFolder déchiffre le contenu d'un dossier de chemin path avec les données doc. 
func DecryptFolder (path string, d *structure.Documents) {

    // Lecture du chemin à déchiffrer. 
   entries, err := ioutil.ReadDir(path)
    if err != nil {
        fmt.Println("- \033[31mFailure Decryption\033[0m : impossible d'ouvrir "+path)
    }

    // Déchiffrement de chaque élément du dossier. 
    for _, entry := range entries {

        newPath := path+entry.Name()
        fi, err := os.Stat(newPath)
        valid := true
        if err != nil {
            fmt.Println("- \033[31mFailure Decryption\033[0m : "+newPath+" n'existe pas ")
            valid = false
        }

        // si l'élément du dossier existe. 
        if valid == true {

            mode := fi.Mode();

            //si l'objet spécifié par le chemin est un dossier.
            if(mode.IsDir()==true){

                //Si l'utilisateur a oublié le "/" à la fin du chemin du fichier
                if(strings.LastIndexAny(newPath, "/") != len(newPath) - 1){
                    newPath=newPath+ string(os.PathSeparator)
                      
                }
                DecryptFolder(newPath, d)

            // si l'objet spécifié par le chemin est un fichier.
            }else if mode.IsRegular()== true {

                // si l'extension du fichier est différent de .gsh on peut chiffrer le fichier.
                if newPath[len(newPath)-4:]==".gsh"{
                	errFile := DecryptFileAES(newPath,d)
                    if errFile != nil {
                        fmt.Println(errFile)
                    }
                }     
            }
        }
    }
}

// DecryptFileFolder déchiffre les éléments choisis par l'utilisateur avec les données doc. 
func DecryptFileFolder(d *structure.Documents) {

	// Pour chaque élément choisi par l'utilisateur. 
    for i := 0; i < len(d.Doc); i++ {

    	// Ouverture de l'élément. 
        fi, err := os.Stat(d.Doc[i])
        valid := true
        if err != nil {
            fmt.Println("- \033[31mFailure Decryption\033[0m : "+d.Doc[i]+" n'existe pas ")
            valid = false
        }

        // Si l'élément est valide. 
        if valid == true {
            mode := fi.Mode();

            // Si l'élément spécifié par le chemin est un dossier.
            if(mode.IsDir()==true){

                // Si l'utilisateur a oublié le "/" à la fin du chemin du fichier.
                if(strings.LastIndexAny(d.Doc[i], "/") != len(d.Doc[i]) - 1){
                  d.Doc[i]=d.Doc[i]+ string(os.PathSeparator)
                }

                // Déchiffrement du dossier.
                DecryptFolder(d.Doc[i], d)

            // si l'objet spécifié par le chemin est un fichier.
            }else if mode.IsRegular()== true {

            	if d.Doc[i][len(d.Doc[i])-4:]==".gsh"{
            		// Déchiffrement du fichier. 
                	errFile := DecryptFileAES(d.Doc[i],d)
                	if errFile != nil {
                   		fmt.Println(errFile)
                	}
            	}
            }
        }
    }
}

	
