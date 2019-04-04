package crypto

import "crypto/aes"
import "crypto/cipher"
import "errors"
import "os"
import "fmt"

import "github.com/Heisenberk/goshield/structure"

func DecryptBlocAES(iv []byte, key []byte, input []byte) ([]byte, error){
	
fmt.Printf("\n");
    	fmt.Printf("input=")
    	fmt.Println(input)
    	fmt.Printf("IV changé tt seul=")
    	fmt.Println(iv)

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

	
    	
    	fmt.Printf("KEY=")
    	fmt.Println(key)
    	fmt.Printf("output=")
    	fmt.Println(input)

	return input, nil
}

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
    fmt.Printf(": %s\n", string(signature))

    // lecture du salt
    salt := make([]byte, 15)
    _, err22 := inputFile.Read(salt)
    if err22 != nil {
		var texteError string = "Failure Decryption : Impossible de lire le salt du fichier chiffré "+pathFile+". "
		return errors.New(texteError)
	}
	fmt.Printf("Salt : ")
	doc.Salt=salt
	fmt.Println(salt)

	DeductHash(doc)
	fmt.Printf("Hash: ")
	fmt.Println(doc.Hash)

	// lecture de la valeur IV
	IV := make([]byte, 16)
	_, err23 := inputFile.Read(IV)
    if err23 != nil {
		var texteError string = "Failure Decryption : Impossible de lire la valeur d'initialisation du fichier chiffré "+pathFile+". "
		return errors.New(texteError)
	}
	fmt.Printf("IV: ")
	fmt.Println(IV)

	// lecture de la taille du dernier bloc
	lengthTab := make([]byte, 1)
	_, err24 := inputFile.Read(lengthTab)
    if err24 != nil {
		var texteError string = "Failure Decryption : Impossible de lire la taille du dernier bloc du fichier chiffré "+pathFile+". "
		return errors.New(texteError)
	}

	fmt.Printf("Taille du dernier bloc chiffré : ")
	fmt.Println(lengthTab[0])

	stat, err2 := inputFile.Stat()
	if err2 != nil {
  		var texteError string = "Failure Decryption : Impossible d'interpréter le fichier à déchiffrer "+pathFile+". "
		return errors.New(texteError)
	}

	fmt.Printf("The file is %d bytes long\n", stat.Size())

	// on soustrait la taille de la signature (8) + le salt (15) + IV (16) + taille du dernier bloc (1)
	var division int = (int)((stat.Size()-8-15-16-1)/aes.BlockSize) 
	var iterations int = division
	if (int)(stat.Size()-8-15-16-1)%aes.BlockSize != 0 {
		var texteError string = "Failure Decryption : Fichier" + pathFile +" non conforme pour le déchiffrement AES. "
		return errors.New(texteError)
	}
	fmt.Printf("Nb iterations : %d\n",iterations)

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
		fmt.Println("-------------------------------------------------")

		fmt.Printf("IV a choisir=")
    	fmt.Println(temp)

    	// si on est au tour i (i!=0), IV vaut le chiffré du tour i-1
    	if (i) != 0 {
    		IV = temp
    		fmt.Printf("IV choisi=")
    		fmt.Println(IV)
    	}

		input =[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

		// lecture de chaque bloc de 16 octets
		_, err8 := inputFile.Read(input)
		if err8 != nil {
  			var texteError string = "Failure Decryption : Impossible de lire dans le fichier à déchiffrer "+pathFile+". "
			return errors.New(texteError)
		}

		fmt.Printf("IV choisi2=")
    		fmt.Println(IV)

    	temp =[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		copy(temp, input)
    	


    	
fmt.Printf("IV choisi3=")
    		fmt.Println(IV)
		
		// déchiffrement de chaque bloc et écriture
		var err10 error
		cipherBlock, err10 = DecryptBlocAES(IV, doc.Hash, input)
		if err10 != nil {
			var texteError string = "Failure Decryption : Impossible de déchiffrer le fichier "+pathFile+". "
			return errors.New(texteError)
		}
		_, err11 := outputFile.Write(cipherBlock)
		if err11 != nil {
  			var texteError string = "Failure Decryption : Impossible d'écrire dans le fichier "+nameOutput+". "
			return errors.New(texteError)
		}

			_, err13 := outputFile.Write(cipherBlock)
			if err13 != nil {
  				var texteError string = "Failure Decryption : Impossible d'écrire dans le fichier "+nameOutput+". "
				return errors.New(texteError)
			}
		

		
    	
		
    	
	}


	inputFile.Close()
	outputFile.Close()

	return nil
}

	
