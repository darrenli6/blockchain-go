package BlockChain

// 钱包集合的文件

// 钱包的集合结构
type Wallets struct {
	// key : string => 钱包地址
	// value 钱包结构
	Wallets map[string]*Wallet
}

//  初始化 创建一个钱包的集合
func NewWallets() *Wallets {
	wallets := &Wallets{}
	wallets.Wallets = make(map[string]*Wallet)
	return wallets
}

// 创建新的钱包

func (wallets *Wallets) CreateWallet() {
	wallet := NewWallet() // 新建钱包对象

	wallets.Wallets[string(wallet.GetAddress())] = wallet

}
