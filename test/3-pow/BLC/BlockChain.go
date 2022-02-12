package BLC

type BlockChain struct {
	Blocks []*Block
}

// 初始化区块链

func CreateBlockChainWithGenesisBlock() *BlockChain {
	//添加创世区块

	genesisBlock := CreateGenesisBlock("init block")

	return &BlockChain{[]*Block{genesisBlock}}
}

// 添加新区块到区块链中
func (bc *BlockChain) AddBlock(height int64, data []byte, prevBlockHash []byte) {
	newBlock := NewBlock(height, prevBlockHash, data)
	bc.Blocks = append(bc.Blocks, newBlock)
}
