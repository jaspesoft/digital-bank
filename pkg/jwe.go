package pkg

import (
	"fmt"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwe"
)

func EncryptData(payload, publicKey string) (string, error) {

	// Assert raw PKIX public key to *rsa.PublicKey
	rsaPublicKey, err := RSAPublicKey(publicKey)
	if err != nil {
		return "", nil
	}

	encrypted, err := jwe.Encrypt([]byte(payload), jwe.WithKey(jwa.RSA_OAEP, rsaPublicKey))
	if err != nil {
		return "", err

	}
	// Devolver el resultado encriptado como una cadena
	return string(encrypted), nil
}

func DecryptData(encryptedData string, privateKeyStr string) (string, error) {
	// Parse PEM block to a raw PKCS1 private key
	privateKeyRaw, err := RSAPrivateKey(privateKeyStr)
	if err != nil {
		fmt.Println("Error parsing PKCS1 private key:", err)
		return "", nil
	}

	// Decrypt the data
	decrypted, err := jwe.Decrypt([]byte(encryptedData), jwe.WithKey(jwa.RSA_OAEP, privateKeyRaw))
	if err != nil {
		return "", err
	}

	// Return the decrypted data as a string
	return string(decrypted), nil
}
