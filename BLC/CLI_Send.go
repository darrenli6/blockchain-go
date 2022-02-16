package BLC

import (
	"fmt"
	"os"
)

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
