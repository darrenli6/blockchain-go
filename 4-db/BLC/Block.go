package BLC

import (
	"time"
	"bytes"
	"crypto/sha256"
)

// 实现一个区块结构

type Block struct {

	TimeStamp     int64  // 区块时间戳，区块产生的时间
	Height        int64  // 区块的高度
	PreBlockHash  []byte
	Hash          []byte  //当前区块的哈希
	Data          []byte  // 交易数据
    Nonce         int64

}


// 创建新的区块
func NewBlock(height int64,prevBlockHash []byte,Data []byte) *Block{
	  var block Block
	  block =Block{Height:height,PreBlockHash:prevBlockHash,Data:Data,TimeStamp:time.Now().Unix()}
	  //block.SetHash() // 生成当前的hash

	  pow:=NewProofOfWork(&block)
	  hash,nonce :=pow.Run()
	  block.Hash=hash
	  block.Nonce=nonce

	  return &block
}


// 计算区块hash
func (b *Block) SetHash(){
	// int64 转化为字节数组
	 heightBytes :=IntToHex(b.Height)

     timeStampBytes :=IntToHex(b.TimeStamp)

     // 拼接所有的属性 进行哈希
     blockBytes :=bytes.Join([][]byte{heightBytes,timeStampBytes,b.PreBlockHash,b.Data},[]byte{})

     hash :=sha256.Sum256(blockBytes)

     b.Hash=hash[:]


}
// 生成一个创世区块
func CreateGenesisBlock(data string) *Block{
	return NewBlock(1,nil,[]byte(data))
}
