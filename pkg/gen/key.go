package gen

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

type KeyGen struct {
	Path string
	Size int
	Type string
}

func (k *KeyGen) Gen() error {
	p, err := rsa.GenerateKey(rand.Reader, k.Size)
	if err != nil {
		return err
	}

	file, err := os.Create(k.Path)
	defer file.Close()
	if err != nil {
		return err
	}

	if err := pem.Encode(file, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(p),
	}); err != nil {
		return fmt.Errorf("failed to generate key '%s': %s", k.Path, err)
	}

	return nil
}
