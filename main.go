package main

import "blockchain-go/4-pow/BLC"

func main() {

	blockChain := BLC.CreateBlockChainWithGenesisBlock()
	//BLC.PrintUsage()
	cli := BLC.CLI{blockChain}
	cli.Run()

}
