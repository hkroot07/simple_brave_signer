package cryptoutils

import (
	"fmt"

	"golang.org/x/crypto/argon2"
)

type KeyDerivationConfig struct {
	Passphrase []byte
	Salt       []byte
}

func DeriveKey(config KeyDerivationConfig) ([]byte, error) {
	if len(config.Passphrase) == 0 || len(config.Salt) == 0 {
		return nil, fmt.Errorf("passphrase and salt  cannot be empty")
	}

	return argon2.IDKey(config.Passphrase, config.Salt, 1, 64*1024, 4, 32), nil
}
