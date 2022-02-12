package main

import (
	"blockchain-go/4-pow/BLC"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {

	db, err := bolt.Open("bc.db", 0600, nil)
	if err != nil {
		log.Panicf("open  the db failed %v \n", err)
	}

	defer db.Close()

	genesisBlock := BLC.NewBlock(1, nil, []byte("test"))

	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("blocks"))
		if err != nil {
			log.Panicf("create bucket failed %v \n", err)
		}
		blockData := genesisBlock.Serialize() // 序列化的操作
		err = b.Put([]byte("1"), blockData)
		if nil != err {
			log.Panicf("put data to db failed ! error %v \n", err)
		}
		return nil

	})

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		if nil != b {
			data := b.Get([]byte("1"))
			fmt.Printf("data : %v \n", genesisBlock.DecerializeBlock(data))

		}
		return nil
	})
	if nil != err {
		log.Panicf("get data of block failed ! error %v \n", err)
	}

	//blockChain := BLC.CreateBlockChainWithGenesisBlock()
	//
	//fmt.Printf("blockchain : %v \n", blockChain)

	//blockChain.AddBlock(blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
	//	[]byte("darren send 100btc to bob"),
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	//
	//blockChain.AddBlock(blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
	//	[]byte("darren send 120btc to Alice"),
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	//
	//blockChain.AddBlock(blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
	//	[]byte("bob send 20btc to Alice"),
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Hash)

	//lenth := len(blockChain.Blocks)
	//
	//fmt.Printf("length of block %d \n", lenth)
	//
	//for i := 0; i < lenth; i++ {
	//	fmt.Printf("the %d th block is %v\n", i, blockChain.Blocks[i])
	//}
}
