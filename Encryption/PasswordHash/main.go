package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	pass := os.Args[1]

	encrypt(pass)
}

func encrypt(pass string) {
	hash := md5.New()
	hash.Write([]byte(pass))

	hashThePass := hex.EncodeToString(hash.Sum(nil))

	fmt.Println(hashThePass)
}
