package BlockChain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

// 钱包相关

//钱包结构

type Wallet struct {
	//1 私钥
	PrivateKey ecdsa.PrivateKey
	// 2 公钥
	PublicKey []byte
}

// 创建钱包
func NewWallet() *Wallet {
	privateKey, publicKey := newKeyPair()

	return &Wallet{PrivateKey: privateKey, PublicKey: publicKey}
}

// 生成公钥 私钥对

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	//椭圆加密
	// curve 椭圆
	priv, err := ecdsa.GenerateKey(curve, rand.Reader)

	if nil != err {
		log.Panicf("ecdsa generate ket fiailed ! %v \n", err)
	}

	pubkey := append(priv.PublicKey.X.Bytes(), priv.PublicKey.Y.Bytes()...)
	return *priv, pubkey
}
