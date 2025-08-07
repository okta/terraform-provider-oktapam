package pam

import (
	"fmt"

	"github.com/go-jose/go-jose/v4"
)

type JWKEncryptor struct {
	PublicKeys []jose.JSONWebKey
}

// Encrypt encrypts the data with the public key and returns the encrypted value, default content encryption algorithm is jose.A256GCM
func (jwke *JWKEncryptor) Encrypt(data []byte, contentAlgo jose.ContentEncryption) (string, error) {
	var encrypter jose.Encrypter
	var err error

	if len(jwke.PublicKeys) == 0 {
		return "", fmt.Errorf("can't find the public key from the JWKS")
	}
	// we will use the prefered key algorithm by the priority from the JWKS
	key, err := getPreferredKey(jwke.PublicKeys)
	if err != nil {
		return "", err
	}

	recipient := jose.Recipient{
		Algorithm: jose.KeyAlgorithm(key.Algorithm),
		Key:       key,
	}
	encrypter, err = jose.NewEncrypter(contentAlgo, recipient, &jose.EncrypterOptions{})
	if err != nil {
		return "", err
	}

	jwe, err := encrypter.Encrypt(data)
	if err != nil {
		return "", err
	}
	return jwe.FullSerialize(), nil
}

// add more key algorithms and their rank here
var keyAlgorithmRank = map[string]int{
	"ECDH-ES+A256KW": 1, // 1st choice
	"ECDH-ES+A192KW": 2, // 2nd choice
	"ECDH-ES+A128KW": 3, // 3rd choice
	"ECDH-ES":        4, // etc..
	"RSA-OAEP-256":   5,
}

// getPreferredKey returns the preferred key from the slice of keys based on the rank
func getPreferredKey(keys []jose.JSONWebKey) (*jose.JSONWebKey, error) {
	var selectedKey *jose.JSONWebKey
	var selectedRank int

	for _, key := range keys {
		if rank, ok := keyAlgorithmRank[key.Algorithm]; ok && (selectedKey == nil || rank < selectedRank) {
			selectedKey = &key
			selectedRank = rank
		}
	}
	if selectedKey == nil {
		return nil, fmt.Errorf("can't find the preferred key from the JWKS")
	}
	return selectedKey, nil
}
