package object

import (
	"fmt"
	"strings"
)

type MapObj struct {
	m map[string]string
}

//change  :type MapObj map[string]string  ?

func NewMapObj() *MapObj {
	return &MapObj{
		m: make(map[string]string),
	}
}

func (m *MapObj) GetElem(elem string) (string, bool) {
	if val, ok := m.m[elem]; ok {
		return val, true
	}
	return "", false
}

func (m *MapObj) DelElem(elem string) {
	delete(m.m, elem)
}

func (m *MapObj) SetElem(elem, value string) {
	m.m[elem] = value
}

// value string format: elem1,value1,elem2,value2,elem3,value3...
// the laset value if not exist, it is nil
func (m *MapObj) Get() interface{} {
	return m.m
}

func (m *MapObj) Set(value string) error {
	kv := strings.Split(value, ",")
	kvlen := len(kv)
	if kvlen%2 == 0 {
		for i := 0; i < kvlen; i = i + 2 {
			key := strings.Trim(kv[i], " ")
			val := strings.Trim(kv[i+1], " ")
			if _, ok := m.GetElem(key); !ok {
				m.SetElem(key, val)
			} else {
				return fmt.Errorf("[error] There is duplicate key : %s !", key)
			}
		}
		return nil
	} else {
		for i := 0; i < kvlen-1; i = i + 2 {
			key := strings.Trim(kv[i], " ")
			val := strings.Trim(kv[i+1], " ")
			if _, ok := m.GetElem(key); !ok {
				m.SetElem(key, val)
			} else {
				return fmt.Errorf("[error] There is duplicate key : %s !", key)
			}
		}
		key := strings.Trim(kv[kvlen-1], " ")
		if _, ok := m.GetElem(key); !ok {
			m.SetElem(key, "")
			return nil
		} else {
			return fmt.Errorf("[error] There is duplicate key : %s !", key)
		}
	}
}

func (m *MapObj) String() string {
	var ret string
	for k, v := range m.m {
		ret = ret + fmt.Sprintf("%s,%s,", k, v)
	}
	ret = strings.TrimRight(ret, ",")
	return ret
}

func (m *MapObj) Len() int {
	return len(m.m)
}
