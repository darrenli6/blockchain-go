package BLC

import "bytes"

// 交易输入
type TxInput struct {

	// 交易hash(不是当前的交易hash,上一笔的)
	TxHash []byte
	// 引用上一笔交易的output索引
	Vout int
	// 用户名 锁定脚本  转账人
	// ScriptSig string
	// 数字签名
	Signature []byte

	// 公钥
	PublicKey []byte
}

// // 判断能不能花费，能不能引用指定地址的output
// func (in *TxInput) UnLockWithAddress(address string) bool {
// 	return in.ScriptSig == address
// }

func (in *TxInput) UnLockRipemd160Hash(ripemd160Hash []byte) bool {
	// 获取到ripemd160哈希
	inputRipedmd160 := Ripemd160Hash(in.PublicKey)

	return (bytes.Compare(inputRipedmd160, ripemd160Hash)) == 0
}
