package server

import (
	"db"
	"fmt"
	"strconv"
	"object"
	"reflect"
)

/* Three types of input
 * command
 * command + key
 * command + key + value
*/

//---------------------------------------------------no key----------------------------------------------------------------

func Help() string {
	data := ""
	for k,v := range db.Cmd {
		data = data +  k + ": " + "\"" + v + "\"" + "\n"
	}
	return data
}

//------------------------------------------------one key no value--------------------------------------------------------

func Select(conn *Connection) {
	fmt.Println("Select")
	dbNum,err := strconv.Atoi(conn.ReqData.Key)
	if err != nil {
		conn.RespData = "invalid string,input the num"
		return
	}
	conn.DB = MyServer.DB[dbNum]
	conn.RespData = "select db[" + strconv.Itoa(dbNum) + "]"
	return
}

func Exist(conn *Connection) {
	fmt.Println("Exist")
	conn.DB.RLock()
	if _,err := conn.DB.GetValue(conn.ReqData.Key); err == nil {
		conn.RespData = "true"
	}else {
		conn.RespData = "false"
	}
	conn.DB.RUnlock()
}

func Get(conn *Connection) {
	fmt.Println("Get")
	conn.DB.RLock()
	value,err := conn.DB.GetValue(conn.ReqData.Key)
	conn.DB.RUnlock()
	if err != nil {
		conn.RespData = err.Error()
	}else {
		conn.RespData = value.String()
	}
	fmt.Println(conn.RespData)
}

func Del(conn *Connection) {
	fmt.Println("Del")
	conn.DB.Lock()
	ok := conn.DB.DelKey(conn.ReqData.Key)
	conn.DB.Unlock()
	if ok {
		conn.RespData = "true"
	}else {
		conn.RespData = "false"
	}
}

func Mget(conn *Connection) {
	Get(conn)
}

func Lget(conn *Connection) {
	Get(conn)
}

//----------------------------------------------------key and value-----------------------------------------------------

func Set(conn *Connection) {
	fmt.Println("Set")
	conn.DB.Lock()
	value, err := conn.DB.GetValue(conn.ReqData.Key)
	if err == nil {
		if reflect.TypeOf(value).String() != "*object.StrObj" {
			conn.RespData = conn.ReqData.Key +  " is not a string key"
			conn.DB.Unlock()
			return
		}
	}
	strobj := object.NewStringObj()
	strobj.Set(conn.ReqData.Data)
	ret := conn.DB.SetValue(conn.ReqData.Key, strobj)
	conn.DB.Unlock()
	switch ret {
	case db.UPDATE : conn.RespData = conn.ReqData.Command + " Update " + conn.ReqData.Key
	case db.CREATE : conn.RespData = conn.ReqData.Command + " Create " + conn.ReqData.Key
	}
}

func Cmap(conn *Connection) {
	fmt.Println("Cmap")
	conn.DB.Lock()
	value, err := conn.DB.GetValue(conn.ReqData.Key)
	if err == nil {
		if reflect.TypeOf(value).String() != "*object.MapObj" {
			conn.RespData = conn.ReqData.Key +  " is not a map key"
			conn.DB.Unlock()
			return
		}
	}
	mapobj := object.NewMapObj()
	mapobj.Set(conn.ReqData.Data)
	ret := conn.DB.SetValue(conn.ReqData.Key,mapobj)
	conn.DB.Unlock()
	switch ret {
	case db.UPDATE : conn.RespData = conn.ReqData.Command + " Update " + conn.ReqData.Key
	case db.CREATE : conn.RespData = conn.ReqData.Command + " Create " + conn.ReqData.Key
	}
}

func Clist(conn *Connection) {
	fmt.Println("Clist")
	conn.DB.Lock()
	value, err := conn.DB.GetValue(conn.ReqData.Key)
	if err == nil {
		if reflect.TypeOf(value).String() != "*object.ListObj" {
			conn.RespData = conn.ReqData.Key +  " is not a list key"
			conn.DB.Unlock()
			return
		}
	}
	listobj := object.NewListObj()
	listobj.Set(conn.ReqData.Data)
	ret := conn.DB.SetValue(conn.ReqData.Key,listobj)
	conn.DB.Unlock()
	switch ret {
	case db.UPDATE : conn.RespData = conn.ReqData.Command + " Update " + conn.ReqData.Key
	case db.CREATE : conn.RespData = conn.ReqData.Command + " Create " + conn.ReqData.Key
	}
}

//----------------------------------------------------------------------------------------------------------------------
//TODO: try to use reflect.methodbyname to call the command,but this would make the commands as the methods of the db.MyDB object
func ExecCommand(conn *Connection) error{
	switch conn.ReqData.Command {
	case "help"		:	Help()
	case "select"	:	Select(conn)
	case "exist"	:	Exist(conn)
	case "get"		:	Get(conn)
	case "del"		:	Del(conn)
	case "mget"		:	Mget(conn)
	case "lget"		:	Lget(conn)
	case "set"		:	Set(conn)
	case "cmap"		:	Cmap(conn)
	case "clist"	:	Clist(conn)
	default:
		Help()
	}
	_,err := conn.Conn.Write([]byte(conn.RespData))
	if err != nil {
		fmt.Println("write data to client error")
		return err
	}
	return nil
}