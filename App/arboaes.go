package main

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




func check(e error) {
    if e != nil {
        panic(e)
    }
}
func chemin(path string,name string,isdir bool){
        if(isdir ){
        path =path+name+"/"

        lister(path)
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

    }else if(Mode!=493 || Mode!=420){
        fmt.Print("ce fichier")
        fmt.Print(name)
        fmt.Println("ne possède pas les droits en lectures /écritures")
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

    }else if(Mode!=493 || Mode!=420){

        fmt.Print("ce fichier ")
        fmt.Print(name)
        fmt.Println(" ne possède pas les droits en lectures/écritures")
    	}	
        fmt.Println()
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

    }else if(Mode!=493 || Mode!=420){
        fmt.Println()
    	fmt.Print("ce fichier ")
    	fmt.Print(name)
    	fmt.Println(" ne possède pas les droits en lectures/écritures")
        fmt.Println()
    	}	
}	

func lister(path string){
    
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

 
   ecrire_encryption(path,entry.Name(),entry.IsDir(),entry.Mode()) 
chemin(path,entry.Name(),entry.IsDir())

}
}else if(mode.IsRegular()==true){


  ecrire_encryption_file(getname_file(path),mode) 

    }



       
}
func lister_dech(path string){
        entries, err := ioutil.ReadDir(path)
         //path=strings.TrimRight(path,"/")

    if err != nil {

        fmt.Println(err)
    }
    for _, entry := range entries {

        ecrire_dech(path,entry.Name(),entry.IsDir(),entry.Mode())
    chemin(path,entry.Name(),entry.IsDir())
   
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
func main() {


    path := "/home/user/Bureau/R"

    fmt.Println(len(path))
    

    //give_me_the_name_of(path)
    //lister_dech(path)
    lister(path)




}