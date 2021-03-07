package common

import (
	"log"
	"time"

	"hello_19/fernet"
)

func Encrypt(key string, acc string) string {
	k, _ := fernet.DecodeKey(key)

	t, err := fernet.EncryptAndSign([]byte(acc), k)
	if err != nil {
		log.Fatalln(err)
	}

	return string(t)
}

func Decrypt(key string, enc_acc string) string {
	k := fernet.MustDecodeKeys(key)
	s := fernet.VerifyAndDecrypt([]byte(enc_acc), 60*time.Second, k)

	return string(s)
}
