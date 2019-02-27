package main
import (

	"bytes"

	"crypto/aes"

	"crypto/cipher"

	"fmt"

)



type AesEncrypt struct {

	Key string

	Iv  string

}



func NewEnc() *AesEncrypt {

	return &AesEncrypt{}

}



func (this *AesEncrypt) getKey() []byte {



	keyLen := len(this.Key)

	if keyLen < 16 {

		panic("res key 16")

	}

	arrKey := []byte(this.Key)

	if keyLen >= 32 {


		return arrKey[:32]

	}

	if keyLen >= 24 {



		return arrKey[:24]

	}



	return arrKey[:16]

}





func (this *AesEncrypt) Encrypt(strMesg string) ([]byte, error) {



	plantText := []byte(strMesg)

	fmt.Println(plantText)

	key := this.getKey()

	block, err := aes.NewCipher(key) 

	if err != nil {

		return nil, err

	}


	fmt.Println(block.BlockSize())

	plantText = this.PKCS7Padding(plantText, block.BlockSize())

	fmt.Println(len(plantText))

	blockModel := cipher.NewCBCEncrypter(block, []byte(this.Iv)[:aes.BlockSize])



	ciphertext := make([]byte, len(plantText))



	blockModel.CryptBlocks(ciphertext, plantText)

	return ciphertext, nil

}




func (this *AesEncrypt) Decrypt(src []byte) (strDesc string, err error) {



	defer func() {

		

		if e := recover(); e != nil {

			err = e.(error)

		}

	}()



	key := this.getKey()

	keyBytes := []byte(key)

	block, err := aes.NewCipher(keyBytes) 

	if err != nil {

		return "", err

	}

	blockModel := cipher.NewCBCDecrypter(block, []byte(this.Iv)[:aes.BlockSize])

	plantText := make([]byte, len(src))

	blockModel.CryptBlocks(plantText, src)

	plantText = this.PKCS7UnPadding(plantText, block.BlockSize())

	return string(plantText), nil

}





func (this *AesEncrypt) PKCS7UnPadding(plantText []byte, blockSize int) []byte {

	length := len(plantText)

	unpadding := int(plantText[length-1])

	return plantText[:(length - unpadding)]

}





func (this *AesEncrypt) PKCS7Padding(ciphertext []byte, blockSize int) []byte {

	padding := blockSize - len(ciphertext)%blockSize

	padtext := bytes.Repeat([]byte{byte(padding)}, padding)


	return append(ciphertext,padtext...)

	//return append(ciphertext,bytes.Repeat([]byte("7"), 7)...)

}
func main() {

	aesEnc := NewEnc()

	aesEnc.Iv = `sdf234wef34efrfT`

	aesEnc.Key = `aaC5p6c5L2g6KeJ5`

	source := `i want go`

	des, err := aesEnc.Encrypt(source,&aesEnc)

	fmt.Println(des)
	if err != nil {

		fmt.Println("hahaha watele")

	}

	resource, err := aesEnc.Decrypt(des)

	fmt.Println(resource)

}