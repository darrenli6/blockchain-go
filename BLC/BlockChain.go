package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

const dbName = "bc.db" // 存储数据的数据库文件

const blockTableName = "blocks" // 表名称

type BlockChain struct {
	DB  *bolt.DB // 数据库
	Tip []byte   // 最新区块的hash值

}

// 初始化区块链

func CreateBlockChainWithGenesisBlock() *BlockChain {

	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {

		log.Panicf("open  the db failed %v \n", err)
	}
	defer db.Close()

	var blockHash []byte // 需要存储到数据库中的区块hash
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b == nil {
			b, err = tx.CreateBucket([]byte(blockTableName))
			if err != nil {
				log.Panicf("create bucket [%s] failed %v \n", blockTableName, err)
			}

		}
		if nil != b {

			// 创建创世区块
			genesisBlock := CreateGenesisBlock("the data of genesis block")
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if nil != err {
				log.Panicf("put data of genesisBlock to db failed ! error %v \n", err)
			}
			// 存储最新区块的hash value

			err = b.Put([]byte("1"), genesisBlock.Hash)
			if err != nil {
				log.Panicf("put the hash of latest block to db failed %v \n", err)

			}

			blockHash = genesisBlock.Hash
		}

		return nil

	})

	return &BlockChain{db, blockHash}
}

// 添加新区块到区块链中
func (bc *BlockChain) AddBlock(height int64, data []byte, prevBlockHash []byte) {
	//newBlock := NewBlock(height, prevBlockHash, data)
	//bc.Blocks = append(bc.Blocks, newBlock)
}
