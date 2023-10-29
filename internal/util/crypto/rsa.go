package crypto

import (
	"crypto/rsa"
	"crypto/x509"
	_ "embed"
	"encoding/pem"
	"errors"
)

//go:embed private_key.pem
var PrivateKeyPEM []byte

//go:embed public_key.pem
var PublicKeyPEM []byte

// GetPrivateKey 从PEM编码的数据中解析并返回*rsa.PrivateKey
func GetRsaPrivateKey() *rsa.PrivateKey {
	block, _ := pem.Decode(PrivateKeyPEM)
	if block == nil {
		panic(errors.New("failed to parse PEM block containing the private key"))
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	return priv
}

// GetPublicKey 从PEM编码的数据中解析并返回*rsa.PublicKey
func GetRsaPublicKey() *rsa.PublicKey {
	block, _ := pem.Decode(PublicKeyPEM)
	if block == nil {
		panic(errors.New("failed to parse PEM block containing the public key"))
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		pubInterface, err = x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			panic(err)
		}
	}

	pub, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		panic(errors.New("not an RSA public key"))
	}
	return pub
}
