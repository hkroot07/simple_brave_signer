package crypto_utils

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type KeyDerivationConfig struct {
	Passphrase []byte
	Salt       []byte
}

func MakeCrypter(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher block: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM cipher: %v", err)
	}
	return gcm, nil
}

func DeriveKey(config KeyDerivationConfig) ([]byte, error) {
	if len(config.Passphrase) == 0 || len(config.Salt) == 0 {
		return nil, fmt.Errorf("passphrase and salt  cannot be empty")
	}

	return argon2.IDKey(config.Passphrase, config.Salt, 1, 64*1024, 4, 32), nil
}
