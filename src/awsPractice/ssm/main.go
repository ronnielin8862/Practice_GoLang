package main

import (
	"flag"
	"fmt"
)

var cliName = flag.String("name", "nick", "Input Your Name")
var cliAge = flag.Int("age", 28, "Input Your Age")
var cliGender = flag.String("gender", "male", "Input Your Gender")

var cliFlag int

func Init() {
	flag.IntVar(&cliFlag, "flagname", 1234, "Just for demo")
}

func main() {
	// 初始化變量 cliFlag
	Init()
	// 把用戶傳遞的命令行參數解析為對應變量的值
	flag.Parse()

	// flag.Args() 函數返回沒有被解析的命令行參數
	// func NArg() 函數返回沒有被解析的命令行參數的個數
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}

	// 輸出命令行參數
	fmt.Println("name=", *cliName)
	fmt.Println("age=", *cliAge)
	fmt.Println("gender=", *cliGender)
	fmt.Println("flagname=", cliFlag)
}
