package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

// 实现一个区块结构

type Block struct {
	TimeStamp    int64 // 区块时间戳，区块产生的时间
	Height       int64 // 区块的高度
	PreBlockHash []byte
	Hash         []byte //当前区块的哈希
	//Data         []byte        // 交易数据
	Txs   []*Transaction // 交易
	Nonce int64
}

// 创建新的区块
func NewBlock(height int64, prevBlockHash []byte, txs []*Transaction) *Block {
	fmt.Println("NewBlock ...")
	var block Block
	block = Block{Height: height, PreBlockHash: prevBlockHash, Txs: txs, TimeStamp: time.Now().Unix()}
	//block.SetHash() // 生成当前的hash

	pow := NewProofOfWork(&block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return &block
}

// 计算区块hash
//func (b *Block) SetHash() {
//	// int64 转化为字节数组
//	heightBytes := IntToHex(b.Height)
//
//	timeStampBytes := IntToHex(b.TimeStamp)
//
//	// 拼接所有的属性 进行哈希
//	blockBytes := bytes.Join([][]byte{heightBytes, timeStampBytes, b.PreBlockHash, b.Txs}, []byte{})
//
//	hash := sha256.Sum256(blockBytes)
//
//	b.Hash = hash[:]
//
//}

// 生成一个创世区块
func CreateGenesisBlock(txs []*Transaction) *Block {
	return NewBlock(1, nil, txs)
}

// 序列化 将区块结构序列化为[]byte (字节数组)
func (block *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)            //新建一个encoder对象
	if err := encoder.Encode(block); err != nil { // 编码
		log.Printf("serialize the block to byte failed %v\n", err)
	}
	return result.Bytes()

}

// 反序列化 ，将字节数组结构化为区块
func DecerializeBlock(blockBytes []byte) *Block {

	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))

	if err := decoder.Decode(&block); err != nil {
		log.Panicf("Decerialize the to []byte to block failed ! %v \n", err)
	}
	return &block
}

// 把区块中所有的区块结构转化为[]byte
func (block *Block) HashTransactions() []byte {
	var txHashes [][]byte
	for _, tx := range block.Txs {
		txHashes = append(txHashes, tx.TxHash)
	}

	txHash := sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}
