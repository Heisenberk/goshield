// Package crypto contenant les fonctions de chiffrement/déchiffrement.
package crypto

import "math/rand"
import "time"
import "crypto/aes"
import "crypto/cipher"
import "os"
import "errors"
import "fmt"
import "io/ioutil"
import "strings"
import "sync"

import "github.com/Heisenberk/goshield/structure"

// CreateIV génère une valeur initiale IV aléatoire. 
func CreateIV() []byte {
	rand.Seed(time.Now().UnixNano())
	iv := make([]byte, 16)
	rand.Read(iv)
	return iv
}

// EncryptBlocAES chiffre 1 bloc input avec la clé key et la valeur initiale iv pour donner le bloc chiffré. 
func EncryptBlocAES(iv []byte, key []byte, input []byte) ([]byte, error) {

	// Résultat du chiffrement sera dans output.
	output := make([]byte, aes.BlockSize)

	// Si la taille de l'entrée est invalide on lance une erreur. 
	if len(input)%aes.BlockSize != 0 {
		return output, errors.New("\033[31mFailure Encryption\033[0m : Taille du bloc invalide.")
	}

	// Preparation du bloc qui sera chiffré. 
	block, err := aes.NewCipher(key)
	if err != nil {
		return output, errors.New("\033[31mFailure Encryption\033[0m : Erreur lors du chiffrement d'un bloc.")
	}

	// Chiffrement AES avec le mode opératoire CBC.
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(output[:aes.BlockSize], input)

	return output, nil
}

