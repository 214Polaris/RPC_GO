package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
	"time"

	function "./pkg/function"
)

type MethodArgs = function.MethodArgs
type MethodResult = function.MethodResult

type ServiceMethod struct {
	Name       string   // 方法名
	NumArgs    int      // 参数个数
	ResultType string   // 返回结果类型
	ArgType    []string // 参数类型
}

func main() {
	// 解析命令行参数s
	serverIP := flag.String("i", "", "服务端 IP 地址")
	serverPort := flag.String("p", "", "服务端端口号")
	flag.Parse()

	if *serverIP == "" || *serverPort == "" {
		log.Fatal("服务端 IP 地址和端口号不能为空")
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		// 创建建立连接
		client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%s", *serverIP, *serverPort))
		if err != nil {
			log.Fatal("Dial error:", err)
		}

		// 获取可调用方法的接口
		var serviceMethods []ServiceMethod
		err = client.Call("InterfaceInfo.GetServiceInfo", struct{}{}, &serviceMethods)
		if err != nil {
			log.Fatal("Call InterfaceInfo.GetServiceInfo error:", err)
		}

		log.Println("提供的服务如下:")
		for _, method := range serviceMethods {
			log.Println(method.Name)
		}

		// 获取要调用的方法名
		fmt.Print("请输入要调用的方法名 (输入 'end' 退出程序): ")
		methodName, _ := reader.ReadString('\n')
		methodName = strings.TrimSpace(methodName)

		if methodName == "end" {
			fmt.Println("退出程序")
			return
		}

		// 查找服务方法
		var method ServiceMethod
		for _, m := range serviceMethods {
			if m.Name == methodName {
				method = m
				break
			}
		}

		if method.Name == "" {
			log.Fatalf("无效的方法名: %s", methodName)
		}

		// 获取参数
		var methodArgs MethodArgs
		var result MethodResult

		for i := 0; i < method.NumArgs; i++ {
			fmt.Printf("请输入参数 %s: ", method.ArgType[i])
			input, _ := reader.ReadString('\n')
			input = strings.TrimSuffix(input, "\n") // 去掉换行符

			switch method.ArgType[i] {
			case "int":
				val := parseInputInt(input)
				// 根据参数的位置设置对应的参数值
				switch i {
				case 0:
					methodArgs.Args.A = val
				case 1:
					methodArgs.Args.B = val
				}
			case "string":
				// 根据参数的位置设置对应的参数值
				switch i {
				case 0:
					if method.NumArgs == 1 {
						methodArgs.Arg.A = input
					} else {
						methodArgs.Args.C = input
					}
				case 1:
					methodArgs.Args.D = input
				}
			}
		}

		// 开启一个管道，并用 Go 异步调用远端方法
		resultChan := make(chan *rpc.Call, 1)
		client.Go("Arithmetic."+methodName, methodArgs, &result, resultChan)

		// 等待异步调用完成
		call := <-resultChan

		// 检查异步调用是否发生错误
		if call.Error != nil {
			log.Fatalf("异步调用 %s 方法发生错误: %s", methodName, call.Error)
		} else {
			switch method.ResultType {
			case "int":
				log.Printf("异步调用 %s 方法的结果为: %d", methodName, result.ResultInt)
			case "string":
				log.Printf("异步调用 %s 方法的结果为: %s", methodName, result.ResultString)
			}
		}

		// 关闭客户端连接
		client.Close()

		time.Sleep(2 * time.Second)
	}
}

// 辅助函数：将输入解析为整数
func parseInputInt(input string) int {
	value := 0
	_, err := fmt.Sscan(input, &value)
	if err != nil {
		log.Fatal("无效的参数值:", err)
	}
	return value
}
