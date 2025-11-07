package passgen

import (
	"crypto/rand"
	"errors"
	"math/big"
)

type Config struct {
	Length int
	UseUpper bool
	UseLower bool
	UseNumbers bool
	UseSymbols bool
}

func Generate(cfg Config)(string, error) {
	if cfg.Length <= 0 {
		return "", errors.New("length must be greater than 0")
	}

	var charset string
	if cfg.UseUpper {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if cfg.UseLower {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}
	if cfg.UseNumbers {
		charset += "0123456789"
	}
	if cfg.UseSymbols {
		charset += "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	}
	if charset == "" {
		return "", errors.New("at least one character set must be selected")
	}

	pass := make([]byte, cfg.Length)
	for i := range pass {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		pass[i] = charset[n.Int64()]
	}
	return string(pass), nil
}
