package BLC

import (
	"fmt"
	"os"
)

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
