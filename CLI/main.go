package main

import (
	"flag"
	"fmt"
)

////命令行操作
//func main() {
//
//	args := os.Args
//
//	if len(args) >= 2 {
//		for i := 0; i < len(args); i++ {
//			fmt.Printf("args[%d]: %v  \n", i, args[i])
//		}
//	} else {
//		fmt.Printf("arg[0]: %v \n", args[0])
//
//	}
//}

func main() {

	flagPrintChain := flag.String("printchain", "ls", "print the info ")
	flag.Parse()
	fmt.Printf("the flag of string : %v \n", *flagPrintChain)

}