package server

import  (
	"db"
	"net"
)

const DBnum int = 16

var MyServer *Zserver

type Zserver struct {
	DB [DBnum]*db.MyDB
}

type Req struct {
	Command 	string
	Key			string
	Data		string
}


type Connection struct {
	Conn		net.Conn
	DB 			*db.MyDB
	ReqData	*Req
	RespData	string
}


func init(){
	MyServer = &Zserver{}
	for i:=0 ; i < DBnum; i++ {
		MyServer.DB[i] = db.NewMyDB()
	}
}