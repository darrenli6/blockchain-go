package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 目标难度值 代表生成的哈希前targetBit位为0，才能满足条件
const targetBit = 20

// 工作量证明
type ProofOfWork struct {
	Block  *Block
	target *big.Int
}

// 创建新的pow对象
func NewProofOfWork(block *Block) *ProofOfWork {

	target := big.NewInt(1)

	target.Lsh(target, 256-targetBit)
	//8
	// 前2为都为0
	// 左移
	// 8-2
	// 1 << 6
	// 0010 0000

	return &ProofOfWork{block, target}

}

//开始工作量证明
func (proofOfWork *ProofOfWork) Run() ([]byte, int64) {
	var nonce = 0 //碰撞次数

	var hash [32]byte

	var hashInt big.Int

	for {
		//1 数据拼接
		dataBytes := proofOfWork.prepareData(nonce)
		hash = sha256.Sum256(dataBytes)
		hashInt.SetBytes(hash[:])
		fmt.Printf("hash: \r %x   ", hash)

		//难度比较
		if proofOfWork.target.Cmp(&hashInt) == 1 {
			break
		}

		nonce++

	}

	return hash[:], int64(nonce)

}

// 准备数据 ，将区块属性链接起来，返回一个字节数组
func (pow *ProofOfWork) prepareData(nonce int) []byte {

	data := bytes.Join([][]byte{
		pow.Block.PreBlockHash,
		pow.Block.Data,
		IntToHex(pow.Block.TimeStamp),
		IntToHex(pow.Block.Height),
		IntToHex(int64(nonce)),
		IntToHex(targetBit),
	}, []byte{})

	return data
}
