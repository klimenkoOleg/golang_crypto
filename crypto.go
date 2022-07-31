package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	ecies "github.com/ecies/go/v2"
	"io"
	"log"
)

func main() {
	hasher := sha1.New()
	str := "asd"
	hasher.Write([]byte(str))
	//sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	sha := hex.EncodeToString(hasher.Sum(nil))

	fmt.Println("sha: " + sha)
	fmt.Printf("% x \n", hasher.Sum(nil))

	fmt.Println("<<<<<<<<<<------------->>>>>>>>>>>>>")
	key1 := "68544020247570407220244063724074"
	key2 := "54684020247570407220244063724074" // this is the correct key
	key3 := "54684020247570407220244063727440"
	///
	msg := "f28fe539655fd6f7275a09b7c3508a3f81573fc42827ce34ddf1ec8d5c2421c3"
	//data, err := hex.DecodeString()
	fmt.Println(hex.EncodeToString(sha256b(haxStr2bytes(key1))))
	fmt.Println(hex.EncodeToString(sha256b(haxStr2bytes(key2))))
	fmt.Println(hex.EncodeToString(sha256b(haxStr2bytes(key3))))
	fmt.Println("<<<<<<<<<<------------->>>>>>>>>>>>>")
	fmt.Println(msg)

	fmt.Println(decryptAes())

	elliptic()
}

func elliptic() {
	k, err := ecies.GenerateKey()
	if err != nil {
		panic(err)
	}
	log.Println("key pair has been generated")
	fmt.Println("Public key: ")
	fmt.Println(hex.EncodeToString(k.PublicKey.Bytes(false)))
	//hex.EncodeToString(k.)

	plaintext := []byte("Hello Blockchain!")

	ciphertext, err := ecies.Encrypt(k.PublicKey, []byte("Hello Blockchain!"))
	if err != nil {
		panic(err)
	}
	log.Printf("plaintext encrypted: %v\n", ciphertext)

	plaintext, err = ecies.Decrypt(k, ciphertext)
	if err != nil {
		panic(err)
	}
	log.Printf("ciphertext decrypted: %s\n", string(plaintext))
}

func haxStr2bytes(str string) []byte {
	data, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return data
}

func sha1b(str string) string {
	hasher := sha1.New()
	hasher.Write([]byte(str))
	//sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	sha := hex.EncodeToString(hasher.Sum(nil))
	return sha
}

func sha256b(data []byte) []byte {
	hash := sha256.Sum256([]byte(data))
	return hash[:]
}

func EncryptMessage(key []byte, message string) (string, error) {
	byteMsg := []byte(message)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("could not create new cipher: %v", err)
	}

	cipherText := make([]byte, aes.BlockSize+len(byteMsg))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("could not encrypt: %v", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], byteMsg)

	resultHex, err := base64.StdEncoding.EncodeToString(cipherText), nil
	fmt.Println(resultHex)
	fmt.Printf("ddd %s\n", hex.Dump(cipherText))
	return resultHex, err
}

// data []byte, key []byte, initVector []byte
func decryptAes() string {
	key := haxStr2bytes("54684020247570407220244063724074")
	aes, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encryptedMsg := haxStr2bytes("876b4e970c3516f333bcf5f16d546a87aaeea5588ead29d213557efc1903997e")
	initVector := haxStr2bytes("656e6372797074696f6e496e74566563")

	//int11 := aes.BlockSize
	//iv := encryptedMsg[:int11]
	//ciphertext = encryptedMsg[aes.BlockSize:]

	out, err := DecryptCbc(aes, encryptedMsg, initVector)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ddd %s\n", hex.Dump(out))

	return hex.EncodeToString(out)
}

//48656c6c6f20426c6f636b636861696e210f0f0f0f0f0f0f0f0f0f0f0f0f0f0f
// result: 48656c6c6f20426c6f636b636861696e21
func DecryptCbc(aes cipher.Block, encrypted []byte, iv []byte) ([]byte, error) {
	//aescipher, _ := aes.NewCipher([]byte(util.Md5sum(key)))

	decryptor := cipher.NewCBCDecrypter(aes, iv)

	//decryptedBytes := make([]byte, len(encrypted))
	//decryptor.CryptBlocks(decryptedBytes, encrypted)
	decryptor.CryptBlocks(encrypted, encrypted)

	//return encrypted, nil

	decryptedBytes, err := PKCS7Unpad2(encrypted)
	//decryptedBytes, err := pad.PKCS7Unpad(decryptedBytes, decryptor.BlockSize())
	if err != nil {
		return nil, err
	}

	//return string(decryptedBytes[:len(decryptedBytes)]), nil
	return decryptedBytes[:len(decryptedBytes)], nil
}

func DecryptMessage(key []byte, message string) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", fmt.Errorf("could not base64 decode: %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("could not create new cipher: %v", err)
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("invalid ciphertext block size")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

/*
func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if blockSize < 1 {
		return nil, fmt.Errorf("Block size looks wrong")
	}

	if len(data)%blockSize != 0 {
		return nil, fmt.Errorf("Data isn't aligned to blockSize")
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("Data is empty")
	}

	paddingLength := int(data[len(data)-1])
	for _, el := range data[len(data)-paddingLength:] {
		if el != byte(paddingLength) {
			return nil, fmt.Errorf("Padding had malformed entries. Have '%x', expected '%x'", paddingLength, el)
		}
	}

	return data[:len(data)-paddingLength], nil
}*/

// PKCS7Unpad removes PKCS7 padding from the data block, http://en.wikipedia.org/wiki/Padding_(cryptography)#PKCS7
// this function may return an error id padding is incorrect,
// however it will return unpadded data in any case
func PKCS7Unpad2(padded []byte) (message []byte, err error) {
	// read padding length
	plen := len(padded)
	last_byte := padded[plen-1]
	padlen := int(last_byte)

	// check validity of PKCS7 padding
	for i := padlen; i > 1; i-- {
		if padded[plen-i] != last_byte {
			err = fmt.Errorf("Invalid padding (byte -%d: %d). Is the message supplied PKCS7 padded?", i, padded[plen-i])
			break
		}
	}

	// remove padding
	return padded[:plen-padlen], err
}
