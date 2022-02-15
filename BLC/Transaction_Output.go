package BLC

// 交易输出
type TxOutput struct {

	// 有多少钱 金额
	Value int64

	// 这个钱是谁的 用户名 解锁脚本
	ScriptPubkey string
}

// output 身份验证
func (txOutput *TxOutput) UnLockScriptPubkeyWithAddress(address string) bool {
	return address == txOutput.ScriptPubkey
}
