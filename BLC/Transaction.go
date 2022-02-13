package BLC

// 存放交易相关的
type Transaction struct {
	//tx hash  交易唯一的表示
	TxHash []byte
	// 输入
	Vins []*TxInput
	// 输出
	Vouts []*TxOutput
}

// 生成coinbase 交易 挖矿的交易 没有输入 由系统输入

// 生成转账交易
