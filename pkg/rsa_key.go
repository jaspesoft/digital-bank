package pkg

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

func PKCPrivateKey(privateKeyStr string) (any, error) {
	// Decode base64 string
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyStr)
	if err != nil {
		fmt.Println("Error decoding base64 string:", err)
		return nil, err
	}

	// Parse bytes to a PEM block
	privateKeyBlock, _ := pem.Decode(privateKeyBytes)
	if privateKeyBlock == nil {
		fmt.Println("Error decoding PEM block")
		return nil, errors.New("error decoding PEM block")
	}

	return x509.ParsePKCS8PrivateKey(privateKeyBlock.Bytes)
}

func RSAPrivateKey(privateKeyStr string) (*rsa.PrivateKey, error) {
	// Decode base64 string
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyStr)
	if err != nil {
		fmt.Println("Error decoding base64 string:", err)
		return nil, err
	}

	// Parse bytes to a PEM block
	privateKeyBlock, _ := pem.Decode(privateKeyBytes)
	if privateKeyBlock == nil {
		fmt.Println("Error decoding PEM block")
		return nil, errors.New("error decoding PEM block")
	}

	// Parse PEM block to a raw PKCS1 private key
	return x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
}

func RSAPublicKey(publicKey string) (*rsa.PublicKey, error) {
	// Decode base64 string
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		fmt.Println("Error decoding base64 string:", err)
		return nil, err
	}

	// Parse bytes to a PEM block
	publicKeyBlock, _ := pem.Decode(publicKeyBytes)
	if publicKeyBlock == nil {
		fmt.Println("Error decoding PEM block")
		return nil, errors.New("Error decoding PEM block")
	}

	// Parse PEM block to a raw PKIX public key
	publicKeyRaw, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		fmt.Println("Error parsing PKIX public key:", err)
		return nil, err
	}

	// Assert raw PKIX public key to *rsa.PublicKey
	rsaPublicKey, ok := publicKeyRaw.(*rsa.PublicKey)
	if !ok {
		fmt.Println("Error asserting to *rsa.PublicKey")
		return nil, errors.New("Error asserting to *rsa.PublicKey")
	}

	return rsaPublicKey, nil
}
