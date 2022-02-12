package main

import (
	"time"
	"crypto/sha256"
	"encoding/hex"
	"github.com/labstack/echo"
	"fmt"
)



//区块模型
type  BlockModel struct {
	Id int64 //ID索引
	Timestamp string//区块创建的时间标识
	BPM  int //每分钟心跳频率
	Hash string//区块哈希sha256
	PreHash string //上一块的哈希sha256
}
//区块链，数组
var  BlockChain= make([]BlockModel,0)
//创建第一个区块，创世区块
func  init(){
	//创建了一个区块
	block:=BlockModel{0,time.Now().String(),0,"",""}
	BlockChain=append(BlockChain,block)//加入数组
}
//哈希处理
func calcHash(block BlockModel)string{
	record:=string(block.Id)+block.Timestamp+string(block.BPM)+block.PreHash//字符串
	myhash:=sha256.New()//创建算法，sha256对象
	myhash.Write([]byte(record))//加入数据
	hashed:=myhash.Sum(nil) //计算哈希
	return hex.EncodeToString(hashed) //编码为字符串


}

func Is_BlockValid(newBlock,lastBlock BlockModel) bool{
	//id不相等，不是顺序模式
	if  lastBlock.Id+1 !=newBlock.Id{
		return false
	}
	//前一块的哈希，不等于新区快的上一个哈希，
	if  lastBlock.Hash!=newBlock.PreHash{
		return false
	}
	if calcHash(newBlock)!=newBlock.Hash{//数据被纂改
		return false
	}


	return true
}



//处理区块的创建
func createBlock(ctx  echo.Context) error{
	//处理心跳信息
	type message struct {
		BPM int
	}
	var mymessage=message{}
	if  err:=ctx.Bind(&mymessage);err!=nil{//绑定消息处理
		panic(err)//处理错误
	}
	lastblock:=BlockChain[len(BlockChain)-1]//前一个区块
	//使用前一个区块，创建新的区块
	newblock:=BlockModel{}
	newblock.Id=lastblock.Id+1  //序列号+1
	newblock.Timestamp=time.Now().String() //当前时间
	newblock.BPM=mymessage.BPM   //心跳信息
	newblock.PreHash=lastblock.Hash  //哈希
	newblock.Hash=calcHash(newblock)  //计算哈希
	if Is_BlockValid(newblock,lastblock){
		BlockChain=append(BlockChain,newblock)
		fmt.Println("创建区块成功",BlockChain[len(BlockChain)-1].Id)
	}else{
		fmt.Println("创建区块失败")
	}

	return ctx.JSON(200,newblock)
}


func main(){
	echosever:=echo.New()//创建服务器
	echosever.GET("/", func(context echo.Context) error {
		return context.JSON(200,BlockChain)
	})
	echosever.POST("/",createBlock)
	echosever.Logger.Fatal(echosever.Start(":8848"))
}
