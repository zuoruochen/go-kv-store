package db

import (
	"errors"
	"fmt"
	//	"reflect"
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

// all object should implement this interface
type Value interface {

	//get value
	Get() interface{}

	//set value,the origin value type is string
	//this method should turn the string type value into specific object type
	Set(value string) error

	//turn value into string type
	String() string

	Len() int
}

// new a MyDB struct
func NewMyDB() *MyDB {
	return &MyDB{
		DB: make(map[string]Object),
	}
}

// new a Object. Parm is a specific object type
// now there are three types object: *MapObj,*StrObj,*ListObj
// we can call NewMapObj(),NewStringObj() or NewListObj() to get a specific object
func NewObject(value Value) *Object {
	return &Object{
		Value: value,
	}
}

// what should be emphasized is that the return value's type is *Object
// however any operation on *Object would not affect the db
// there are two methods to affect db:  SetValue() and DelKey()
func (db *MyDB) GetValue(key string) (*Object, error) {
	obj, ok := db.DB[key]
	if ok {
		return &obj, nil
	} else {
		return nil, errors.New("the key is not in db!")
	}
}

func (db *MyDB) SetValue(key string, obj *Object) int {
	_, ok := db.DB[key]
	if ok {
		fmt.Println("The key has value,update value")
		db.DB[key] = *obj

		return 0
	} else {
		//		fmt.Printf("set value:%v\n", reflect.TypeOf(obj))
		db.DB[key] = *obj

		return 1
	}
}

// return val: 0 reps delete key successfully
//			   1  reps there is no key in db
func (db *MyDB) DelKey(key string) int {
	_, ok := db.DB[key]
	if ok {
		delete(db.DB, key)
		return 1
	}
	return 0
}
