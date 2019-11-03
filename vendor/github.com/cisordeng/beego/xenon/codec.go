package xenon

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

func EncodeMD5(unencrypted string) string {
	encrypted := fmt.Sprintf("%x", md5.Sum([]byte(unencrypted)))
	return encrypted
}


func EncodeAesWithCommonKey(unencrypted string, commonKey string) (string, error) {
	block, err := aes.NewCipher([]byte(commonKey))
	if err != nil {
		return "", err
	}
	cipherText := make([]byte, aes.BlockSize+len(unencrypted))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(cipherText[aes.BlockSize:],
		[]byte(unencrypted))
	return hex.EncodeToString(cipherText), nil

}
func DecodeAesWithCommonKey(encrypted string, commonKey string) (string, error) {
	cipherText, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(commonKey))
	if err != nil {
		return "", err
	}
	if len(cipherText) < aes.BlockSize {
		return "", errors.New("cipherText too short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	cipher.NewCFBDecrypter(block, iv).XORKeyStream(cipherText, cipherText)
	return string(cipherText), nil
}
