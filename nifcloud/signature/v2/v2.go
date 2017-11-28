package v2

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	// "encoding/base64"

	// "github.com/heriet/funicula/nifcloud"
	"github.com/heriet/funicula/nifcloud/credential"
)

// CalcSignature calcs sigunature v2
func CalcSignature(cred *credential.Credential, method string, path string, encodedValues string, host string) string {
	payload := method + "\n" + host + "\n" + path + "\n" + encodedValues
	hash := hmac.New(sha256.New, []byte(cred.SecretAccessKey))
	hash.Write([]byte(payload))
	sig := make([]byte, base64.StdEncoding.EncodedLen(hash.Size()))
	base64.StdEncoding.Encode(sig, hash.Sum(nil))

	return string(sig)
}
