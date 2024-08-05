package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func AESEncrypt(plaintext, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	origData := PKCS7Padding([]byte(plaintext), blockSize)
	blockMode := NewECBEncrypt(block)
	encrypted := make([]byte, len(origData))
	blockMode.CryptBlocks(encrypted, origData)

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func AESDecrypt(ciphertext, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	encrypted, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	blockMode := NewECBDecrypt(block)
	decrypted := make([]byte, len(encrypted))
	blockMode.CryptBlocks(decrypted, encrypted)

	decrypted = PKCS7UnPadding(decrypted)
	return string(decrypted), nil
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypt ecb

// NewECBEncrypt creates a new ECB encrypter
func NewECBEncrypt(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypt)(newECB(b))
}

func (x *ecbEncrypt) BlockSize() int { return x.blockSize }

func (x *ecbEncrypt) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypt ecb

// NewECBDecrypt creates a new ECB decrypt
func NewECBDecrypt(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypt)(newECB(b))
}

func (x *ecbDecrypt) BlockSize() int { return x.blockSize }

func (x *ecbDecrypt) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func PKCS7UnPadding(src []byte) []byte {
	length := len(src)
	unPadding := int(src[length-1])
	return src[:(length - unPadding)]
}
