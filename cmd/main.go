package main

import "fmt"
import "../common"

func main()  {
	fmt.Println("hello go_common!!!")
	var num uint8 = 43
	str := common.Uint8ToStr(num)
	fmt.Println(str)
}