package server

import (
	"db"
	"strconv"
)

/* Three types input
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
	dbNum,err := strconv.Atoi(conn.ReqData.Command)
	if err != nil {
		conn.RespData = "invalid string,input the num"
		return
	}
	conn.DB = MyServer.DB[dbNum]
	conn.RespData = "select db[" + dbNum + "]"
	return
}

func Exist(conn *Connection) {
	conn.DB.RLock()
	if _,err := conn.DB.GetValue(conn.ReqData.Key); err == nil {
		conn.RespData = "true"
	}else {
		conn.RespData = "false"
	}
	conn.DB.RUnlock()
}

func Get(conn *Connection) {
	conn.DB.RLock()
	value,err := conn.DB.GetValue(conn.ReqData.Key)
	conn.DB.RUnlock()
	if err != nil {
		conn.RespData = err.Error()
	}else {
		conn.RespData = value.String()
	}
}

func Del(conn *Connection) {
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

func Set(conn *Connection){
	conn.DB.Lock()
	ret := conn.DB.SetValue(conn.ReqData.Key,conn.ReqData.Data)
	conn.DB.Unlock()
	switch ret {
	case db.UPDATE : conn.RespData = conn.ReqData.Command + " Update " + conn.ReqData.Key
	case db.CREATE : conn.RespData = conn.ReqData.Command + " Create " + conn.ReqData.Key
	}
}

func Cmap(conn *Connection) {
	Set(conn)
}

func Clist(conn *Connection) {
	Set(conn)
}

//----------------------------------------------------------------------------------------------------------------------
