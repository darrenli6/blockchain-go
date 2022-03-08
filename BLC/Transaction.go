package BLC

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"time"
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
	tm := time.Now().Unix() // 添加时间标识,没有时间标识 所有coinbase交易完全一样
	txHashBytes := bytes.Join([][]byte{result.Bytes(), IntToHex(tm)}, []byte{})

	hash := sha256.Sum256(txHashBytes)

	tx.TxHash = hash[:]

}

// 生成coinbase 交易 挖矿的交易 没有输入 由系统输入
func NewCoinbaseTransaction(address string) *Transaction {
	//输入

	txInput := &TxInput{[]byte{}, -1, nil, nil}

	// 输出
	txOutput := NewTxOutput(10, address)
	// hash
	txCoinbase := &Transaction{nil, []*TxInput{txInput}, []*TxOutput{txOutput}}
	txCoinbase.HashTransaction()

	return txCoinbase

}

// 生成转账交易
func NewSimpleTransaction(from string, to string, amount int, blockchain *BlockChain, txs []*Transaction) *Transaction {

	var txInputs []*TxInput // 输入

	var txOutputs []*TxOutput // 输出

	// 查找指定地址的可用UTXO
	money, spendableUXTODic := blockchain.FindSpendableUTXO(from, int64(amount), txs)

	// fmt.Printf("money : %v \n", money)

	// 获取钱包集合
	wallets, _ := NewWallets()
	wallet := wallets.Wallets[from] // 指定地址得到钱包结构

	for txHash, indexArray := range spendableUXTODic {
		txHashBytes, _ := hex.DecodeString(txHash)
		for _, index := range indexArray {
			//此处的输出是需要被消息的，必然会被其他交易的输入所引用

			// 从地址 找到wallet  然后可以找到公钥

			txInput := &TxInput{txHashBytes, index, nil, wallet.PublicKey}
			txInputs = append(txInputs, txInput)
		}
	}

	// 挖矿
	// 消费

	NewTxOutput(int64(amount), to)
	// 转账
	txOutput := NewTxOutput(int64(amount), to)

	txOutputs = append(txOutputs, txOutput)
	// 找零
	txOutput = NewTxOutput(money-int64(amount), from)

	txOutputs = append(txOutputs, txOutput)

	// 生成交易
	tx := &Transaction{nil, txInputs, txOutputs}
	tx.HashTransaction()

	//对交易进行签名
	// 主要参数tx, wallet.PrivateKey
	blockchain.SignTransaction(tx, wallet.PrivateKey)

	return tx
}

// 判断指定交易是否是一个coinbase交易
func (tx *Transaction) IsCoinbaseTransaction() bool {
	return len(tx.Vins[0].TxHash) == 0 && tx.Vins[0].Vout == -1

}

//交易签名
func (tx *Transaction) Sign(privKey ecdsa.PrivateKey, prevTxs map[string]Transaction) {
	//判断是否是挖矿交易
	if tx.IsCoinbaseTransaction() {
		return
	}

	for _, vin := range tx.Vins {
		if prevTxs[hex.EncodeToString(vin.TxHash)].TxHash == nil {
			log.Panicf("ERROR: prev transaction is not correct \n")
		}

	}
	// 提取需要签名的属性

	txCopy := tx.TrimmedCopy()
	for vin_id, vin := range txCopy.Vins {
		prevTx := prevTxs[hex.EncodeToString(vin.TxHash)] // 获取关联交易
		txCopy.Vins[vin_id].Signature = nil
		txCopy.Vins[vin_id].PublicKey = prevTx.Vouts[vin.Vout].Ripemd160Hash
		//
	}

	//获取copy tx
	txCopy = tx.TrimmedCopy()
	for vin_id, vin := range txCopy.Vins {
		prevTx := prevTxs[hex.EncodeToString(vin.TxHash)] //获取关联交易
		txCopy.Vins[vin_id].Signature = nil
		txCopy.Vins[vin_id].PublicKey = prevTx.Vouts[vin.Vout].Ripemd160Hash
		txCopy.TxHash = tx.TxHash
		txCopy.Vins[vin_id].PublicKey = nil

		r, s, err := ecdsa.Sign(rand.Reader, &privKey, txCopy.TxHash)

		if err != nil {
			log.Panicf("sign to tx %x failed ! %v \n", tx.TxHash, err)
		}
		// fmt.Printf("signature: %x , %x\n", r, s)
		// ECDSA的签名算法就是一对数字
		//SIG=(R,S)
		signature := append(r.Bytes(), s.Bytes()...)
		// fmt.Printf("signature2: %x \n", signature)
		rr := big.Int{}
		sr := big.Int{}
		sigLen := len(signature)

		rr.SetBytes(signature[:(sigLen / 2)])
		sr.SetBytes(signature[(sigLen / 2):])
		// fmt.Printf("signature[:(sigLen / 2)] signature[(sigLen / 2):]    : %x , %x\n", signature[:(sigLen/2)], signature[(sigLen/2):])
		// fmt.Printf("signature得到: %x , %x\n", rr, sr)

		tx.Vins[vin_id].Signature = signature
	}

}

