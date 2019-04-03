package crypto

import "math/rand"
import "time"
import "crypto/aes"
import "crypto/cipher"
import "os"
import "fmt"
import "errors"

import "github.com/Heisenberk/goshield/structure"

func CreateIV() []byte {
	rand.Seed(time.Now().UnixNano())
	iv := make([]byte, 16)
	rand.Read(iv)
	return iv
}

func Xor(a []byte, b []byte) []byte {
	if len(a)!= len(b){
		return nil
	}
	c := make([]byte, len(a))
	for i:=0 ; i< len(a) ;i++ {
		c[i]=a[i]^b[i]
	}
	return c
}

func EncryptBlocAES(iv []byte, key []byte, input []byte) ([]byte, error) {

	// Résultat du chiffrement sera dans output.
	output := make([]byte, aes.BlockSize)

	// Si la taille de l'entrée est invalide on lance une erreur. 
	if len(input)%aes.BlockSize != 0 {
		return output, errors.New("Failure Encryption : Taille du bloc invalide.")
	}

	// Preparation du bloc qui sera chiffré. 
	block, err := aes.NewCipher(key)
	if err != nil {
		return output, errors.New("Failure Encryption : Erreur lors du chiffrement d'un bloc.")
	}

	// Chiffrement AES avec le mode opératoire CBC.
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(output[:aes.BlockSize], input)

	return output, nil
}

func EncryptFileAES(pathFile string, doc *structure.Documents) error{
	// ouverture du fichier a chiffrer
	inputFile, err1 := os.Open(pathFile) 
	if err1 != nil {
		var texteError string = "Failure Encryption : Impossible d'ouvrir le fichier à chiffrer "+pathFile+". "
		return errors.New(texteError)
	}

	stat, err2 := inputFile.Stat()
	if err2 != nil {
  		var texteError string = "Failure Encryption : Impossible de lire le fichier à chiffrer "+pathFile+". "
		return errors.New(texteError)
	}

	//fmt.Printf("The file is %d bytes long", stat.Size())

	var division int = (int)(stat.Size()/aes.BlockSize)
	var iterations int = division
	if (int)(stat.Size())%aes.BlockSize != 0 {
		iterations=iterations+1
	}
	//fmt.Printf("\nnb iterations : %d\n",iterations)

	// ouverture du fichier résultat
    outputFile, err3 := os.Create(pathFile+".gsh")
    if err3 != nil {
  		var texteError string = "Failure Encryption : Impossible d'écrire le fichier chiffré "+pathFile+".gsh. "
		return errors.New(texteError)
	}

	// ecriture de la signature
	_, err4 := outputFile.WriteString("GOSHIELD")
    if err4 != nil {
  		var texteError string = "Failure Encryption : Impossible d'écrire dans le fichier chiffré "+pathFile+".gsh. "
		return errors.New(texteError) 
	}

	// ecriture du salt 
	CreateHash(doc)
	_, err5 := outputFile.Write(doc.Salt)
	if err5 != nil {
  		var texteError string = "Failure Encryption : Impossible de générer le salt. "
		return errors.New(texteError)
	}

	// ecriture de la valeur d'initialisation IV
	IV := CreateIV()
	_, err6 := outputFile.Write(IV)
    if err6 != nil {
  		var texteError string = "Failure Encryption : Impossible d'écrire la valeur d'initialisation IV. "
		return errors.New(texteError) 
	}

	// ecriture de la taille du dernier bloc (sans padding) en octets
	var length int = 0
	if (int)(stat.Size())<aes.BlockSize {
		length=(int)(stat.Size())
	}else {
		length=(int)(stat.Size())%aes.BlockSize
	}
	//fmt.Printf("Taille du dernier bloc chiffré : %d\n", length)
	_, err7 := outputFile.WriteString(fmt.Sprintf("%d", length))
	if err7 != nil {
  		var texteError string = "Failure Encryption : Impossible d'écrire la taille du dernier bloc chiffré. "
		return errors.New(texteError)
	}

	// chiffrement de chaque bloc de données et ecriture des donnees chiffrees
	input := make([]byte, 16)
	var seek int64 = 0
	var cipherBlock []byte
	for i:= 0; i<iterations; i++{
		input =[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

		// lecture de chaque bloc de 16 octets
		_, err8 := inputFile.Read(input)
		if err8 != nil {
  			var texteError string = "Failure Encryption : Impossible de lire dans le fichier à chiffrer "+pathFile+". "
			return errors.New(texteError)
		}
    	//fmt.Println(input)

    	// si on est au tour i (i!=0), IV vaut le chiffré du tour i-1
    	if i != 0 {
    		IV = cipherBlock
    	}

    	seek = aes.BlockSize*((int64)(i+1))
   		_, err9 := outputFile.Seek(seek, 0)
   		if err9 != nil {
  			var texteError string = "Failure Encryption : Impossible de lire dans le fichier à chiffrer "+pathFile+". "
			return errors.New(texteError)
		}

		// chiffrement de chaque bloc et écriture
		var err10 error
		cipherBlock, err10 = EncryptBlocAES(IV, doc.Hash, input)
		if err10 != nil {
			var texteError string = "Failure Encryption : Impossible de chiffrer le fichier "+pathFile+". "
			return errors.New(texteError)
		}
		_, err11 := outputFile.Write(cipherBlock)
		if err11 != nil {
  			var texteError string = "Failure Encryption : Impossible d'écrire dans le fichier "+pathFile+".gsh. "
			return errors.New(texteError)
		}
	}

    outputFile.Close()
    inputFile.Close()
    return nil

}