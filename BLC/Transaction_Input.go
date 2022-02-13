package BLC

// 交易输入
type TxInput struct {

	// 交易hash(不是当前的交易hash,上一笔的)
	TxHash []byte
	// 引用上一笔交易的output索引
	Vout int
	// 用户名 锁定脚本
	ScriptSig string
}
