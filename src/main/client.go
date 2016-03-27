package main

import (
	"bufio"
	//	"encoding/binary"
	"fmt"
	"net"
	"os"
	//	"strconv"
	//"strings"
	//"util"
)

// TODO : try to realize pseudo terminal
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println(err)
	buf := make([]byte, 1024)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	for {
		//command data
		line, _ := reader.ReadBytes('\n')
		data := line[:len(line)-1]
		/*
			fmt.Println(line)
			list := util.TrimSpace(string(line[:]))
			fmt.Println(list)
			com, _ := strconv.Atoi(list[0])
			test.command = uint8(com)
			test.data = []byte(list[1])
			test.data = test.data[:len(list[1])-1]
			binary.Write(conn, binary.LittleEndian, &test.command)
		*/

		_,err := conn.Write(data)
		if err != nil {
			fmt.Println(err)
		}
		_,err = conn.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		writer.Write(buf)
		writer.Flush()
	}
	fmt.Println("exit client")
}
