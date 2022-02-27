package BLC

type UTXO struct {

	//UTXO 所对应的交易hash

	TxHash []byte

	// UTXO 所在所属交易中的索引
	Index int

	// OUTPUT
	Output *TxOutput
}
