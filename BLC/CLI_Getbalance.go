package BLC

import "fmt"

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
