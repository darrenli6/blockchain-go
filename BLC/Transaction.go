package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// 存放交易相关的
type Transaction struct {
	//tx hash  交易唯一的表示
	TxHash []byte
	// 输入
	Vins []*TxInput
	// 输出
	Vouts []*TxOutput
}

// 生成交易hash

func (tx *Transaction) HashTransaction() {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(tx)

	if err != nil {
		log.Panicf("tx hash generate failed %v \n ", err)
	}
	hash := sha256.Sum256(result.Bytes())
	tx.TxHash = hash[:]

}

// 生成coinbase 交易 挖矿的交易 没有输入 由系统输入
func NewCoinbaseTransaction(address string) *Transaction {
	//输入
	txInput := &TxInput{[]byte{}, -1, "Genesis Data"}
	// 输出
	txOutput := &TxOutput{10, address}

	// hash
	txCoinbase := &Transaction{nil, []*TxInput{txInput}, []*TxOutput{txOutput}}
	txCoinbase.HashTransaction()

	return txCoinbase

}

// 生成转账交易
