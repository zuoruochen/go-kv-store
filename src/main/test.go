package main

import (
	"db"
	"fmt"
	"object"
	"reflect"
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
	mapobj.Set(" z , zzz , r, rrr,c,  ccc ")
	if val, ok := mapobj.GetElem("c"); ok {
		fmt.Println(val)
	}
	fmt.Println(mapobj.String())
	obj := db.NewObject(mapobj)

	fmt.Println(reflect.TypeOf(obj))
	cc := *obj
	mydb.SetValue("zrc", cc)
	value, err := mydb.GetValue("zzz")
	if err != nil {
		fmt.Printf("%s", reflect.TypeOf(value.Value))
	}
	return
}
