package ioread

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	IO "io/ioutil"
	"os"

	"encoding/json"
)

var mail_account = "mail_account.txt"
var passToDecriptFile = "passtodecrpty"

type ContentFileEncrpt struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

/*
** To Create Crtpy-Text that store Gmail Account
 */

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

func encryptFile(filename string, data []byte, passphrase string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(encrypt(data, passphrase))
}

func decryptFile(filename string, passphrase string) []byte {
	data, _ := IO.ReadFile(filename)
	return decrypt(data, passphrase)
}

/*Call When U want to gen new EmailID&Pass*/
func EncryptByAdmin() {
	// fmt.Println("Starting the application...")
	// ciphertext := encrypt([]byte("Hello World"), "password")
	// fmt.Printf("Encrypted: %x\n", ciphertext)
	// plaintext := decrypt(ciphertext, "password")
	// fmt.Printf("Decrypted: %s\n", plaintext)

	// contentPaint := []byte(`{"id":"dentiplansw@gmail.com","password":"pass"}`)
	contentPaint := []byte(`{"id":"myemail@gmail.com","password":"pass"}`)
	encryptFile(mail_account, []byte(contentPaint), passToDecriptFile)
}

func DecrptEmailId() (string, string) {
	contentDecrypt := string(decryptFile(mail_account, passToDecriptFile))
	//fmt.Println("contentDecrypt : " + contentDecrypt)
	contentDecryptStruct := ContentFileEncrpt{}
	err := json.Unmarshal([]byte(contentDecrypt), &contentDecryptStruct)

	if err != nil {
		fmt.Println("something error")
	}

	return contentDecryptStruct.ID, contentDecryptStruct.Password
}
