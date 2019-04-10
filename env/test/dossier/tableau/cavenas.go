

package main

import (
    "fmt"
    "math/rand"
    "time"
)


type Graphe struct {

    Nombre_sommet int

    Matrice_adjacent [][]int

}

func Init_graphe(n int) Graphe {

    var g Graphe

    g.Nombre_sommet=n

    g.Matrice_adjacent=make([][]int,n)

    for i:=0 ; i< n ; i++ {

        g.Matrice_adjacent[i]=make([]int,n)

    }

    return g

}
func Affichage(g Graphe){

    fmt.Println("nombre de sommet:",g.Nombre_sommet)

    for i:= 0 ; i<g.Nombre_sommet ; i++ {

        fmt.Print("Voisin de ",i,":")

        for j:=0 ; j<g.Nombre_sommet ; j++ {

            if g.Matrice_adjacent[i][j]==1 {

                fmt.Print(" ",j)

            }
        }

        fmt.Println()

    }

}

func Degre (g Graphe,n int) int {

    var compteur  int 

    for i := 0 ; i<g.Nombre_sommet; i++ {

        if g.Matrice_adjacent[n][i]==1 {

            compteur++

        }
        
    }

    return compteur

}

func Alea (n int) Graphe {   

    rand.Seed(time.Now().UTC().UnixNano())

    g:=Init_graphe(n)

    for n := 0; n<g.Nombre_sommet; n++{

        g.Matrice_adjacent[rand.Intn(g.Nombre_sommet)][rand.Intn(g.Nombre_sommet)]=1

    }

    return g
    
} 

func Cycle_graphe (n int) Graphe {

    g:=Init_graphe(n)

    for i:=0 ;i<n ; i++ {

        g.Matrice_adjacent[i][(i+1)%n]=1

        g.Matrice_adjacent[(i+1)%n][i]=1

    } 

    return g
}

func Complet_graphe (n int) Graphe {

    g:=Init_graphe(n)

    for i:=0 ;i<n ; i++ {

        for j:=0 ;j<n ; j++ {

            if i!=j {

                g.Matrice_adjacent[i][j]=1

            }

        }
    } 

    return g
}

func Cycle_graphe (n int) Graphe {

    g:=Init_graphe(n)

    for i:=0 ;i<n ; i++ {

        g.Matrice_adjacent[i][(i+1)%n]=1

        g.Matrice_adjacent[(i+1)%n][i]=1

    } 

    return g
}

func main() {
    var g Graphe

    //var degre int

    //degre=4

    g=Complet_graphe(10)

    Affichage(g)




}
