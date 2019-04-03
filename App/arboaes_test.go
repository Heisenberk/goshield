package App

import "testing"
import "strings"
import "os"


func TestURLPathSeparator(t *testing.T){

	path := "/home/user/Bureau/R/"
	Separator :=strings.LastIndexAny(path, "/")
	if(Separator!=len(path)-1){
	t.Errorf("PathSeparator manquant à la fin du chemin absolu du fichier")
}
}

func TestURLvoid(t *testing.T){

	path := "/home/user/Bureau/R/"

	if(len(path)==0){
	t.Errorf("Path vide")
}
}

func TestURLcaracteristique(t *testing.T) {

		path := "/home/user/Bureau/R/"

		PS :=0
		Word :=0
	    for i := 0; i < len(path); i++ {
        if(path[i]=='/'){
        	PS=PS+1

        }
        if ((path[i]!='/') && (path[i-1]=='/')) {
        		Word=Word+1
        }

    }

	if(Word != PS-1){
		t.Errorf("problème path caractéristique")
	}
}
func TestOpenFileOrDir(t *testing.T){
	path :="/home/user/Bureau/R"
	_, err := os.Open(path)

	path =strings.Trim(path,"/")
    path=path[strings.LastIndexAny(path, "/")+1:]

	if(err!=nil){
	t.Errorf("Impossible d'ouvrir le fichier ou dossier" + path )
	}
}




