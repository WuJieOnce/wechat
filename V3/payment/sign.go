package payment

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

type Signer struct{}

func (*Signer) GenerateSign(data map[string]string, apiKey string) string {
	// 按照 key 的字典序排序
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 拼接字符串
	signStr := ""
	for _, k := range keys {
		if data[k] != "" {
			signStr += fmt.Sprintf("%s=%s&", k, data[k])
		}
	}
	signStr += "key=" + apiKey

	// MD5 加密并转为大写
	hash := md5.Sum([]byte(signStr))
	return strings.ToUpper(hex.EncodeToString(hash[:]))
}

func (*Signer) VerifySign(data map[string]string, signStr string, apiKey string) bool {
	return false
}
