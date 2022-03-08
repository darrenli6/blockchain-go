package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	msg := "hello, world"
	// 生成hash
	hash := sha256.Sum256([]byte(msg))

	//签名
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("signature: %x , %x\n", r, s)

	valid := ecdsa.Verify(&privateKey.PublicKey, hash[:], r, s)
	fmt.Println("signature verified:", valid)
}
