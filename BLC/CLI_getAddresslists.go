package BLC

import (
	"fmt"
)

func (cli *CLI) getAddressLists() {

	fmt.Println("打印所有钱包地址...")

	wallets, _ := NewWallets()

	for address, _ := range wallets.Wallets {
		fmt.Printf("address: [%s] \n ", address)
	}
}
