package BLC

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"

	"github.com/iocn-io/ripemd160"
)

// 钱包相关

const version = byte(0x00)

// checksum 长度

const addressChecksumLen = 4

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

	// fmt.Printf("private key:%v \n  ", privateKey)
	// fmt.Printf("public key:%v \n  ", publicKey)
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

// 对公钥进行双hash sha256 -> ripemd160
func Ripemd160Hash(pubKey []byte) []byte {
	//1 sha256
	hash256 := sha256.New()
	hash256.Write(pubKey)
	hash := hash256.Sum(nil)

	// 2 ripemd160
	rmd160 := ripemd160.New()
	rmd160.Write(hash)

	return rmd160.Sum(nil)

}

// 通过公钥生成地址
func (w *Wallet) GetAddress() []byte {

	// 1 获取160哈希
	ripemd160Hash := Ripemd160Hash(w.PublicKey)
	// 2 生成version 并且加入hash中
	version_ripemd160hash := append([]byte{version}, ripemd160Hash...)
	// 3生成校验和数据   为了以后判断地址有效性
	checkSumBytes := CheckSum(version_ripemd160hash)
	// 4 拼接校验和
	bytes := append(version_ripemd160hash, checkSumBytes...)
	// 5生成base58的编码了
	base58 := Base58Encode(bytes)

	return base58

}

// 生成校验和数据
func CheckSum(payload []byte) []byte {

	hash1 := sha256.Sum256(payload)
	hash2 := sha256.Sum256(hash1[:])

	return hash2[:addressChecksumLen] // 取4个字节
}

// 判断地址有效性
func IsValidForAddress(address []byte) bool {
	// 1. 地址通过base58解码
	version_pubkey_checksumBytes := Base58Decode(address)

	// 2 拆分 进行校验和的校验
	checkSumBytes := version_pubkey_checksumBytes[len(version_pubkey_checksumBytes)-addressChecksumLen:]

	version_ripemd160 := version_pubkey_checksumBytes[:len(version_pubkey_checksumBytes)-addressChecksumLen]

	checkBytes := CheckSum(version_ripemd160)

	if bytes.Compare(checkBytes, checkSumBytes) == 0 {
		return true
	}

	return false
}
