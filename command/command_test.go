// Package command contenant les fonctions de parsing pour interpréter les commandes de l'utilisateur.
package command

import "testing"
//import "io/ioutil"
import "github.com/Heisenberk/goshield/structure"

// TestSansArgument teste la commande "./goshield"
func TestSansArgument(t *testing.T) {
    arg := []string{"./goshield"}
    _, err := Parse(arg[1:])
    if err == nil {
       t.Errorf("Assertion TestSansArgument de command_test FAILED.")
    }
}

// TestChiffrementSeul teste la commande "./goshield -c"
func TestChiffrementSeul(t *testing.T) {
    arg := []string{"./goshield", "-c"}
    _, err := Parse(arg[1:])
    if err == nil {
       t.Errorf("Assertion TestChiffrementSeul de command_test FAILED.")
    }
}

// TestDechiffrementSeul teste la commande "./goshield -d"
func TestDechiffrementSeul(t *testing.T) {
    arg := []string{"./goshield", "-d"}
    _, err := Parse(arg[1:])
    if err == nil {
       t.Errorf("Assertion TestChiffrementSeul de command_test FAILED.")
    }
}

// TestPasswordArg teste la commande "./goshield -d -p"
func TestPasswordArg(t *testing.T) {
    arg := []string{"./goshield", "-d", "-p"}
    _, err := Parse(arg[1:])
    if err == nil {
       t.Errorf("Assertion TestPasswordArg de command_test FAILED.")
    }
}

// TestPassword teste la commande "./goshield -d -p password"
func TestPassword(t *testing.T) {
    arg := []string{"./goshield", "-d", "-p", "password"}
    _, err := Parse(arg[1:])
    if err == nil {
       t.Errorf("Assertion TestPassword de command_test FAILED.")
    }
}

// TestChiffrementFichierDossier teste la commande "./goshield -e -p password file1.txt file2.txt project/"
func TestChiffrementFichierDossier(t *testing.T) {

    arg := []string{"./goshield", "-e", "-p", "password", "file1.txt", "file2.txt", "project/"}
    goshield, err := Parse(arg[1:])

    if err != nil {
       t.Errorf("Assertion 1 TestChiffrementFichierDossier de command_test FAILED.")
    }

    if goshield.Mode != structure.ENCRYPT {
        t.Errorf("Assertion 2 TestChiffrementFichierDossier de command_test FAILED.")
    }

    if goshield.Password != "password" {
        t.Errorf("Assertion 3 TestChiffrementFichierDossier de command_test FAILED.")
    }

    if len(goshield.Doc) != 3 {
        t.Errorf("Assertion 4 TestChiffrementFichierDossier de command_test FAILED.")
    }

    if goshield.Doc[0]!="file1.txt" {
        t.Errorf("Assertion 5 TestChiffrementFichierDossier de command_test FAILED.")
    }

    if goshield.Doc[1]!="file2.txt" {
        t.Errorf("Assertion 6 TestChiffrementFichierDossier de command_test FAILED.")
    }

    if goshield.Doc[2]!="project/" {
        t.Errorf("Assertion 7 TestChiffrementFichierDossier de command_test FAILED.")
    }
}

// TestDechiffrementFichierDossier teste la commande "./goshield -d -p password file1.txt file2.txt project/"
func TestDechiffrementFichierDossier(t *testing.T) {

    arg := []string{"./goshield", "-d", "-p", "password", "file1.txt", "file2.txt", "project/"}
    goshield, err := Parse(arg[1:])

    if err != nil {
       t.Errorf("Assertion 1 TestDechiffrementFichierDossier de command_test FAILED.")
    }

    if goshield.Mode != structure.DECRYPT {
        t.Errorf("Assertion 2 TestDechiffrementFichierDossier de command_test FAILED.")
    }

    if goshield.Password != "password" {
        t.Errorf("Assertion 3 TestDechiffrementFichierDossier de command_test FAILED.")
    }

    if len(goshield.Doc) != 3 {
        t.Errorf("Assertion 4 TestDechiffrementFichierDossier de command_test FAILED.")
    }

    if goshield.Doc[0]!="file1.txt" {
        t.Errorf("Assertion 5 TestDechiffrementFichierDossier de command_test FAILED.")
    }

    if goshield.Doc[1]!="file2.txt" {
        t.Errorf("Assertion 6 TestDechiffrementFichierDossier de command_test FAILED.")
    }

    if goshield.Doc[2]!="project/" {
        t.Errorf("Assertion 7 TestDechiffrementFichierDossier de command_test FAILED.")
    }
}
