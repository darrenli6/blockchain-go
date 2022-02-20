package main

import (
	"blockchain-go/4-pow/demo/BlockChain"
	"crypto/sha256"
	"fmt"
	"github.com/iocn-io/ripemd160"
)

func main1() {
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

func main() {
	// 创建wallet
	wallet := BlockChain.NewWallet()
	address := wallet.GetAddress()
	fmt.Printf("address %s \n", address)
}