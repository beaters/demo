package main

import (
	"fmt"
	mypackage "github.com/beaters/utilModule/mypackage"
)
func main() {
	fmt.Println("开始调用mypackage的方法")
	msg := "input message"
	mypackage.Util4Print(&msg)
}