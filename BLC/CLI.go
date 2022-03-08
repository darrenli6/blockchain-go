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
	fmt.Printf("\tcreatewallet  -- 创建钱包. \n")
	fmt.Printf("\taddresslists  -- 获取钱包地址列表. \n")
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

//
//// 添加区块
//
//func (cli *CLI) addBlock(txs []*Transaction) {
//
//	if dbExists() == false {
//		fmt.Println("数据库不存在")
//		os.Exit(1)
//	}
//	blockchain := BlockchainObject() //获取区块链对象
//	defer blockchain.DB.Close()
//	blockchain.AddBlock(txs)
//
//}
//

// 运行函数
func (cli *CLI) Run() {
	// 检查参数
	IsValidArgs()

	// 2 新建命令

	// // 添加区块
	// addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	// 创建钱包
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	// 获取钱包地址
	getAddressListCmd := flag.NewFlagSet("addresslists", flag.ExitOnError)
	// 打印区块链的信息
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	// 创建创世区块
	crearteBlcWithGenesisCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	// 转账的命令
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	// 查询余额
	getBalanceCmd := flag.NewFlagSet("getBalanceCmd", flag.ExitOnError)

	// 获取命令行参数
	//flagAddBlockArg := addBlockCmd.String("data", "send 100 BTC TO eveyone", "交易数据..")
	flagCreateBlockchainWithAddress := crearteBlcWithGenesisCmd.String("address", "", "地址..")

	//./main send -from "[\"darren\"]" -to "[\"lijia\"]" -amount  "[\"2\"]"

	// 多人转
	//./main send -from "[\"darren\",\"lijia\"]" -to "[\"lijia\",\"darren\"]" -amount  "[\"2\",\"1\"]"
	flagFromArg := sendCmd.String("from", "", "转账源地址...")
	flagToArg := sendCmd.String("to", "", "转账目标地址...")
	flagAmountArg := sendCmd.String("amount", "", "转账金额")

	// 查询余额
	flagBalanceArg := getBalanceCmd.String("address", "", "查询地址..")

	// 转账命令行参数
	switch os.Args[1] {
	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if nil != err {
			log.Panicf("parse cmd of createWallet failed ! %v \n", err)
		}
	case "addresslists":
		err := getAddressListCmd.Parse(os.Args[2:])
		if nil != err {
			log.Panicf("parse cmd of getAddressList failed ! %v \n", err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if nil != err {
			log.Panicf("parse cmd of send failed ! %v \n", err)

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

	//解析
	if getAddressListCmd.Parsed() {

		cli.getAddressLists()
	}

	// 创建钱包
	if createWalletCmd.Parsed() {
		cli.CreateWallets()
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

		cli.send(JSONToArray(*flagFromArg), JSONToArray(*flagToArg), JSONToArray(*flagAmountArg)) // 发送交易

	}
	//// 添加区块链
	//if addBlockCmd.Parsed() {
	//	if *flagAddBlockArg == "" {
	//		PrintUsage()
	//		os.Exit(1)
	//	}
	//
	//	cli.addBlock([]*Transaction{})
	//}

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
