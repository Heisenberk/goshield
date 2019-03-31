package crypto

import "math/rand"
import "time"
import "crypto/aes"
import "crypto/cipher"
import "os"
import "fmt"

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

func EncryptBlocAES(iv []byte, key []byte, input []byte) []byte {

	// Si la taille de l'entrée est invalide on lance une erreur. 
	if len(input)%aes.BlockSize != 0 {
		panic("EXCEPTION A LANCER") 
	}

	// Preparation du bloc qui sera chiffré. 
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Résultat du chiffrement sera dans output.
	output := make([]byte, aes.BlockSize)

	// Chiffrement AES avec le mode opératoire CBC.
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(output[:aes.BlockSize], input)

	return output
}

func EncryptFileAES(pathFile string, doc *structure.Documents){
	// ouverture du fichier a chiffrer
	file, err := os.Open(pathFile) 
	if err != nil {
		panic("EXCEPTION A LANCER1") 
	}

	stat, err := file.Stat()
	if err != nil {
  		panic("EXCEPTION A LANCER2") 
	}

	fmt.Printf("The file is %d bytes long", stat.Size())

	var division int = (int)(stat.Size()/aes.BlockSize)
	var iterations int = division
	if (int)(stat.Size())%aes.BlockSize != 0 {
		iterations=iterations+1
	}
	fmt.Printf(">%d\n",iterations)

	// ouverture du fichier résultat
    f, err := os.Create(pathFile+".gsh")
    if err != nil {
  		panic("EXCEPTION A LANCER3") 
	}

	// ecriture de la signature
	_, err5 := f.WriteString("GOSHIELD")
    if err5 != nil {
  		panic("EXCEPTION A LANCER4") 
	}

	// ecriture du salt 
	CreateHash(doc)
	_, err0 := f.Write(doc.Salt)
	if err0 != nil {
  		panic("EXCEPTION A LANCER44") 
	}

	// ecriture de la valeur d'initialisation IV
	IV := CreateIV()
	_, err6 := f.Write(IV)
    if err6 != nil {
  		panic("EXCEPTION A LANCER5") 
	}

	// ecriture de la taille du dernier bloc (sans padding) en octets
	var length int = 0
	if (int)(stat.Size())<aes.BlockSize {
		length=(int)(stat.Size())
	}else {
		length=(int)(stat.Size())%aes.BlockSize
	}
	fmt.Printf(">>%d\n", length)
	_, err7 := f.WriteString(fmt.Sprintf("%d", length))
	if err7 != nil {
  		panic("EXCEPTION A LANCER6") 
	}

	// chiffrement de chaque bloc de données et ecriture des donnees chiffrees
	//d2 := []byte{115, 111, 109, 101, 10}
	input := make([]byte, 16)
	var seek int64 = 0
	var cipherBlock []byte
	for i:= 0; i<iterations; i++{
		input =[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

		// lecture de chaque bloc de 16 octets
		_, err8 := file.Read(input)
		if err8 != nil {
  			panic("EXCEPTION A LANCER7") 
		}
    	fmt.Println(input)

    	// si on est au tour i (i!=0), IV vaut le chiffré du tour i-1
    	if i != 0 {
    		IV = cipherBlock
    	}

    	seek = aes.BlockSize*((int64)(i+1))
   		_, err9 := f.Seek(seek, 0)
   		if err9 != nil {
  			panic("EXCEPTION A LANCER8") 
		}

		// chiffrement de chaque bloc et écriture
		cipherBlock = EncryptBlocAES(IV, doc.Hash, input)
		_, err10 := f.Write(cipherBlock)
		if err10 != nil {
  			panic("EXCEPTION A LANCER9") 
		}
	}
	
    f.Close()

}