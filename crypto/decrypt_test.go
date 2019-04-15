// Package crypto contenant les fonctions de chiffrement/déchiffrement.
package crypto

import "testing"
import "encoding/hex"
import "os"
import "io/ioutil"
import "bytes"
import "sync"

import "github.com/Heisenberk/goshield/structure"

// Test de chiffrement sur un bloc. 
func TestDecryptBlocAES(t *testing.T){

    // IV sur 16 octets (128 bits).
	iv := []byte{170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170}

	// Clé sur 256 bits (AES256).
	key, _ := hex.DecodeString("6368616e6765207468697320706173736368616e676520746869732070617373")

	// Input sur 16 octets (128 bits).
	input := []byte{126, 119, 20, 94, 251, 169, 63, 50, 62, 9, 220, 143, 72, 168, 19, 24}

	output, err := DecryptBlocAES(iv, key, input)

	if err != nil {
    	t.Errorf("Assertion 1 TestEncryptBlocAES de encrypt_test FAILED.")
    }

	test := []byte{84, 69, 83, 84, 84, 69, 83, 84, 84, 69, 83, 84, 84, 69, 83, 84} 
	if testEgaliteSlice(test, output) == false {
    	t.Errorf("Assertion 2 TestDecryptBlocAES de decrypt_test FAILED.")
    }
	
}

// Test de chiffrement suivi de déchiffrement sur le fichier env/test/test6.md
func TestEncryptDecryptFile(t * testing.T){

	wg := &sync.WaitGroup{}
	wg.Add(1)
	channel := make (chan error)

	var d structure.Documents
	d.Password = "password"
	go EncryptFileAES("../env/test/test6.md", &d, channel, wg)
	err1 := <- channel 
	if err1 != nil {
		t.Errorf("Erreur 1 TestEncryptDecryptFile de decrypt_test.")
	}
	wg.Wait()

	err2 := os.Rename("../env/test/test6.md.gsh", "../env/test6.md.gsh")
	if err2 != nil {
		t.Errorf("Erreur 2 TestEncryptDecryptFile de decrypt_test.")
	}

	wg2 := &sync.WaitGroup{}
	wg2.Add(1)
	channel2 := make (chan error)

	var d2 structure.Documents
	d2.Password = "password"
	go DecryptFileAES("../env/test6.md.gsh", &d2, channel2, wg2)
	err3 := <- channel2
	if err3 != nil {
		t.Errorf("Erreur 3 TestEncryptDecryptFile de decrypt_test.")
	}
	wg2.Wait()

    file1, err1 := ioutil.ReadFile("../env/test/test6.md")
    if err1 != nil {
        t.Errorf("Erreur 4 TestEncryptDecryptFile de decrypt_test.")
    }
    file2, err2 := ioutil.ReadFile("../env/test6.md")
    if err2 != nil {
        t.Errorf("Erreur 5 TestEncryptDecryptFile de decrypt_test.")
    }

    donnees1 := []byte(string(file1))
    donnees2 := []byte(string(file2))

    if bytes.Equal(donnees1, donnees2) == false {
    	t.Errorf("Assertion TestEncryptDecryptFile de decrypt_test FAILED.")
    }

    os.Remove("../env/test6.md")
    os.Remove("../env/test6.md.gsh")
}