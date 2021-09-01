package api

import (
	"crypto/md5"
	"encoding/hex"
)

func GetHashFromPassword(password string) string {
	data := []byte(password)
	md := md5.Sum(data)
	//fmt.Printf("Password = %s\t\tHash = %s", password, hex.EncodeToString(md[:]))
	return hex.EncodeToString(md[:])
}
