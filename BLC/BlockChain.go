package BLC

import (
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"strconv"
)

const dbName = "bc.db" // 存储数据的数据库文件

const blockTableName = "blocks" // 表名称

type BlockChain struct {
	DB  *bolt.DB // 数据库
	Tip []byte   // 最新区块的hash值

}

// 判断数据库是否存在
func dbExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}
	return true
}

// 初始化区块链

func CreateBlockChainWithGenesisBlock(address string) *BlockChain {

	if dbExists() {
		fmt.Println("创世区块已经存在")
		os.Exit(1) //退出
	}
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {

		log.Panicf("open  the db failed %v \n", err)
	}

	var blockHash []byte // 需要存储到数据库中的区块hash
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b == nil {
			b, err = tx.CreateBucket([]byte(blockTableName))
			if err != nil {
				log.Panicf("create bucket [%s] failed %v \n", blockTableName, err)
			}

		}
		if nil != b {

			// 生成交易
			txCoinbase := NewCoinbaseTransaction(address)

			// 创建创世区块
			genesisBlock := CreateGenesisBlock([]*Transaction{txCoinbase})
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

	if err != nil {
		log.Panicf("update the data of genesis block failed %v \n", err)
	}

	return &BlockChain{db, blockHash}
}

// 添加新区块到区块链中
func (bc *BlockChain) AddBlock(txs []*Transaction) {
	//newBlock := NewBlock(height, prevBlockHash, data)
	//bc.Blocks = append(bc.Blocks, newBlock)

	// 更新数据
	err := bc.DB.Update(func(tx *bolt.Tx) error {

		// 获取数据表
		b := tx.Bucket([]byte(blockTableName))

		if nil != b { // 明确表存在
			// 3 将获取最新区块的hash
			//newEastHash := b.Get([]byte("1"))

			blockBytes := b.Get(bc.Tip)
			lastest_block := DecerializeBlock(blockBytes)
			//4 创建新的区块
			newBlock := NewBlock(lastest_block.Height+1, lastest_block.Hash, txs)
			// 5. 存入数据库
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panicf("put the data of new block info db failed %v \n", err)
			}
			// 6 更新最新区块的值
			err = b.Put([]byte("1"), newBlock.Hash)
			if err != nil {
				log.Panicf("put the hash of new block info db failed %v \n", err)
			}

			bc.Tip = newBlock.Hash

		}
		return nil
	})

	if err != nil {
		log.Panicf("update the db of block failed , %v \n", err)

	}

}

// 遍历输出区块链中所有区块的信息
func (bc *BlockChain) PrintChain() {

	fmt.Println("区块链信息:")

	var curBlock *Block

	//var currentHash []byte = bc.Tip

	// 创建一个迭代器的对象
	bcit := bc.Iterator()

	for {
		fmt.Println("----------------------------------")
		curBlock = bcit.Next()
		fmt.Printf(" \t Height : %d \n", curBlock.Height)
		fmt.Printf(" \t TimeStamp : %d \n", curBlock.TimeStamp)
		fmt.Printf(" \t PrevBlockHash : %x \n", curBlock.PreBlockHash)
		fmt.Printf(" \t Hash : %x \n", curBlock.Hash)
		fmt.Printf(" \t Transaction : %x \n", curBlock.Txs)
		for _, tx := range curBlock.Txs {
			fmt.Printf("\t\t tx-hash : %x \n", tx.TxHash)
			fmt.Println("\t\t 输入..")
			for _, vin := range tx.Vins {
				fmt.Printf("\t\t\tvin-txhash: %v \n", vin.TxHash)
				fmt.Printf("\t\t\tvin-vout: %v \n", vin.Vout)
				fmt.Printf("\t\t\tvin-scripsig: %v \n", vin.ScriptSig)
			}
			fmt.Println("\t\t 输出..")
			for _, vout := range tx.Vouts {
				fmt.Printf("\t\t\tvout-value: %v \n", vout.Value)
				fmt.Printf("\t\t\tvout-ScriptPubkey: %v \n", vout.ScriptPubkey)

			}
		}
		fmt.Printf(" \t Nonce : %d \n", curBlock.Nonce)

		// 判断是否已经遍历到创世区块
		var hashInt big.Int
		hashInt.SetBytes(curBlock.PreBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}

		// 更新一下
		//currentHash = curBlock.PreBlockHash

	}

}

