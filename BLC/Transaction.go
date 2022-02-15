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
func NewSimpleTransaction(from string, to string, amount int) *Transaction {

	var txInputs []*TxInput // 输入

	var txOutputs []*TxOutput // 输出

	// 消费
	txInput := &TxInput{[]byte("40956b11157f060b80a227452a89eaef6f4f2a2f7b75e9ffa833dce56d454a19"), 0, from}

	txInputs = append(txInputs, txInput)

	// 转账
	txOutput := &TxOutput{int64(amount), to}

	txOutputs = append(txOutputs, txOutput)
	// 找零
	txOutput = &TxOutput{10 - int64(amount), from}

	txOutputs = append(txOutputs, txOutput)

	// 生成交易
	tx := &Transaction{nil, txInputs, txOutputs}
	tx.HashTransaction()

	return tx
}

// 判断指定交易是否是一个coinbase交易
func (tx *Transaction) IsCoinbaseTransaction() bool {
	return len(tx.Vins[0].TxHash) == 0 && tx.Vins[0].Vout == -1

}
