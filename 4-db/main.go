package main

import (
	"blockchain-go/3-pow/BLC"
	"fmt"
)

func main(){



	 //
	 //block :=BLC.NewBlock(1,nil,[]byte("my name is darren"))
	 //
	 //fmt.Printf(" block %v ",block)

	 blockChain :=BLC.CreateBlockChainWithGenesisBlock()

	 fmt.Printf("blockchain : %v \n",blockChain)

	 blockChain.AddBlock(blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
	 	  []byte("darren send 100btc to bob"),
	 	  blockChain.Blocks[len(blockChain.Blocks)-1].Hash)


	blockChain.AddBlock(blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
		[]byte("darren send 120btc to Alice"),
		blockChain.Blocks[len(blockChain.Blocks)-1].Hash)

	blockChain.AddBlock(blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
		[]byte("bob send 20btc to Alice"),
		blockChain.Blocks[len(blockChain.Blocks)-1].Hash)

	 lenth:=len(blockChain.Blocks)

	 fmt.Printf("length of block %d \n",lenth)

	 for i:=0;i<lenth; i++{
	 	fmt.Printf("the %d th block is %v\n",i,blockChain.Blocks[i])
	 }
}
