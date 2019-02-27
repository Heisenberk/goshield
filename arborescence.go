package main

import "io/ioutil"
import "fmt"
import "strings"

func cut(path string) string {

   


    path=path[:strings.LastIndexAny(path, "/")]


    return path
}
func lister(path string){
        entries, err := ioutil.ReadDir(path)

    if err != nil {

        fmt.Println(err)
    }
    for i, entry := range entries {

    fmt.Println(entry.Name())    // Nom du fichier ("myphoto.jpg")
    fmt.Println(entry.Size())    // Taille en octet (/1024 = Ko)
    fmt.Println(entry.Mode())    // Droits d'écritures "-rw-rw-rw-"
    fmt.Println(entry.ModTime()) // Date de dernière modification
    fmt.Println(entry.IsDir())   // "false" par défaut (car on ne liste pas des "directories" / répertoires)
    fmt.Println(i)
}
}
func lister_arbo(path string){

    //path := "/home/user/Bureau/app/"
    path=strings.TrimRight(path,"/")
    fmt.Println(cut(path))
    fmt.Println(strings.HasSuffix(path, "user"))
    for(strings.HasSuffix(path, "user")!=true){
    lister(path)
    path=cut(path)
    }
    fmt.Println(path)
}
func main() {

    /*
    path := "/home/user/Bureau/app/"
     path=strings.TrimRight(path,"/")
    fmt.Println(cut(path))
    //fmt.Println(strings.LastIndexAny(cut(path), "/"))
    //path=path[:strings.LastIndexAny(cut(path), "/")]
    //path=path[:strings.LastIndexAny(cut(path), "/")]
    //path=path[:strings.LastIndexAny(cut(path), "/")]
    fmt.Println(strings.HasSuffix(path, "user"))
    for(strings.HasSuffix(path, "user")!=true){
    lister(path)
    path=cut(path)
    }
    fmt.Println(path)
    */
    path := "/home/user/Bureau/app/"
    lister_arbo(path)


}