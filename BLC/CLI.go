package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// CLI 结构

type CLI struct {
	BC *BlockChain
}

// 展示用法
func PrintUsage() {
	fmt.Println("Usage: ")
	fmt.Printf("\tcreateblcwothgenensis -- 创建区块链. \n")
	fmt.Printf("\taddblock -add DATA -- 交易数据 \n")
	fmt.Printf("\tprintblockchain -- 输出区块链的信息. \n")

}

// 校验 如果只输入了程序命令，就输出指令用户并且推出程序
func IsValidArgs() {
	if len(os.Args) < 2 {
		PrintUsage()
		//退出程序
		os.Exit(1)
	}
}

// 添加区块

func (cli *CLI) addBlock(txs []*Transaction) {

	if dbExists() == false {
		fmt.Println("数据库不存在")
		os.Exit(1)
	}
	blockchain := BlockchainObject() //获取区块链对象
	defer blockchain.DB.Close()
	blockchain.AddBlock(txs)

}

// 输出区块链信息

func (cli *CLI) printchain() {
	if dbExists() == false {
		fmt.Println("数据库不存在")
		os.Exit(1)
	}
	blockchain := BlockchainObject() //获取区块链对象
	defer blockchain.DB.Close()
	blockchain.PrintChain()
}

// 创建区块链

func (cli *CLI) createBlockChainWithGenesis(txs []*Transaction) {

	CreateBlockChainWithGenesisBlock(txs)
}

// 运行函数
func (cli *CLI) Run() {
	// 检查参数
	IsValidArgs()

	// 2 新建命令

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	crearteBlcWithGenesisCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)

	// 获取命令行参数
	flagAddBlockArg := addBlockCmd.String("data", "send 100 BTC TO eveyone", "交易数据")
	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if nil != err {
			log.Panicf("parse cmd of addblock falield ! %v", err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if nil != err {
			log.Panicf("parse cmd of printchain falield ! %v", err)
		}

	case "createblockchain":
		err := crearteBlcWithGenesisCmd.Parse(os.Args[2:])
		if nil != err {
			log.Panicf("parse cmd of createblockchain  falield ! %v", err)
		}

	default:
		PrintUsage()
		os.Exit(1)

	}

	// 添加区块链
	if addBlockCmd.Parsed() {
		if *flagAddBlockArg == "" {
			PrintUsage()
			os.Exit(1)
		}

		cli.addBlock([]*Transaction{})
	}

	// 输出区块链信息命令
	if printChainCmd.Parsed() {
		cli.printchain()
	}

	// 创建区块链
	if crearteBlcWithGenesisCmd.Parsed() {
		cli.createBlockChainWithGenesis([]*Transaction{})
	}
}
