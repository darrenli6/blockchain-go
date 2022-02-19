package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/iocn-io/ripemd160"
)

func main() {
	//sha256

	hasher := sha256.New()
	hasher.Write([]byte("the"))
	bytes := hasher.Sum(nil)
	fmt.Printf("%x\n", bytes)

	//ripemd160

	hash160 := ripemd160.New()
	hash160.Write([]byte("themonnstone"))
	byteRipemd := hash160.Sum(nil)
	fmt.Printf("%x\n", byteRipemd)

}
