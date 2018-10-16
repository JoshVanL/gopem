package gen

import (
	"crypto/rsa"
	//"crypto/x509"
)

type Flags struct {
	Path   string
	Type   string
	Size   int
	Prefix string
}

type CAGen struct {
	dir     string
	size    int
	keyType string
	prefix  string
	pk      *rsa.PrivateKey
}

func NewCAGen(f *Flags, pk *rsa.PrivateKey) *CAGen {
	return &CAGen{
		dir:     f.Path,
		size:    f.Size,
		keyType: f.Type,
		prefix:  f.Prefix,
		pk:      pk,
	}
}

func (c *CAGen) Gen() error {
	return nil
}
