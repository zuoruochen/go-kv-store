package main

import (
	"db"
	"fmt"
	"object"
	//	"reflect"
	//"strings"
	"util"
)


func main() {
	in := "   fauns         sakjdkl  sakd "
	out := util.TrimSpace(in)
	fmt.Println(out)
	mydb := db.NewMyDB()

	mapobj := object.NewMapObj()
	errs := mapobj.Set(" z , zzz , r, rrr,c,mmm ")
	if errs != nil {
		panic(errs)
	}
	mydb.SetValue("zrc", mapobj)

	strobj := object.NewStringObj()
	strobj.Set("make you feel my love!")
	mydb.SetValue("qiqi", strobj)

	listobj := object.NewListObj()
	listobj.Set("all,out,of,love")
	listobj.Sort()
	mydb.SetValue("shanghai", listobj)

	value1, err := mydb.GetValue("zrc")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value1.String())
		fmt.Println(value1.(*object.MapObj).GetElem("c"))
	}

	value2, err := mydb.GetValue("qiqi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value2.String())

	}
	mydb.DelKey("shanghai")
	value3, err := mydb.GetValue("shanghai")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value3.String())
		fmt.Println(value3.(*object.ListObj).Vals(4))
	}

	return
}
