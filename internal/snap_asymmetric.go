package internal

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

// A Signer can create signatures that verify against a public key.
type Signer interface {
	// Sign returns a raw signature for the given data. This method
	// will apply the hash specified for the key type to the data.
	Sign(data []byte) ([]byte, error)
}

type rsaPrivateKey struct {
	*rsa.PrivateKey
}

func (r *rsaPrivateKey) Sign(data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	d := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, r.PrivateKey, crypto.SHA256, d)
}

func GenerateAsymmetricSignature(data string) (string, error) {
	privateKey, err := base64.StdEncoding.DecodeString(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return "", fmt.Errorf("failed to load private key: %s", err)
	}
	signer, err := parsePrivateKey([]byte(string(privateKey)))
	if err != nil {
		return "", err
	}
	signed, err := signer.Sign([]byte(data))
	if err != nil {
		return "", err
	}

	sig := base64.StdEncoding.EncodeToString(signed)
	return sig, nil
}

// this function is used for parse private key to create new signer
func parsePrivateKey(pemBytes []byte) (Signer, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("ssh: no key found")
	}

	var rawkey interface{}

	switch block.Type {
	case "RSA PRIVATE KEY":
		rsa, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rawkey = rsa
	case "PRIVATE KEY":
		rsa, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rawkey = rsa
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %q", block.Type)
	}

	var sshKey Signer
	switch t := rawkey.(type) {
	case *rsa.PrivateKey:
		sshKey = &rsaPrivateKey{t}
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %T", rawkey)
	}

	return sshKey, nil
}
