package gen

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

type KeyGen struct {
	path    string
	size    int
	keyType string
	prefix  string
}

func NewKeyGen(f *Flags) *KeyGen {
	return &KeyGen{
		path:    f.Path,
		size:    f.Size,
		keyType: f.Type,
		prefix:  f.Prefix,
	}
}

func (k *KeyGen) Gen() error {
	path, err := homedir.Expand(k.path)
	if err != nil {
		return err
	}
	k.path, err = filepath.Abs(path)
	if err != nil {
		return err
	}

	dir, fl := filepath.Split(k.path)
	k.path = filepath.Join(dir, k.prefix+fl)

	p, err := rsa.GenerateKey(rand.Reader, k.size)
	if err != nil {
		return err
	}

	file, err := os.Create(k.path)
	defer file.Close()
	if err != nil {
		return err
	}

	if err := pem.Encode(file, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(p),
	}); err != nil {
		return fmt.Errorf("failed to generate key '%s': %s", k.path, err)
	}

	return nil
}
