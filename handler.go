package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
)

func GetHashFromPassword(password string) string {
	data := []byte(password)
	md := md5.Sum(data)
	//fmt.Printf("Password = %s\t\tHash = %s", password, hex.EncodeToString(md[:]))
	return hex.EncodeToString(md[:])
}

func NotFoundHandler(w http.ResponseWriter, message string) {
	w.WriteHeader(404)
	fmt.Fprint(w, message)
}

func JsonResponse(w http.ResponseWriter, json_ans []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(json_ans)
}

func InnerErrorHandler(w http.ResponseWriter, err error) {
	w.WriteHeader(500)
	fmt.Fprint(w, err)
}
