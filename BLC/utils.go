package BLC

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
)

func IntToHex(data int64) []byte {

	buffer := new(bytes.Buffer)

	err := binary.Write(buffer, binary.BigEndian, data)
	if nil != err {
		log.Panicf("int to []byte failed %v \n", err)
	}
	return buffer.Bytes()
}

// 标准json格式转组数
// ./main send -from "[\"lijia\",\"lisi\"]" -to "[\"darren\",\"zhangsan\"]" -amount  "[\"10\",\"1\"]"
func JSONToArray(jsonString string) []string {

	var strArr []string
	// json.unmarshal
	if err := json.Unmarshal([]byte(jsonString), &strArr); err != nil {
		log.Panicf("json to []string failed ! %v\n", err)
	}

	return strArr
}

func Reverse1(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}

}
