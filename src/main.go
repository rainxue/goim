package main

import (
	"fmt"
	imcore "github.com/rainxue/goim/imcore"
	// api "github.com/rainxue/goim/api"
)

func main() {
	fmt.Println("hello goim.")
	
	// api.InitRouters()
	imcore.StartServer()
}