package crypto

import "crypto/aes"
import "crypto/cipher"
import "errors"
import "os"
import "fmt"

import "github.com/Heisenberk/goshield/structure"

// DecryptBlocAES déchiffre 1 bloc input avec la clé key et la valeur initiale iv pour donner le bloc déchiffré. 
func DecryptBlocAES(iv []byte, key []byte, input []byte) ([]byte, error){

	// Résultat du chiffrement sera dans output.
	output := make([]byte, aes.BlockSize)

	// Si la taille de l'entrée est invalide on lance une erreur. 
	if len(input)%aes.BlockSize != 0 {
		return output, errors.New("Failure Decryption : Taille du bloc invalide.")
	}

	// Preparation du bloc qui sera chiffré.
	block, err := aes.NewCipher(key)
	if err != nil {
		return output, errors.New("Failure Decryption : Erreur lors du déchiffrement d'un bloc.")
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
		var texteError string = "Failure Decryption : Impossible d'ouvrir le fichier à déchiffrer "+pathFile+". "
		return errors.New(texteError)
	}

	// renvoie une erreur si l'extension n'est pas la bonne
	if pathFile[(len(pathFile)-4):]!= ".gsh"{
		var texteError string = "Failure Decryption : L'extension de "+pathFile+" est invalide (doit être \".gsh\"). "
		return errors.New(texteError)
	}

	// renvoie une erreur si la signature n'est pas correcte
	signature := make([]byte, 8)
    _, err2 := inputFile.Read(signature)
    if err2 != nil {
		var texteError string = "Failure Decryption : Format du fichier à déchiffrer "+pathFile+" invalide. "
		return errors.New(texteError)
	}

    // lecture du salt et déduction de la clé
    salt := make([]byte, 15)
    _, err22 := inputFile.Read(salt)
    if err22 != nil {
		var texteError string = "Failure Decryption : Impossible de lire le salt du fichier chiffré "+pathFile+". "
		return errors.New(texteError)
	}
	doc.Salt=salt
	DeductHash(doc)

	// lecture de la valeur IV
	IV := make([]byte, 16)
	_, err23 := inputFile.Read(IV)
    if err23 != nil {
		var texteError string = "Failure Decryption : Impossible de lire la valeur d'initialisation du fichier chiffré "+pathFile+". "
		return errors.New(texteError)
	}

	// lecture de la taille du dernier bloc
	lengthTab := make([]byte, 1)
	_, err24 := inputFile.Read(lengthTab)
    if err24 != nil {
		var texteError string = "Failure Decryption : Impossible de lire la taille du dernier bloc du fichier chiffré "+pathFile+". "
		return errors.New(texteError)
	}

	stat, err2 := inputFile.Stat()
	if err2 != nil {
  		var texteError string = "Failure Decryption : Impossible d'interpréter le fichier à déchiffrer "+pathFile+". "
		return errors.New(texteError)
	}

	// on soustrait la taille de la signature (8) + le salt (15) + IV (16) + taille du dernier bloc (1)
	var division int = (int)((stat.Size()-8-15-16-1)/aes.BlockSize) 
	var iterations int = division
	if (int)(stat.Size()-8-15-16-1)%aes.BlockSize != 0 {
		var texteError string = "Failure Decryption : Fichier" + pathFile +" non conforme pour le déchiffrement AES. "
		return errors.New(texteError)
	}

    // ouverture du fichier résultat
    var nameOutput string=pathFile[:(len(pathFile)-4)]
    outputFile, err3 := os.Create(nameOutput)
    if err3 != nil {
  		var texteError string = "Failure Decryption : Impossible d'écrire le fichier chiffré "+nameOutput+". "
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
  			var texteError string = "Failure Decryption : Impossible de lire dans le fichier à déchiffrer "+pathFile+". "
			return errors.New(texteError)
		}

    	temp =[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		copy(temp, input)
		
		// déchiffrement de chaque bloc et écriture
		var err10 error
		cipherBlock, err10 = DecryptBlocAES(IV, doc.Hash, input)
		if err10 != nil {
			var texteError string = "Failure Decryption : Impossible de déchiffrer le fichier "+pathFile+". "
			return errors.New(texteError)
		}
		
		// dans le dernier bloc, il faut enlever les bits de padding qui ne sont pas dans le message initial.
		if i==(iterations-1) {
			_, err11 := outputFile.Write(cipherBlock[:lengthTab[0]])
			if err11 != nil {
  				var texteError string = "Failure Decryption : Impossible d'écrire dans le fichier "+nameOutput+". "
				return errors.New(texteError)
			}
		}else {
			_, err11 := outputFile.Write(cipherBlock)
			if err11 != nil {
  				var texteError string = "Failure Decryption : Impossible d'écrire dans le fichier "+nameOutput+". "
				return errors.New(texteError)
			}
		}
	}

	// fermeture des fichiers. 
	inputFile.Close()
	outputFile.Close()

	var messageSuccess string = "- Success Decryption "+pathFile+" : resultat dans le fichier "+nameOutput
    fmt.Println(messageSuccess)

	return nil
}

	
