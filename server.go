package main

import (
	"flag"
	"log"
	"net"
	"net/rpc"
	"sync"
	"time"

	function "./pkg/function"
)

func main() {
	// 读取命令行参数
	ip := flag.String("l", "0.0.0.0", "服务端监听的 IP 地址")
	port := flag.String("p", "", "服务端监听的端口号")
	flag.Parse()

	if *port == "" {
		log.Fatal("端口号不能为空")
	}

	interfaceInfo := new(function.InterfaceInfo)
	arithmetic := new(function.Arithmetic)

	server := rpc.NewServer()
	err := server.Register(interfaceInfo)
	if err != nil {
		log.Fatal("注册InterfaceInfo失败:", err)
	}

	err = server.Register(arithmetic)
	if err != nil {
		log.Fatal("注册Arithmetic失败:", err)
	}

	addr := *ip + ":" + *port
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("监听错误:", err)
	}

	log.Println("RPC服务器运行在地址:", addr)

	// 创建一个线程池
	pool := sync.Pool{
		New: func() interface{} {
			conn, err := l.Accept()
			if err != nil {
				log.Fatal("接收错误:", err)
			}
			// 设置连接的读取和写入超时时间为10秒
			err = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
			if err != nil {
				log.Fatal("设置读取超时时间失败:", err)
			}
			err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err != nil {
				log.Fatal("设置写入超时时间失败:", err)
			}
			return conn
		},
	}

	concurrentClients := make(chan bool, 10) // 限制并发处理的客户端数量为10个

	for {
		conn := pool.Get().(net.Conn) // 从线程池获取一个连接
		concurrentClients <- true     // 占用一个并发处理的客户端数量

		go func(c net.Conn) {
			defer func() {
				if r := recover(); r != nil {
					log.Println("连接异常:", r)
				}
				c.Close() // 关闭连接
			}()

			defer func() {
				pool.Put(c)         // 将连接放回线程池
				<-concurrentClients // 释放一个并发处理的客户端数量
			}()

			server.ServeConn(c) // 使用rpc.Server的ServeConn方法处理连接
		}(conn)
	}
}
