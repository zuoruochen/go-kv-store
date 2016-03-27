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
	"util"
)

var connected = 0

func handleEvent(conn *server.Connection) {
	fmt.Println("in handle!")
	buf := make([]byte, 1024)
	defer func() {
		connected--
		fmt.Printf("Conection closed! Now there are %d connection remained",connected)
		conn.Conn.Close()
	}()
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
		err = server.ExecCommand(conn)
		if err != nil {
			return
		}
	}
}


// return command, key(if exist) ,data(if exist),error
func getCommandAndData(commandLine string)(string,string,string,error) {
	command,left,err := util.GetWord(commandLine)
	command = strings.ToLower(command)
	if err != nil {
		return "","","",errors.New("command is null")
	}
	if _,ok := db.Cmd[command]; !ok {
		return "","","",errors.New("unsupport command")
	}
	if command == "help" {
		return command,"","",nil
	}

	key,left,err := util.GetWord(left)
	if command != "set" && command != "cmap" && command != "clist" {
		if err != nil {
			return command,"","",errors.New(command+" need args")
		}
		if left != "" {
			return command,key,"",errors.New(command + " " + key + " too many args")
		}
		return command,key,"",nil
	}
	data,left,err := util.GetWord(left)
	if err != nil {
		return command,key,"",errors.New(command + " " + key + " need data")
	}
	if left != "" {
		return command, key, "", errors.New(command + " " + key + " too many args")
	}
	return command, key, data, nil
}

func main() {
	//var test interface{} = "jsahdjakjd"
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		connected++
		fmt.Printf("New connection! There are %d connections",connected)
		go handleEvent(&server.Connection{Conn:conn,DB:server.MyServer.DB[0]})
	}

}
