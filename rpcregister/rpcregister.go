package rpcregister

import "fmt"

type RegisterServer struct {
	RegisterMap map[string]string
}

type Args struct {
	Name     string
	Passward string
}

func (rs *RegisterServer) RegusterUser(args Args, result *bool) error {
	_, ok := rs.RegisterMap[args.Name]
	if ok {
		*result = false
		return fmt.Errorf("用户名 %s 已存在，请重新注册\n", args.Name)
	}

	rs.RegisterMap[args.Name] = args.Passward

	*result = true
	return nil
}
