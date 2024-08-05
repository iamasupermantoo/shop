package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// MarshalPKCS1PrivateKey 生成 Rsa PKCS1 私钥方法
func MarshalPKCS1PrivateKey() ([]byte, []byte) {
	rng := rand.Reader
	priKey, err := rsa.GenerateKey(rng, 2048)
	if err != nil {
		panic(err)
	}

	derStream := x509.MarshalPKCS1PrivateKey(priKey)
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}

	prvKey := pem.EncodeToMemory(block)
	puKey := &priKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(puKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubKey := pem.EncodeToMemory(block)
	return prvKey, pubKey
}

// ParsePKCS1PrivateKey 私钥解析 Rsa PKCS1
func ParsePKCS1PrivateKey(priKey string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(priKey))
	if block == nil {
		return nil, errors.New("PrivateKey Error")
	}

	//	解析私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// ParsePKIXPublicKey 公钥解析 PKIX
func ParsePKIXPublicKey(pubKey string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubKey))
	if block == nil {
		return nil, errors.New("PublicKey Error")
	}

	//	解析公钥
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey := publicKeyInterface.(*rsa.PublicKey)
	return publicKey, nil
}

// ParsePrivatePublicPKCS1 解析私钥解析公钥
func ParsePrivatePublicPKCS1(priKey, pubKey string) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	//	解析私钥
	privateKey, err := ParsePKCS1PrivateKey(priKey)
	if err != nil {
		return nil, nil, err
	}

	//	解析公钥
	publicKey, err := ParsePKIXPublicKey(pubKey)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, publicKey, nil
}
