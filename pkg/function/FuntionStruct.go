package function

type ServiceMethod struct {
	Name       string   // 方法名
	NumArgs    int      // 参数个数
	ResultType string   // 返回结果类型
	ArgType    []string // 参数类型
}

func (i *InterfaceInfo) GetServiceInfo(_ struct{}, methods *[]ServiceMethod) error {
	methodsList := []ServiceMethod{
		{
			Name:       "Add",
			NumArgs:    2,
			ResultType: "int",
			ArgType:    []string{"int", "int"},
		},
		{
			Name:       "Multiply",
			NumArgs:    2,
			ResultType: "int",
			ArgType:    []string{"int", "int"},
		},
		{
			Name:       "ToCapital",
			NumArgs:    1,
			ResultType: "string",
			ArgType:    []string{"string"},
		},
		{
			Name:       "Subtract",
			NumArgs:    2,
			ResultType: "int",
			ArgType:    []string{"int", "int"},
		},
		{
			Name:       "Divide",
			NumArgs:    2,
			ResultType: "int",
			ArgType:    []string{"int", "int"},
		},
		{
			Name:       "Power",
			NumArgs:    2,
			ResultType: "int",
			ArgType:    []string{"int", "int"},
		},
		{
			Name:       "Concatenate",
			NumArgs:    2,
			ResultType: "string",
			ArgType:    []string{"string", "string"},
		},
		{
			Name:       "ToUpper",
			NumArgs:    1,
			ResultType: "string",
			ArgType:    []string{"string"},
		},
		{
			Name:       "ToLower",
			NumArgs:    1,
			ResultType: "string",
			ArgType:    []string{"string"},
		},
		{
			Name:       "Length",
			NumArgs:    1,
			ResultType: "int",
			ArgType:    []string{"string"},
		},
	}

	*methods = methodsList

	return nil
}

type ArgsDouble struct {
	A, B int
	C, D string
}

type ArgsSingle struct {
	A string
}

type MethodArgs struct {
	Args ArgsDouble
	Arg  ArgsSingle
}

type MethodResult struct {
	ResultInt    int
	ResultString string
}
