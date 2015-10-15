package main

import (
	//	"encoding/binary"
	"db"
	"fmt"
	"net"
	"strings"
)

var dbnum int = 16

type zserver struct {
	db [dbnum]*db.MyDB
}

func trimspace(s string) (ret []string) {
	s = strings.Trim(s, " ")
	i := 0
	j := 0

	for i < len(s) {
		if s[i] == ' ' {
			ret = append(ret, s[j:i])
			i++
			for s[i] == ' ' {
				i++
			}
			j = i
			i++
		} else {
			i++
		}
	}
	ret = append(ret, s[j:i])
	return ret
}

func handleEvent(conn net.Conn) {
	fmt.Println("in handle!")
	buf := make([]byte, 1024)
	defer conn.Close()
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("client closed!")
			} else {
				fmt.Println("read data error:", err)
			}
			return
		}
		//todo: parse command
		//todo: exec command,before exec ,get the key's rwlock
		fmt.Println(string(buf[0:n]))
	}
}

func main() {
	var test interface{} = "jsahdjakjd"
	var obj db.Object
	obj.Obj = &test
	obj.Types = 1
	fmt.Println("hello world!")
	fmt.Println(*obj.Obj)

	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("new connection!")
		go handleEvent(conn)
	}

}
