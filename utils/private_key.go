package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// ParseRSAPrivateKey parses a private key from a PEM-encoded string
func ParseRSAPrivateKey(content string) (*rsa.PrivateKey, error) {
	// Decode the PEM block from the private key text
	block, _ := pem.Decode([]byte(content))
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	// Parse the private key
	var privateKey *rsa.PrivateKey
	var err error
	if block.Type == "RSA PRIVATE KEY" {
		// Parse PKCS#1 format
		privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse PKCS#1 private key: %w", err)
		}
	} else if block.Type == "PRIVATE KEY" {
		var key any
		var ok bool
		// Parse PKCS#8 format
		if key, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
			return nil, fmt.Errorf("failed to parse PKCS#8 private key: %w", err)
		}
		// Type assert to *rsa.PrivateKey
		if privateKey, ok = key.(*rsa.PrivateKey); !ok {
			return nil, fmt.Errorf("not an RSA private key")
		}
	} else {
		return nil, fmt.Errorf("unsupported key type: %s", block.Type)
	}

	return privateKey, nil
}

func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	// Read the private key file using os.ReadFile
	pemData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %w", err)
	}
	// Decode the PEM block
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}
	// Parse the private key
	var privateKey *rsa.PrivateKey
	if block.Type == "RSA PRIVATE KEY" {
		// Parse PKCS#1 format
		privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse PKCS#1 private key: %w", err)
		}
	} else if block.Type == "PRIVATE KEY" {
		var ok bool
		var key any
		// Parse PKCS#8 format
		if key, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
			return nil, fmt.Errorf("failed to parse PKCS#8 private key: %w", err)
		}
		// Type assert to *rsa.PrivateKey
		if privateKey, ok = key.(*rsa.PrivateKey); !ok {
			return nil, fmt.Errorf("not an RSA private key")
		}
	} else {
		return nil, fmt.Errorf("unsupported key type: %s", block.Type)
	}

	return privateKey, nil
}
