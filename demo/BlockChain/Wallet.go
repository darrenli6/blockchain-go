package BlockChain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/iocn-io/ripemd160"
	"log"
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
	// 3生成校验和数据
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
