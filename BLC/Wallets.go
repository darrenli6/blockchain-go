package BLC

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"io/ioutil"
	"log"
	"os"
)

// 钱包集合的文件
// 存储钱包的文件
const walletFile = "Wallets.dat"

// 钱包的集合结构
type Wallets struct {
	// key : string => 钱包地址
	// value 钱包结构
	Wallets map[string]*Wallet
}

//  初始化 创建一个钱包的集合
func NewWallets() (*Wallets, error) {
	// 判断文件是否存在
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		wallets := &Wallets{}
		wallets.Wallets = make(map[string]*Wallet)
		return wallets, err
	}
	// 文件存在读取内容

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Panicf("get content failed %v \n", err)
	}
	var wallets Wallets
	// register适用于需要解析的参数中包含interface
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	if nil != err {
		log.Panicf("decode file failed %v \n", err)
	}

	return &wallets, nil
}

// 创建新的钱包

func (wallets *Wallets) CreateWallet() {
	wallet := NewWallet() // 新建钱包对象

	wallets.Wallets[string(wallet.GetAddress())] = wallet

	// 将钱包存储到文件中
	wallets.SaveWallets()

}

// 持久化钱包信息  写入硬盘中
func (w *Wallets) SaveWallets() {

	// 序列化钱包数据
	var content bytes.Buffer
	// 注册
	gob.Register(elliptic.P256())
	encoder := gob.NewEncoder(&content)

	// 序列化
	err := encoder.Encode(&w)
	if nil != err {
		log.Panicf("encode the struct of walltes failed %v \n", err)
	}
	//
	// r w x  4 2 1
	// 先清空文件 再存储 writeFile 函数 只保留一条数据 但是该数据存储会到目前位置所有地址的集合
	err = ioutil.WriteFile(walletFile, content.Bytes(), 0777)
	if err != nil {
		log.Printf("write the content of wallets to file [%s] failed %v \n", walletFile, err)
	}

}
