package function

import (
	"math"
	"strings"
	"time"
)

type InterfaceInfo struct{}

func (i *InterfaceInfo) GetInterfaceInfo(_ struct{}, reply *[]string) error {
	methods := []string{"Add", "Multiply", "ToCapital", "Subtract", "Divide", "Power", "Concatenate", "ToUpper", "ToLower", "Length"}
	*reply = methods
	return nil
}

type Arithmetic struct{}

func (a *Arithmetic) Add(args MethodArgs, reply *MethodResult) error {
	reply.ResultInt = args.Args.A + args.Args.B
	return nil
}

func (a *Arithmetic) Multiply(args MethodArgs, reply *MethodResult) error {
	reply.ResultInt = args.Args.A * args.Args.B
	return nil
}

func (a *Arithmetic) ToCapital(args MethodArgs, reply *MethodResult) error {
	reply.ResultString = strings.ToUpper(args.Arg.A)
	time.Sleep(10 * time.Second)
	return nil
}

// Subtract 减法
func (a *Arithmetic) Subtract(args MethodArgs, reply *MethodResult) error {
	reply.ResultInt = args.Args.A - args.Args.B
	return nil
}

// Divide 除法
func (a *Arithmetic) Divide(args MethodArgs, reply *MethodResult) error {
	if args.Args.B != 0 {
		reply.ResultInt = args.Args.A / args.Args.B
	} else {
		return nil // Handle division by zero error
	}
	return nil
}

// Power 次方
func (a *Arithmetic) Power(args MethodArgs, reply *MethodResult) error {
	reply.ResultInt = int(math.Pow(float64(args.Args.A), float64(args.Args.B)))
	return nil
}

// Concatenate 拼接字符串
func (a *Arithmetic) Concatenate(args MethodArgs, reply *MethodResult) error {
	reply.ResultString = args.Args.C + args.Args.D
	return nil
}

// ToUpper 大写
func (a *Arithmetic) ToUpper(args MethodArgs, reply *MethodResult) error {
	reply.ResultString = strings.ToUpper(args.Arg.A)
	return nil
}

// ToLower 小写
func (a *Arithmetic) ToLower(args MethodArgs, reply *MethodResult) error {
	reply.ResultString = strings.ToLower(args.Arg.A)
	return nil
}

// Length 判断字符串长度
func (a *Arithmetic) Length(args MethodArgs, reply *MethodResult) error {
	reply.ResultInt = len(args.Arg.A)
	return nil
}
