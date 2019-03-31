
package crypto

import "testing"
import "encoding/hex"
//import "fmt"


func TestXor(t *testing.T) {
	a := []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	b := []byte{254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254}
	c := Xor (a, b)
	test := []byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}

	if testEgaliteSlice(test, c) == false {
    	t.Errorf("Assertion TestXor de encrypt_test FAILED.")
    }

}

func TestEncryptBlocAES(t *testing.T){
	// IV sur 16 octets (128 bits).
	iv := []byte{170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170}

	// Clé sur 256 bits (AES256).
	key, _ := hex.DecodeString("6368616e6765207468697320706173736368616e676520746869732070617373")

	// Input sur 16 octets (128 bits).
	input := []byte{84, 69, 83, 84, 84, 69, 83, 84, 84, 69, 83, 84, 84, 69, 83, 84} 

	output := EncryptBlocAES(iv, key, input)

	test := []byte{126, 119, 20, 94, 251, 169, 63, 50, 62, 9, 220, 143, 72, 168, 19, 24}
	if testEgaliteSlice(test, output) == false {
    	t.Errorf("Assertion TestEncryptBlocAES de encrypt_test FAILED.")
    }
	
}