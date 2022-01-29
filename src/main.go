package main

import (
	"fmt"
	imcore "github.com/rainxue/goim/imcore"
)

func main() {
	fmt.Println("hello goim.")
	imcore.StartServer()
}