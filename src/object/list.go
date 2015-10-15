package object

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type ListObj struct {
	list []string
}

func NewListObj() *ListObj {
	return &ListObj{
		list: make([]string, 0),
	}
}

func (list *ListObj) Len() int {
	return len(list.list)
}

func (list *ListObj) Get() interface{} {
	return list.list
}

// value string format: value1,value2,value3,value4...
func (list *ListObj) Set(value string) error {
	vals := strings.Split(value, ",")
	for _, v := range vals {
		v = strings.Trim(v, " ")
		list.Append(v)
	}
	return nil
}

func (list *ListObj) String() string {
	var ret string
	for idx, v := range list.list {
		if idx != len(list.list)-1 {
			ret = ret + fmt.Sprintf("%s%s", v, ",")
		} else {
			ret = ret + fmt.Sprintf("%s", v)
		}
	}
	return ret
}

func (list *ListObj) Less(i, j int) bool {
	return list.list[i] < list.list[j]
}

func (list *ListObj) Swap(i, j int) {
	list.list[i], list.list[j] = list.list[j], list.list[i]
}

func (list *ListObj) Append(elem ...string) {
	list.list = append(list.list, elem...)
}

func (list *ListObj) Sort() {
	sort.Sort(list)
}

func (list *ListObj) Push(elem string) {
	list.Append(elem)
}

func (list *ListObj) Pop() (string, error) {
	if list.Len() > 0 {
		ret := list.list[list.Len()-1]
		list.list = list.list[:(list.Len() - 1)]
		return ret, nil
	} else {
		return "", errors.New("list is nil")
	}
}

func (list *ListObj) Have(elem string) bool {
	for _, v := range list.list {
		if v == elem {
			return true
		}
	}
	return false
}
