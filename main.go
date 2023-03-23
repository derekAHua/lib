package main

import "github.com/derekAHua/lib/server"

// @Author: Derek
// @Description:
// @Date: 2023/3/23 08:19
// @Version 1.0

func main() {
	httpServer := server.NewSdkHttpServer("demo", server.MetricFilterBuilder)
	err := httpServer.Start(":8080")
	if err != nil {
		panic(err)
	}
}
