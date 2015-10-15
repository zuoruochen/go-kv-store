package db

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type Object struct {
	//	Types    int
	RW    *sync.RWMutex
	Value Value
}

type MyDB struct {
	DB map[string]Object
}

type Value interface {
	//get value
	Get() interface{}
	//set value
	Set(value string) error
	//turn value to string type
	String() string
	Len() int
}

func NewMyDB() *MyDB {
	return &MyDB{
		DB: make(map[string]Object),
	}
}

func NewObject(value Value) *Object {
	return &Object{
		Value: value,
	}
}

func (db *MyDB) GetValue(key string) (*Object, error) {
	obj, ok := db.DB[key]
	if ok {
		return &obj, nil
	} else {
		return nil, errors.New("the key is not in db!")
	}
}

func (db *MyDB) SetValue(key string, obj Object) int {
	_, ok := db.DB[key]
	if ok {
		fmt.Println("The key had value,update value")
		db.DB[key] = obj

		return 0
	} else {
		fmt.Printf("set value:%v\n", reflect.TypeOf(obj))
		db.DB[key] = obj

		return 1
	}
}
