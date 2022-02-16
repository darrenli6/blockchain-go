package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
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
func NewSimpleTransaction(from string, to string, amount int, blockchain *BlockChain) *Transaction {

	var txInputs []*TxInput // 输入

	var txOutputs []*TxOutput // 输出

	// 查找指定地址的可用UTXO
	money, spendableUXTODic := blockchain.FindSpendableUTXO(from, int64(amount))

	fmt.Printf("money : %v \n", money)

	for txHash, indexArray := range spendableUXTODic {
		txHashBytes, _ := hex.DecodeString(txHash)
		for _, index := range indexArray {
			//此处的输出是需要被消息的，必然会被其他交易的输入所引用
			txInput := &TxInput{txHashBytes, index, from}
			txInputs = append(txInputs, txInput)
		}
	}

	// 消费

	// 转账
	txOutput := &TxOutput{int64(amount), to}

	txOutputs = append(txOutputs, txOutput)
	// 找零
	txOutput = &TxOutput{money - int64(amount), from}

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
