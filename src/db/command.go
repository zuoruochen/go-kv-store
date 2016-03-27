package db

var commands [][]string = [][]string{
	{"help","list all directive"},
	{"select","select db num"},
	{"set", "set key value"},
	{"get", "get value by key"},
	{"exist", "if the key is exist"},
	{"del", "delete key"},
	{"cmap", "create a map key"},
	{"mget", "get value from map key value"},
	{"clist", "create a list key"},
	{"lget", "get value from list key value,param is index"},
}

var Cmd map[string]string

func init(){
	Cmd = make(map[string]string)
	for i := 0;i < len(commands); i++ {
		Cmd[commands[i][0]] = commands[i][1]
	}
}




