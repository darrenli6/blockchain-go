package BLC

import "bytes"

// 交易输出
type TxOutput struct {

	// 有多少钱 金额
	Value int64

	// // 这个钱是谁的 用户名 解锁脚本
	// ScriptPubkey string

	Ripemd160Hash []byte // 哈希值

}

// output 身份验证
func (txOutput *TxOutput) UnLockScriptPubkeyWithAddress(address string) bool {
	hash160 := Lock(address)
	return bytes.Compare(txOutput.Ripemd160Hash, hash160) == 0
}

// 锁定
func Lock(address string) []byte {
	publicKeyHash := Base58Decode([]byte(address))
	hash160 := publicKeyHash[1:(len(publicKeyHash) - addressChecksumLen)]
	return hash160
}

// 创建output对象
func NewTxOutput(value int64, address string) *TxOutput {
	txOutput := &TxOutput{}
	hash160 := Lock(address)
	txOutput.Value = value
	txOutput.Ripemd160Hash = hash160
	return txOutput
}

/*
address: [15exdKbAKQcHkXeyLFwR22XY6efqk2ERhz]
 address: [1Ao5ozQJfPavNirDSoKZHByzkn8KokwzwM]
 address: [1w9FGzeUDB577dv5WEPmDveLsGWc3jTZM]

 ./main send -from "[\"15exdKbAKQcHkXeyLFwR22XY6efqk2ERhz\"]" -to "[\"1Ao5ozQJfPavNirDSoKZHByzkn8KokwzwM\"]" -amount   
*/
