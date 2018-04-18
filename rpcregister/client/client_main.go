package main

import (
	"fmt"
	"github.com/Clodfisher/learngo/rpcregister"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	rpcClient := jsonrpc.NewClient(conn)

	var bIsRegisterOk bool
	err = rpcClient.Call("RegisterServer.RegusterUser",
		rpcregister.Args{Name: "yht", Passward: "123"},
		&bIsRegisterOk)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bIsRegisterOk, "yht 注册成功...")
	}

	err = rpcClient.Call("RegisterServer.RegusterUser",
		rpcregister.Args{Name: "ljj", Passward: "123"},
		&bIsRegisterOk)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bIsRegisterOk, "ljj 注册成功...")
	}

	err = rpcClient.Call("RegisterServer.RegusterUser",
		rpcregister.Args{Name: "yht", Passward: "123"},
		&bIsRegisterOk)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bIsRegisterOk, "yht 注册成功...")
	}

}
