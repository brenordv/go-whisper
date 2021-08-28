package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
)

type Session struct {
	Mine Key
	Partner Key
}

type Key struct {
	Public  *rsa.PublicKey
	Private *rsa.PrivateKey
}

func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	return privateKey, &privateKey.PublicKey, err
}

func (k Key) PublicKeyToBytes() ([]byte, error) {
	pubASN1, err := x509.MarshalPKIXPublicKey(k.Public)
	if err != nil {
		return nil, err
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return pubBytes, nil
}

func (k *Key) BytesToPublicKey(pub []byte) error {
	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block") //todo: remove
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return err
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		return err
	}

	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		return errors.New("could not cast parsed to pubkey")
	}
	k.Public = key
	return nil
}


func (k Key) Encrypt(msg []byte) ([]byte, error) {
	cypher, err := rsa.EncryptPKCS1v15(rand.Reader, k.Public, msg)
	return cypher, err
}


func (k Key) Decrypt(ciphertext []byte) ([]byte, error) {
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, k.Private, ciphertext)
	return plaintext, err
}