//返回blockchain对象
func BlockchainObject() *BlockChain {
	//读取数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if nil != err {
		log.Panicf("get the object of blockchain failed %v \n", err)
	}
	var tip []byte
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if nil != b {
			// 最新区块的hash value
			tip = b.Get([]byte("1")) // 最新的hash value

		}
		return nil
	})

	return &BlockChain{db, tip}

}

// 挖矿 生成新的区块，区块是通过挖矿产生的
// 接受交易 进行打包确认 最终生成新的区块
func (blockchain *BlockChain) MineNewBlock(from, to, amount []string) {
	fmt.Printf("\tFROM:[%s] \n", from)
	fmt.Printf("\tTO:[%s] \n", to)
	fmt.Printf("\tAMOUNT:[%s] \n", amount)
	// 接受交易
	var txs []*Transaction
	value, _ := strconv.Atoi(amount[0])
	tx := NewSimpleTransaction(from[0], to[0], value)
	txs = append(txs, tx)

	// 打包交易

	// 生成新的区块
	var block *Block
	// 从数据库中获取最新的区块
	blockchain.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			// 获取最新区块的hash value
			hash := b.Get([]byte("1"))
			// 得到最新的区块
			blockBytes := b.Get(hash) // 为了得到区块高度

			block = DecerializeBlock(blockBytes)

		}
		return nil

	})

	// 生成新的区块
	block = NewBlock(block.Height+1, block.Hash, txs)
	// 持久化一个新的区块
	blockchain.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if nil != b {
			err := b.Put(block.Hash, block.Serialize())
			if err != nil {
				log.Panicf("put block to db failed ! %v\n", err)
			}
			b.Put([]byte("1"), block.Hash) // 更新最新区块的hash value
			blockchain.Tip = block.Hash
		}

		return nil
	})

}

// 返回指定地址的余额
func (blockchain *BlockChain) UnUTXOS(address string) []*UTXO {

	fmt.Printf("the address is %s \n", address)

	var unUTXOS []*UTXO
	// 1. 遍历区块链,查找与address相关的所有交易
	// 1.1 获取区块链对象
	blockIterator := blockchain.Iterator()

	// 2.存储存储所有已花费的输出  make分配内存
	// key: 每个input所引用的交易hash
	// value: output 索引列表
	spentTxOutputs := make(map[string][]int)

	for {
		block := blockIterator.Next()  // 获取每一块区块信息
		for _, tx := range block.Txs { // 遍历每一区块的交易
			// 先查找输入
			if tx.IsCoinbaseTransaction() {
				// 转账交易的情况下 才查到输入
				for _, in := range tx.Vins {
					// 验证身份地址
					if in.UnLockWithAddress(address) {

						// 添加到已花费输入map中
						key := hex.EncodeToString(in.TxHash)
						spentTxOutputs[key] = append(spentTxOutputs[key], in.Vout)

					}
				}
			}

			// 再查找输出
			for index, vout := range tx.Vouts {
				// 判断地址验证 检查btc是否属于自己传入地址
				if vout.UnLockScriptPubkeyWithAddress(address) {
					// 是否是一个没有花费的输出
					// 判断已花费输出中是否为空
					if len(spentTxOutputs) != 0 {
						for txHash, indexArray := range spentTxOutputs {

							for _, i := range indexArray {
								if txHash == hex.EncodeToString(tx.TxHash) && i == index {
									//已经花费的输出
									continue
								} else {
									utxo := &UTXO{TxHash: tx.TxHash, Index: index, Output: vout}
									unUTXOS = append(unUTXOS, utxo)
								}

							}

						}
					} else {
						// 都是未花费的输出
						utxo := &UTXO{TxHash: tx.TxHash, Index: index, Output: vout}
						unUTXOS = append(unUTXOS, utxo)
					}
				}

			}
		}

		// 退出循环条件
		var hashInt big.Int
		hashInt.SetBytes(block.PreBlockHash)
		// 是否到创世区块这里
		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}

	}

	return unUTXOS
}

// 查询指定地址的余额

func (blockchain *BlockChain) getBalance(address string) int64 {
	utxos := blockchain.UnUTXOS(address)
	var amount int64
	for _, utxo := range utxos {
		amount += utxo.Output.Value
	}
	return amount
}
