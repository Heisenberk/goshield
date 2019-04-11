// Package crypto contenant les fonctions de chiffrement/déchiffrement.
package crypto

import "testing"
import "crypto/sha256"

import "github.com/Heisenberk/goshield/structure"

// testEgaliteSlice renvoie true si les deux slices sont égaux et false sinon. 
func testEgaliteSlice(a, b []byte) bool {
	if len(a)!= len(b){
		return false
	}
	for i:=0 ; i< len(a) ;i++ {
		if a[i]!= b[i] {
			return false
		}
	}
	return true
}

// Test la bonne création d'un hash. 
func TestCreateHash(t *testing.T) {
	salt := []byte{84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84}
	doc := structure.Documents{Mode: structure.ENCRYPT, Password: "password", Salt: salt}

    if doc.Mode != structure.ENCRYPT {
       t.Errorf("Assertion 1 TestCreateHash de hash_test FAILED.")
    }

    if testEgaliteSlice(salt, doc.Salt) == false {
    	t.Errorf("Assertion 2 TestCreateHash de hash_test FAILED.")
    }

    hash := sha256.New()
    concat := append(doc.Salt, doc.Password...) 
    hash.Write(concat)
    doc.Hash = hash.Sum(nil)

    test := []byte{185, 200, 122, 186, 38, 7, 228, 168, 206, 176, 175, 82, 249, 83, 99, 235, 150, 140, 146, 208, 201, 49, 30, 157, 144, 0, 163, 88, 191, 153, 103, 9}

    if testEgaliteSlice(test, doc.Hash) == false {
    	t.Errorf("Assertion 3 TestCreateHash de hash_test FAILED.")
    }

}