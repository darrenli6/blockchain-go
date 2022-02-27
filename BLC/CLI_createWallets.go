package BLC

import "fmt"

// 创建钱包集合
func (cli *CLI) CreateWallets() {
	walllets, _ := NewWallets()
	walllets.CreateWallet()
	fmt.Printf("wallets : %v \n", walllets)

}