// EncryptFileAES chiffre un fichier de chemin pathFile avec les données doc. 
func EncryptFileAES(pathFile string, doc *structure.Documents, channel chan error, wg *sync.WaitGroup) {

	// synchronisation pour les autres goroutines. 
	defer wg.Done()

	// ouverture du fichier a chiffrer
	inputFile, err1 := os.Open(pathFile) 
	if err1 != nil {
		var texteError string = "\033[31mFailure Encryption\033[0m : Impossible d'ouvrir le fichier à chiffrer "+pathFile+". "
		channel <- errors.New(texteError)
		return
	}
	stat, err2 := inputFile.Stat()
	if err2 != nil {
  		var texteError string = "\033[31mFailure Encryption\033[0m : Impossible d'interpréter le fichier à chiffrer "+pathFile+". "
		channel <- errors.New(texteError)
		return 
	}

	// vérification de la bonne permission
	if stat.Mode().String()[1]=='-' {
		var texteError string = "\033[31mFailure Encryption\033[0m : Permission du fichier à chiffrer "+pathFile+" incorrecte . "
		channel <- errors.New(texteError)
		return 
	}

	var division int = (int)(stat.Size()/aes.BlockSize)
	var iterations int = division
	if (int)(stat.Size())%aes.BlockSize != 0 {
		iterations=iterations+1
	}

	// ouverture du fichier résultat
    outputFile, err3 := os.Create(pathFile+".gsh")
    if err3 != nil {
  		var texteError string = "\033[31mFailure Encryption\033[0m : Impossible d'écrire le fichier chiffré "+pathFile+".gsh. "
		channel <- errors.New(texteError)
		return
	}

	// ecriture de la signature
	_, err4 := outputFile.WriteString("GOSHIELD")
    if err4 != nil {
  		var texteError string = "\033[31mFailure Encryption\033[0m : Impossible d'écrire dans le fichier chiffré "+pathFile+".gsh. "
		channel <- errors.New(texteError) 
		return
	}

	// ecriture du salt 
	CreateHash(doc)
	_, err5 := outputFile.Write(doc.Salt)
	if err5 != nil {
  		var texteError string = "\033[31mFailure Encryption\033[0m : Impossible de générer le salt. "
		channel <- errors.New(texteError)
		return
	}

	// ecriture de la valeur d'initialisation IV
	IV := CreateIV()
	_, err6 := outputFile.Write(IV)
    if err6 != nil {
  		var texteError string = "\033[31mFailure Encryption\033[0m : Impossible d'écrire la valeur d'initialisation IV. "
		channel <- errors.New(texteError) 
		return
	}

	// ecriture de la taille du dernier bloc (sans padding) en octets
	var length int = 0
	if (int)(stat.Size())<aes.BlockSize {
		length=(int)(stat.Size())
	}else {
		length=(int)(stat.Size())%aes.BlockSize
	}
	lengthWritten := make([]byte, 1)
	lengthWritten[0]=byte(length)
	_, err7 := outputFile.Write(lengthWritten)
	if err7 != nil {
  		var texteError string = "\033[31mFailure Encryption\033[0m : Impossible d'écrire la taille du dernier bloc chiffré. "
		channel <- errors.New(texteError)
		return
	}

	// chiffrement de chaque bloc de données et ecriture des donnees chiffrees
	input := make([]byte, 16)
	var cipherBlock []byte

	for i:= 0; i<iterations; i++{

		input =[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

		// lecture de chaque bloc de 16 octets
		_, err8 := inputFile.Read(input)
		if err8 != nil {
  			var texteError string = "\033[31mFailure Encryption\033[0m : Impossible de lire dans le fichier à chiffrer "+pathFile+". "
			channel <- errors.New(texteError)
			return
		}

    	// si on est au tour i (i!=0), IV vaut le chiffré du tour i-1
    	if i != 0 {
    		IV = cipherBlock
    	}

		// chiffrement de chaque bloc 
		var err10 error
		cipherBlock, err10 = EncryptBlocAES(IV, doc.Hash, input)
		if err10 != nil {
			var texteError string = "\033[31mFailure Encryption\033[0m : Impossible de chiffrer le fichier "+pathFile+". "
			channel <- errors.New(texteError)
			return
		}

		// écriture du bloc chiffré
		_, err11 := outputFile.Write(cipherBlock)
		if err11 != nil {
  			var texteError string = "\033[31mFailure Encryption\033[0m : Impossible d'écrire dans le fichier "+pathFile+".gsh. "
			channel <- errors.New(texteError)
			return
		}
	}

	// fermeture des fichiers. 
    outputFile.Close()
    inputFile.Close()

    var messageSuccess string = "- \033[32mSuccess Encryption\033[0m : "+pathFile+" : resultat dans le fichier "+pathFile+".gsh"
    fmt.Println(messageSuccess)
    channel <- nil
    return

}

// EncryptFolder chiffre le contenu d'un dossier de chemin path avec les données doc. 
func EncryptFolder (path string, d *structure.Documents) {

	// Permettra de synchroniser le chiffrement des fichiers contenus dans le dossier. 
	wgFolder := &sync.WaitGroup{}

    // On lit dans le dossier visée par le chemin
   entries, err := ioutil.ReadDir(path)
    if err != nil {
        fmt.Println("- \033[31mFailure Encryption\033[0m : impossible d'ouvrir "+path)
    }

    // Comptage des futures goruntines à lancer. 
	var countFiles int = 0
    for _, entry := range entries {
        p:=path + entry.Name()

        // Si l'extension du fichier est différent de .gsh on peut chiffrer le fichier
        if(p[len(p)-4:]!=".gsh"){
           
           newPath := path+entry.Name()
           fi, err := os.Stat(newPath)
            valid := true
            if err != nil {
                fmt.Println("- \033[31mFailure Encryption\033[0m : "+newPath+" n'existe pas ")
                valid = false
            }

            // Si l'élément du dossier est valide. 
            if valid == true {

                mode := fi.Mode();
                if mode.IsRegular()== true {
                	// si l'extension du fichier est différent de .gsh on peut chiffrer le fichier.
                	if newPath[len(newPath)-4:]!=".gsh"{
 						countFiles=countFiles+1
                	} 
                }
            }
        }
    }

    // Initialisation des goroutines
	wgFolder.Add(countFiles)
	channelFolder := make (chan error)

    // Pour chaque élément du dossier, on le chiffre.
    for _, entry := range entries {

        p:=path + entry.Name()

        // Si l'extension du fichier est différent de .gsh on peut chiffrer le fichier
        if(p[len(p)-4:]!=".gsh"){
           
           newPath := path+entry.Name()
           fi, err := os.Stat(newPath)
            valid := true
            if err != nil {
                fmt.Println("- \033[31mFailure Encryption\033[0m : "+newPath+" n'existe pas ")
                valid = false
            }

            // Si l'élément du dossier est valide. 
            if valid == true {

                mode := fi.Mode();

                // Si l'objet spécifié par le chemin est un dossier.
                if(mode.IsDir()==true){

                    // Si l'utilisateur a oublié le "/" à la fin du chemin du fichier.
                    if(strings.LastIndexAny(newPath, "/") != len(newPath) - 1){
                      newPath=newPath+ string(os.PathSeparator)
                    }
                    // Chiffrement du dossier. 
                    EncryptFolder(newPath, d)

                // Si l'objet spécifié par le chemin est un fichier.
                }else if mode.IsRegular()== true {
                	// si l'extension du fichier est différent de .gsh on peut chiffrer le fichier.
                	if newPath[len(newPath)-4:]!=".gsh"{
                		go EncryptFileAES(newPath, d, channelFolder, wgFolder)
                	} 
                }
            }
        }
    }

    // Récupération des codes erreurs
    for _, entry := range entries {
        p:=path + entry.Name()

        // Si l'extension du fichier est différent de .gsh on peut chiffrer le fichier
        if(p[len(p)-4:]!=".gsh"){
           
           newPath := path+entry.Name()
           fi, err := os.Stat(newPath)
            valid := true
            if err != nil {
                fmt.Println("- \033[31mFailure Encryption\033[0m : "+newPath+" n'existe pas ")
                valid = false
            }

            // Si l'élément du dossier est valide. 
            if valid == true {

                mode := fi.Mode();

                if mode.IsRegular()== true {
                	// si l'extension du fichier est différent de .gsh on peut chiffrer le fichier.
                	if newPath[len(newPath)-4:]!=".gsh"{
 						err := <- channelFolder 
						if err != nil {
							fmt.Println(err)
						}
                	} 
                }
            }
        }
    }

    // Attente que toutes les goroutines se terminent. 
    if countFiles!= 0 {
    	wgFolder.Wait()
    }
}

// DecryptFileFolder déchiffre les éléments choisis par l'utilisateur avec les données doc. 
func EncryptFileFolder(d *structure.Documents) {

	// Permettra de synchroniser le chiffrement des fichiers contenus dans le dossier. 
	wg := &sync.WaitGroup{}

	// Comptage des futures goruntines à lancer. 
	var countFiles int = 0
	for j:=0 ; j < len(d.Doc); j++ {
		stat, err := os.Stat(d.Doc[j])
		
		valid := true
        if err != nil {
            valid = false
        }

        if valid == true {
        	mode := stat.Mode()
        	if mode.IsRegular() == true {
				countFiles=countFiles+1
			}
        }
	}

	// Initialisation des goruntines
	wg.Add(countFiles)
	channel := make (chan error)

	// Pour chaque élément choisi par l'utilisateur, on le chiffre. 
    for i := 0; i < len(d.Doc); i++ {

    	// Lecture de cet élément. 
        fi, err := os.Stat(d.Doc[i])
        valid := true
        if err != nil {
            fmt.Println("- \033[31mFailure Encryption\033[0m : "+d.Doc[i]+" n'existe pas ")
            valid = false
        }

        // Si cet élément est valide. 
        if valid == true {
            mode := fi.Mode();

            // Si l'objet spécifié par le chemin est un dossier.
            if(mode.IsDir()==true){

                // Si l'utilisateur a oublié le "/" à la fin du chemin du fichier. 
                if(strings.LastIndexAny(d.Doc[i], "/") != len(d.Doc[i]) - 1){
                  d.Doc[i]=d.Doc[i]+ string(os.PathSeparator)
                  
                }

                // Chiffrement du dossier. 
                EncryptFolder(d.Doc[i], d)

            // Si l'objet spécifié par le chemin est un fichier.
            }else if mode.IsRegular()== true {

            	if d.Doc[i][len(d.Doc[i])-4:]!=".gsh"{
            		// Chiffrement du fichier. 
                	go EncryptFileAES(d.Doc[i], d, channel, wg)
            	}
            }
        }
    }

    // Récupération des erreurs lors du chiffrement. 
    for j:=0 ; j < len(d.Doc); j++ {
		stat2, err := os.Stat(d.Doc[j])
		valid := true
        if err != nil {
            valid = false
        }
        if valid == true {
        	mode := stat2.Mode()
        	if mode.IsRegular() == true {
				err := <- channel 
				if err != nil {
					fmt.Println(err)
				}
			}	
        }
	}

	// Attente que toutes les goroutines se terminent. 
    if countFiles!= 0 {
    	wg.Wait()
    }
}