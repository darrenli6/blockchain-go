package main

import (
	"blockchain-go/1/BLC"
	"fmt"
)

func main(){


	 block :=BLC.NewBlock(1,nil,[]byte("my name is darren"))

	 fmt.Printf(" block %v ",block)
}