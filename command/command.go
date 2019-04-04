// Package command contenant les fonctions de parsing pour interpréter les commandes de l'utilisateur.
package command

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/md5"
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "io"
    "io/ioutil"
    "strings"
    "os"
    "crypto/sha1"
    //"syscall"

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
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func chemin(path string,name string,isdir bool, d  *structure.Documents){
        if(isdir ){
        path =path+name+"/"

        Lister(path,d)
        //fmt.Println(path)
        path=strings.TrimRight(path,"/")
        path=strings.TrimRight(path,name)
        //fmt.Println(path)
    }
}
func ecrire(path string,name string,isdir bool,Mode os.FileMode){
    fmt.Println(int(Mode),name)

      if(!isdir && (Mode==493 || Mode==420) ){
            data, _ := ioutil.ReadFile(path+name)
       // ciphertext := encrypt([]byte("Hello World"), "password")

   d2 := []byte(encrypt([]byte(data), "password"))
 
    err1 := ioutil.WriteFile(path+name, d2, 0644)
    check(err1)
    fmt.Println(path+name)
	}

    }
 
func createHash(key string) string {
    hasher := md5.New()
    hasher.Write([]byte(key))
    return hex.EncodeToString(hasher.Sum(nil))
}
func SHA1(data []byte) []byte {

    h := sha1.New()

    h.Write(data)

    return h.Sum(nil)

}
func encrypt(data []byte, passphrase string) []byte {
    block, _ := aes.NewCipher([]byte(createHash(passphrase)))

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        panic(err.Error())
    }
    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        panic(err.Error())
    }
    ciphertext := gcm.Seal(nonce, nonce, data, nil)
    return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
    key := []byte(createHash(passphrase))
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err.Error())
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        panic(err.Error())
    }
    nonceSize := gcm.NonceSize()
    nonce, ciphertext := data[:nonceSize], data[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        panic(err.Error())
    }
    return plaintext
}
func getname_file(path string) string {

   path =strings.Trim(path,"/")


    path=path[strings.LastIndexAny(path, "/")+1:]



    return path
}
   func ecrire_dech(path string,name string,isdir bool,Mode os.FileMode){
    fmt.Println(int(Mode),name)

      if(!isdir && (Mode==493 || Mode==420) ){
            data, _ := ioutil.ReadFile(path+name)
       // ciphertext := encrypt([]byte("Hello World"), "password")

   d2 := []byte(decrypt([]byte(data), "password"))
 
    err1 := ioutil.WriteFile(path+name, d2, 0777)
    check(err1)
    fmt.Println(path+name)

    }else if(Mode!=493){
            if (Mode!=420 ){
            
        
        fmt.Print("ce fichier")
        fmt.Print(name)
        fmt.Println("ne possède pas les droits en lectures /écritures")
        }
    }
}
func ecrire_encryption(path string,name string,isdir bool,Mode os.FileMode){
    fmt.Println(int(Mode),name)

      if(!isdir && (Mode==493 || Mode==420) ){
               
      
            data, _ := ioutil.ReadFile(path+name)

       // ciphertext := encrypt([]byte("Hello World"), "password")

   d2 := []byte(encrypt([]byte(data), "password"))
 
    err1 := ioutil.WriteFile(path+name, d2, 0644)
    check(err1)
    fmt.Println(path+name)

    }else if(Mode!=493){
            if (Mode!=420 ){
            
        
        fmt.Print("ce fichier")
        fmt.Print(name)
        fmt.Println("ne possède pas les droits en lectures /écritures")
        }
    }
}
func ecrire_encryption_file(name string,Mode os.FileMode){

    fmt.Println(int(Mode),name)

      if((Mode==493 || Mode==420) ){
               
      
            data, _ := ioutil.ReadFile(name)

       // ciphertext := encrypt([]byte("Hello World"), "password")

   d2 := []byte(encrypt([]byte(data), "password"))
 
    err1 := ioutil.WriteFile(name, d2, 0644)
    check(err1)
    fmt.Println(name)

    }else if(Mode!=493){
            if (Mode!=420 ){
            
        
        fmt.Print("ce fichier")
        fmt.Print(name)
        fmt.Println("ne possède pas les droits en lectures /écritures")
        }
    }
}	

func Lister(path string,d *structure.Documents){
    
    fi, err := os.Stat(path)
    
    if err != nil {
        fmt.Println(err)
        return
    }
    mode := fi.Mode();
    //fmt.Println(mode.IsRegular())
    if(mode.IsDir()==true){
    	if(strings.LastIndexAny(path, "/") != len(path) - 1){
        path=path+ string(os.PathSeparator)
    }
    
   entries, err := ioutil.ReadDir(path)
         //path=strings.TrimRight(path,"/")

    if err != nil {

        fmt.Println(err)
    }
    for _, entry := range entries {

 
   //ecrire_encryption(path,entry.Name(),entry.IsDir(),entry.Mode()) 
    	p:=path + entry.Name()

    	if(p[len(p)-4:]!=".gsh"){
    	crypto.EncryptFileAES(path+entry.Name(),d)
    	    fmt.Println(path+entry.Name())
    	}
chemin(path,entry.Name(),entry.IsDir(),d)

}
}else if(mode.IsRegular()==true){

crypto.EncryptFileAES(getname_file(path),d)
  //ecrire_encryption_file(getname_file(path),mode) 

    }



       
}
func lister_dech(path string,d *structure.Documents){
        entries, err := ioutil.ReadDir(path)
         //path=strings.TrimRight(path,"/")

    if err != nil {

        fmt.Println(err)
    }
    for _, entry := range entries {

        ecrire_dech(path,entry.Name(),entry.IsDir(),entry.Mode())
    chemin(path,entry.Name(),entry.IsDir(),d)
   
}


}

func give_me_the_name_of(path string){
	        entries,_ := ioutil.ReadDir(path)
    for _, entry := range entries {
       

    fmt.Printf("%d \n",entry.Mode())    // Droits d'écritures "-rw-rw-rw-"
 

    
    
    if(entry.IsDir()){
        path =path+entry.Name()+"/"

        give_me_the_name_of(path)
        //fmt.Println(path)
        path=strings.TrimRight(path,"/")
        path=strings.TrimRight(path,entry.Name())
        //fmt.Println(path)
    }

}
}	

func Interpret( d  *structure.Documents ,err error ) {


	if (err==nil){
		if(d.Mode == 1){
			for i := 0; i < len(d.Doc); i++ {
				Lister(d.Doc[i],d)
			}
			
		}
		if(d.Mode == 2){

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