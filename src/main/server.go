package main

import (
	//	"encoding/binary"
	"db"
	"fmt"
	"net"
	//"strings"
	//"object"
	"strings"
	"errors"
	"api/server"
)



/*
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
*/


func handleEvent(conn *server.Connection) {
	fmt.Println("in handle!")
	buf := make([]byte, 1024)
	defer conn.Conn.Close()

	for {
		n, err := conn.Conn.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("client closed!")
			} else {
				fmt.Println("read data error:", err)
			}
			return
		}
		//todo: parse command
		fmt.Println(string(buf[0:n]))
		commandLine := string(buf[0:n])
		command,key,data,err := getCommandAndData(commandLine)
		if err != nil {
			n,err = conn.Conn.Write([]byte(err.Error()))
			if err != nil {
				fmt.Println("write data to client error")
				return
			}
		}
		conn.ReqData = &server.Req{Command:command,Key:key,Data:data}
		//todo: exec command,before exec ,get the key's rwlock
		execCommand(conn)
	}
}


func execCommand(conn *server.Connection) (error) {
	return "",nil
}


// return command, key(if exist) ,data(if exist),error
func getCommandAndData(commandLine string)(string,string,string,error) {
	command,left,err := getWord(commandLine)
	if err != nil {
		return "","","",errors.New("command is null")
	}
	if _,ok := db.Cmd[strings.ToLower(command)]; !ok {
		return "","","",errors.New("unsupport command")
	}
	if command == "help" {
		return command,"","",nil
	}

	key,left,err := getWord(left)
	if command != "set" && command != "cmap" && command != "clist" {
		if err != nil {
			return command,"","",errors.New(command+" need args")
		}
		if left != "" {
			return command,key,"",errors.New(command + " " + key + " too many args")
		}
		return command,key,"",nil
	}
	data,left,err := getWord(left)
	if err != nil {
		return command,key,"",errors.New(command + " " + key + " need data")
	}
	if left != "" {
		return command, key, "", errors.New(command + " " + key + " too many args")
	}
	return command, key, data, nil
}

// get a word from a string
// return a word and left string
func getWord(str string)(string,string,error) {
	s := strings.TrimLeft(str," ")
	if len(s) == 0 {
		return "","",errors.New("blank string")
	}
	sl := strings.Split(s," ")
	word := sl[0]
	var idx = len(word)
	if idx == len(s) {
		return word,"",nil
	}
	return word,s[idx:len(s)],nil
}



func main() {
	//var test interface{} = "jsahdjakjd"


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
		go handleEvent(&server.Connection{Conn:conn,DB:server.MyServer.DB[0]})
	}

}
