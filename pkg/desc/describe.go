package desc

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	//"os"

	"github.com/mitchellh/go-homedir"
)

type Desc struct {
	data []byte
	path string
}

//type test struct {
//	check  func([]byte) (interface{}, error)
//	printf func(interface{})
//}

func New(path string) (*Desc, error) {
	p, err := homedir.Expand(path)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}

	return &Desc{
		path: path,
		data: b,
	}, nil
}

func (d *Desc) Describe() error {
	p, _ := pem.Decode(d.data)
	if p == nil {
		return errors.New("failed to parse PEM block")
	}

	return d.printf(p.Bytes)
}

func (d *Desc) printf(pem []byte) error {
	pk, err := x509.ParsePKCS1PrivateKey(pem)
	if err == nil {
		d.rsaPrivateKey(pk)
		return nil
	}

	cert, err := x509.ParseCertificate(pem)
	if err == nil {
		d.cert(cert)
		return nil
	}

	return nil
}

func (d *Desc) rsaPrivateKey(pk *rsa.PrivateKey) {
	fmt.Printf("Private key: %s\n", d.path)
	fmt.Printf("N: %s\n", pk.N.String())
	fmt.Printf("E: %d\n", pk.E)
	fmt.Printf("D: %s\n", pk.D.String())
	fmt.Print("\n")
}

func (d *Desc) cert(cert *x509.Certificate) {
	fmt.Printf("certificate: %s\n", d.path)
	fmt.Printf("Issuer: %s\n", cert.Issuer.String())
	fmt.Printf("Valid from: %s\n", cert.NotBefore.String())
	fmt.Printf("Valid till: %s\n", cert.NotAfter.String())
	fmt.Printf("Emails: %s\n", cert.EmailAddresses)
	fmt.Printf("Subject: %s\n", cert.Subject.String())
	fmt.Print("\n")
}
