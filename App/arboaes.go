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
    "syscall"
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
    func ecrire_dech(path string,name string,isdir bool,Mode os.FileMode){
    fmt.Println(int(Mode),name)

      if(!isdir && (Mode==syscall.O_RDWR) ){
            data, _ := ioutil.ReadFile(path+name)
       // ciphertext := encrypt([]byte("Hello World"), "password")

   d2 := []byte(decrypt([]byte(data), "password"))
 
    err1 := ioutil.WriteFile(path+name, d2, 0777)
    check(err1)
    fmt.Println(path+name)

    }else if(Mode!=syscall.O_RDWR){
    	fmt.Print("ce fichier")
    	fmt.Print(name)
    	fmt.Println("ne possède pas les droits en lectures /écritures")
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
func cut(path string) string {

   


    path=path[:strings.LastIndexAny(path, "/")+1]


    return path
}
func ecrire_encryption(path string,name string,isdir bool,Mode os.FileMode){
    fmt.Println(int(Mode),name)

      if(!isdir && (Mode==syscall.O_RDWR ) ){
               
      
            data, _ := ioutil.ReadFile(path+name)

       // ciphertext := encrypt([]byte("Hello World"), "password")

   d2 := []byte(encrypt([]byte(data), "password"))
 
    err1 := ioutil.WriteFile(path+name, d2, 0644)
    check(err1)
    fmt.Println(path+name)

    }else if(Mode!=syscall.O_RDWR){
    	fmt.Print("ce fichier")
    	fmt.Print(name)
    	fmt.Println(" ne possède pas les droits en lectures/écritures")
    	}	
}	

func lister(path string){

        entries, err := ioutil.ReadDir(path)
         //path=strings.TrimRight(path,"/")

    if err != nil {

        fmt.Println(err)
    }
    for i, entry := range entries {

        /*
    fmt.Println(entry.Name())    // Nom du fichier ("myphoto.jpg")
    fmt.Println(entry.Size())    // Taille en octet (/1024 = Ko)
    fmt.Println(entry.Mode())    // Droits d'écritures "-rw-rw-rw-"
    fmt.Println(entry.ModTime()) // Date de dernière modification
    fmt.Println(entry.IsDir())   // "false" par défaut (car on ne liste pas des "directories" / répertoires)
    */
    fmt.Println(i)
    //fmt.Println(path+entry.Name())
/*
    if(!entry.IsDir() ){
        //ciphertext := encrypt([]byte("Hello World"), "password")
        data, _ := ioutil.ReadFile(path+entry.Name())
    d1 := []byte(encrypt([]byte(data), "password"))
    //d2 := []byte( decrypt(ciphertext, "password"))
    err := ioutil.WriteFile(path+entry.Name(), d1, 0644)
    check(err)
    //err1 := ioutil.WriteFile(path+entry.Name(), d2, 0644)
    //check(err1)
    fmt.Println(path+entry.Name())
    }
    */
   ecrire_encryption(path,entry.Name(),entry.IsDir(),entry.Mode()) 
chemin(path,entry.Name(),entry.IsDir())

}
}
func lister_dech(path string){
        entries, err := ioutil.ReadDir(path)
         //path=strings.TrimRight(path,"/")

    if err != nil {

        fmt.Println(err)
    }
    for _, entry := range entries {
        /*
    fmt.Println(entry.Name())    // Nom du fichier ("myphoto.jpg")
    fmt.Println(entry.Size())    // Taille en octet (/1024 = Ko)
    fmt.Println(entry.Mode())    // Droits d'écritures "-rw-rw-rw-"
    fmt.Println(entry.ModTime()) // Date de dernière modification
    fmt.Println(entry.IsDir())   // "false" par défaut (car on ne liste pas des "directories" / répertoires)
    */
    //fmt.Println(i)
    //fmt.Println(path+entry.Name())
/*
    if(!entry.IsDir() && int(entry.Mode)==290 ){
            

        data, _ := ioutil.ReadFile(path+entry.Name())
       // ciphertext := encrypt([]byte("Hello World"), "password")

   d2 := []byte(decrypt([]byte(data), "password"))
 
    err1 := ioutil.WriteFile(path+entry.Name(), d2, 0644)
    check(err1)
    fmt.Println(path+entry.Name())
    }
    
    
    if(entry.IsDir()){
        path =path+entry.Name()+"/"

        lister(path)
        //fmt.Println(path)
        path=strings.TrimRight(path,"/")
        path=strings.TrimRight(path,entry.Name())
        //fmt.Println(path)
    }
    */
        ecrire_dech(path,entry.Name(),entry.IsDir(),entry.Mode())
    chemin(path,entry.Name(),entry.IsDir())
}


}

func give_me_the_name_of(path string){
	        entries,_ := ioutil.ReadDir(path)
    for _, entry := range entries {
       
    //fmt.Println(entry.Name())    // Nom du fichier ("myphoto.jpg")
    //fmt.Println(entry.Size())    // Taille en octet (/1024 = Ko)
    fmt.Printf("%d \n",entry.Mode())    // Droits d'écritures "-rw-rw-rw-"
    //fmt.Println(entry.ModTime()) // Date de dernière modification
    //fmt.Println(entry.IsDir())   // "false" par défaut (car on ne liste pas des "directories" / répertoires)
   
    //fmt.Println(i)
    //fmt.Println(path+entry.Name())


    
    
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


    path := "/home/user/Bureau/go/Go-Language/"

    //give_me_the_name_of(path)
    //lister_dech(path)
    lister(path)




}