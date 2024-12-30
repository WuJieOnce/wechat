package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// SignWithPrivateKey 使用商户私钥对数据进行签名（SHA256-RSA2048）
func SignWithPrivateKey(message string, privateKey *rsa.PrivateKey) (string, error) {
	// 计算 SHA256 哈希
	hash := sha256.New()
	_, err := hash.Write([]byte(message))
	if err != nil {
		return "", fmt.Errorf("failed to hash message: %v", err)
	}
	hashed := hash.Sum(nil)

	// 使用私钥签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", fmt.Errorf("failed to sign message: %v", err)
	}

	// 返回 Base64 编码的签名
	return base64.StdEncoding.EncodeToString(signature), nil
}
