// TODO :  Add persistent storage command, periodically write data to disk
// TODO :  After start server, the server will load data in memory firstly

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
		fmt.Printf("Conection closed! Now there are %d connection remained\n",connected)
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

		fmt.Println(string(buf[0:n]))
		commandLine := string(buf[0:n])
		command,key,data,err := getCommandAndData(commandLine)
		fmt.Printf("command = %s , key = %s , data = %s\n",command,key,data)
		if err != nil {
			n,err = conn.Conn.Write([]byte(err.Error()))
			if err != nil {
				fmt.Println("write data to client error")
				return
			}
			continue
		}
		conn.ReqData = &server.Req{Command:command,Key:key,Data:data}

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
	/*
	data,left,err := util.GetWord(left)
	if err != nil {
		return command,key,"",errors.New(command + " " + key + " need data")
	}
	if left != "" {
		return command, key, "", errors.New(command + " " + key + " too many args")
	}
	*/
	return command, key, strings.TrimLeft(left," "), nil
}

func main() {
	//var test interface{} = "jsahdjakjd"
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening at 8888")
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		connected++
		fmt.Printf("New connection! There are %d connections\n",connected)
		go handleEvent(&server.Connection{Conn:conn,DB:server.MyServer.DB[0]})
	}

}
