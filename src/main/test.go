package main

import (
	"db"
	"fmt"
	"object"
	//	"reflect"
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
	in := "   fauns         sakjdkl  sakd "
	out := trimspace(in)
	fmt.Println(out)
	mydb := db.NewMyDB()

	mapobj := object.NewMapObj()
	errs := mapobj.Set(" z , zzz , r, rrr,c,mmm ")
	if errs != nil {
		panic(errs)
	}
	obj1 := db.NewObject(mapobj)
	mydb.SetValue("zrc", obj1)

	strobj := object.NewStringObj()
	strobj.Set("make you feel my love!")
	obj2 := db.NewObject(strobj)
	mydb.SetValue("dongjia", obj2)

	listobj := object.NewListObj()
	listobj.Set("all,out,of,love")
	listobj.Sort()
	obj3 := db.NewObject(listobj)
	mydb.SetValue("shanghai", obj3)

	value1, err := mydb.GetValue("zrc")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value1.Value.String())
		fmt.Println(value1.Value.(*object.MapObj).GetElem("c"))
	}

	value2, err := mydb.GetValue("dongjia")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value2.Value.String())

	}

	value3, err := mydb.GetValue("shanghai")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value3.Value.String())
		fmt.Println(value3.Value.(*object.ListObj).Vals(4))
	}

	return
}
