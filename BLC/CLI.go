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
	fmt.Printf("\tcreateblockchain -address  adress  -- 地址. \n")
	fmt.Printf("\taddblock -add DATA -- 交易数据 \n")
	fmt.Printf("\tprintchain -- 输出区块链的信息. \n")
	fmt.Printf("\tsend -from FROM -to TO -amount AMOUNT  -- 转账. \n")

	fmt.Printf("\tgetbalance -address FROM  -- 查询余额. \n")
}

// 校验 如果只输入了程序命令，就输出指令用户并且推出程序
func IsValidArgs() {
	if len(os.Args) < 2 {
		PrintUsage()
		//退出程序
		os.Exit(1)
	}
}

// 发送交易
func (cli *CLI) send(from, to, amount []string) {
	//检查交易
	if dbExists() == false {
		fmt.Println("数据库不存在")
		os.Exit(1)
	}
	blockchain := BlockchainObject()
	defer blockchain.DB.Close()
	blockchain.MineNewBlock(from, to, amount)

}

// 查询余额
func (cli *CLI) getBalance(from string) {
	// 获取指定地址的余额

	//outPuts := UnUTXOS(from)
	//fmt.Println("unUTXO : %v \n", outPuts)

	blockchain := BlockchainObject()

	defer blockchain.DB.Close()
	amount := blockchain.getBalance(from)

	fmt.Printf("\t地址： %s的余额为：%d \n", from, amount)
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

func (cli *CLI) createBlockChainWithGenesis(adress string) {

	CreateBlockChainWithGenesisBlock(adress)
}

// 运行函数
func (cli *CLI) Run() {
	// 检查参数
	IsValidArgs()

	// 2 新建命令

	// 添加区块
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	// 打印区块链的信息
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	// 创建创世区块
	crearteBlcWithGenesisCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	// 转账的命令
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	// 查询余额
	getBalanceCmd := flag.NewFlagSet("getBalanceCmd", flag.ExitOnError)

	// 获取命令行参数
	flagAddBlockArg := addBlockCmd.String("data", "send 100 BTC TO eveyone", "交易数据..")
	flagCreateBlockchainWithAddress := crearteBlcWithGenesisCmd.String("address", "", "地址..")

	//./main send -from "[\"darren\"]" -to "[\"lijia\"]" -amount  "[\"2\"]"
	flagFromArg := sendCmd.String("from", "", "转账源地址...")
	flagToArg := sendCmd.String("to", "", "转账目标地址...")
	flagAmountArg := sendCmd.String("amount", "", "转账金额")

	// 查询余额
	flagBalanceArg := getBalanceCmd.String("address", "", "查询地址..")

	// 转账命令行参数
	switch os.Args[1] {
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if nil != err {
			log.Panicf("parse cmd of send failed ! %v \n", err)

		}
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
	case "getbalance":

		err := getBalanceCmd.Parse(os.Args[2:])
		if nil != err {
			log.Panicf("parse getbalance  failed ! %v \n", err)
		}

	default:
		PrintUsage()
		os.Exit(1)

	}
	if getBalanceCmd.Parsed() {
		if *flagBalanceArg == "" {
			fmt.Println("未指定查询地址... ")
			os.Exit(1)
		}

		cli.getBalance(*flagBalanceArg)
	}

	// 添加转账命令
	if sendCmd.Parsed() {
		if *flagFromArg == "" {
			fmt.Println("源地址不能为空")
			PrintUsage()
			os.Exit(1)
		}
		if *flagToArg == "" {
			fmt.Println("目标地址不能为空")
			PrintUsage()
			os.Exit(1)
		}

		if *flagAmountArg == "" {
			fmt.Println("金额不能为空")
			PrintUsage()
			os.Exit(1)
		}
		fmt.Printf("\tFROM:[%s] \n", JSONToArray(*flagFromArg))
		fmt.Printf("\tTO:[%s] \n", JSONToArray(*flagToArg))
		fmt.Printf("\tAMOUNT:[%s] \n", JSONToArray(*flagAmountArg))

		cli.send(JSONToArray(*flagFromArg), JSONToArray(*flagToArg), JSONToArray(*flagAmountArg)) // 发送交易

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

		if *flagCreateBlockchainWithAddress == "" {
			PrintUsage()
			os.Exit(1)
		}
		cli.createBlockChainWithGenesis(*flagCreateBlockchainWithAddress)
	}
}
