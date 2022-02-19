package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	msg := "this is the eg  f我 啦啦啦"

	encoded := base64.StdEncoding.EncodeToString([]byte(msg))

	fmt.Println(encoded)

	// 解码

	decode, err := base64.StdEncoding.DecodeString("dGhpcyBpcyB0aGUgZWcgIGbmiJEg5ZWm5ZWm5ZWm")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(decode))
}
