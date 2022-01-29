package imcore

import (
    "fmt"
    // "time"
    "github.com/gogf/gf/net/gtcp"
    // "github.com/gogf/gf/os/glog"
    // "github.com/gogf/gf/util/gconv"
)

func StartServer() {
	fmt.Println("server started.")
	gtcp.NewServer("127.0.0.1:8999", func(conn *gtcp.Conn) {
        defer conn.Close()
        for {
            data, err := conn.Recv(-1)
            if len(data) > 0 {
                  fmt.Println(string(data))
            }
            if err != nil {
                break
            }
        }
    }).Run()
}