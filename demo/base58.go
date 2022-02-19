package main

import (
	"bytes"
	"fmt"
	"math/big"
)

//base58字符表

var b58Alphabet = []byte("123456789" +
	"abcdefjhijkmnopqrstuvwxyz" +
	"ABCDEFGHJKLMNPQRSTUVWXYZ")

// 实现编码函数
func Base58Encode(input []byte) []byte {

	var result []byte
	x := big.NewInt(0).SetBytes(input)
	fmt.Printf("x: %v \n", x)
	base := big.NewInt(int64(len(b58Alphabet))) //设置一个基数 58位
	zero := big.NewInt(0)
	mod := &big.Int{} //余数

	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod) // 求余
		result = append(result, b58Alphabet[mod.Int64()])
	}
	// 翻转切片
	Reverse(result)

	for b := range input { // b代表切片下标
		if b == 0x00 {
			result = append([]byte{b58Alphabet[0]}, result...)
		} else {
			break
		}
	}

	fmt.Printf("result: %s \n", result)
	return result
}

// 解码函数

func Base58Decode(input []byte) []byte {

	result := big.NewInt(0)
	zeroBytes := 0

	for b := range input {
		if b == 0x00 {
			zeroBytes++
		}
	}
	data := input[zeroBytes:]

	for _, b := range data {

		// 获取bytes 数组中指定数字第一次出现的索引
		charIndex := bytes.IndexByte(b58Alphabet, b)
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))

	}
	decoded := result.Bytes()
	decoded = append(bytes.Repeat([]byte{byte(0x00)}, zeroBytes), decoded...)

	return decoded
}

func main() {
	msg := "1the massge 345"

	// 编码
	encoded := Base58Encode([]byte(msg))
	fmt.Printf("%x \n", encoded)

	//解码
	decode_data := Base58Decode(encoded)

	fmt.Printf("msg : %v \n", string(decode_data))
}

func Reverse(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}

}
