package main

import (
	"bufio"
	//	"encoding/binary"
	"fmt"
	"net"
	"os"
	//	"strconv"
	"strings"
)

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

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println(err)

	reader := bufio.NewReader(os.Stdin)
	for {
		//command data
		line, _ := reader.ReadBytes('\n')
		data := line[:len(line)-1]
		/*
			fmt.Println(line)
			list := trimspace(string(line[:]))
			fmt.Println(list)
			com, _ := strconv.Atoi(list[0])
			test.command = uint8(com)
			test.data = []byte(list[1])
			test.data = test.data[:len(list[1])-1]
			binary.Write(conn, binary.LittleEndian, &test.command)
		*/

		conn.Write(data)
	}
	fmt.Println("exit client")
}