// 设置用于签名的数据hash
func (tx *Transaction) Hash() []byte {
	txCopy := tx
	txCopy.TxHash = []byte{}
	hash := sha256.Sum256(txCopy.Serialize())
	return hash[:]
}

// 序列化
func (tx *Transaction) Serialize() []byte {

	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)         //新建一个encoder对象
	if err := encoder.Encode(tx); err != nil { // 编码
		log.Printf("serialize the tx to byte failed %v\n", err)
	}
	return result.Bytes()

}

// 添加一个交易的拷贝用于交易签名 返回需要签名的交易
func (tx *Transaction) TrimmedCopy() Transaction {

	var inputs []*TxInput
	var outputs []*TxOutput

	for _, vin := range tx.Vins {
		inputs = append(inputs, &TxInput{vin.TxHash, vin.Vout, nil, nil})
	}

	for _, vout := range tx.Vouts {
		outputs = append(outputs, &TxOutput{vout.Value, vout.Ripemd160Hash})
	}

	txCopy := Transaction{tx.TxHash, inputs, outputs}
	return txCopy
}

// 交易验证
func (tx *Transaction) Verify(prevTxs map[string]Transaction) bool {

	if tx.IsCoinbaseTransaction() {
		return true
	}

	// 检查能否找到交易
	// 查找每一个可以引用的交易是否包含在prevTxs
	if len(tx.Vins) > 0 {
		for _, vin := range tx.Vins {
			if prevTxs[hex.EncodeToString(vin.TxHash)].TxHash == nil {
				log.Panic("error,tx is incorrect!")
				return false
			}
		}
	} else {
		fmt.Println("tx.Vinds len is 0 ")
	}

	txCopy := tx.TrimmedCopy()

	// 获取秘钥对 ,使用相同的椭圆
	curve := elliptic.P256()

	for vinId, vin := range tx.Vins {
		prevTx := prevTxs[hex.EncodeToString(vin.TxHash)]
		txCopy.Vins[vinId].Signature = nil
		txCopy.Vins[vinId].PublicKey = prevTx.Vouts[vin.Vout].Ripemd160Hash
		// 上面生成hash的数据,所以这与签名的数据完全一致
		txCopy.TxHash = txCopy.Hash()
		txCopy.Vins[vinId].PublicKey = nil

		// 私钥ID
		// 获取r s(r和s长度相等,根据椭圆加密计算的结果)
		r := big.Int{}
		s := big.Int{}
		//sigLen := len(vin.Signature)

		//r.SetBytes(vin.Signature[:(sigLen / 2)])
		//s.SetBytes(vin.Signature[(sigLen / 2):])
		//fmt.Printf("signature得到: %x , %x\n", r, s)
		// 生成x,y  首先 签名就是数字对,公钥是x,y坐标这组合
		// 在生成公钥的时候,需要将x,y坐标组合在一起,再验证的时候,需要
		// 公钥的xy拆开
		x := big.Int{}
		y := big.Int{}

		pubkeyLen := len(vin.PublicKey)
		x.SetBytes(vin.PublicKey[:(pubkeyLen / 2)])
		y.SetBytes(vin.PublicKey[(pubkeyLen)/2:])

		//fmt.Printf("vin.PublicKey[:(pubkeyLen / 2)]  %x \n", vin.PublicKey[:(pubkeyLen/2)])
		//fmt.Printf("vin.PublicKey[(pubkeyLen)/2:]  %x \n", vin.PublicKey[(pubkeyLen)/2:])

		// 生成验证签名所需要的公钥
		rawPublicKey := ecdsa.PublicKey{curve, &x, &y}

		// 验证签名
		if !ecdsa.Verify(&rawPublicKey, txCopy.TxHash, &r, &s) {
			return true
		}
	}

	return true
}
