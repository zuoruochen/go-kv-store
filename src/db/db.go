package db

import (
	"errors"
	"fmt"
	//	"reflect"
	"sync"
)

type object struct {
	//if two or  more operations on this object ,we should lock the RW first
	rw    *sync.RWMutex
	value Value
}

type MyDB struct {
	DB map[string]*object
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
		DB: make(map[string]*object),
	}
}

// new a object. Parm is a specific object type
// now there are three types object: *MapObj,*StrObj,*ListObj
// we can call NewMapObj(),NewStringObj() or NewListObj() to get a specific object
func newObject(value Value) *object {
	return &object{
		rw:    new(sync.RWMutex),
		value: value,
	}
}

func (db *MyDB) GetValue(key string) (Value, error) {
	obj, ok := db.DB[key]
	if ok {
		obj.rw.Lock()
		val := obj.value
		obj.rw.Unlock()
		return val, nil
	} else {
		return nil, errors.New("the key is not in db!")
	}
}

func (db *MyDB) SetValue(key string, val Value) int {
	obj, ok := db.DB[key]
	if ok {
		obj.rw.Lock()
		fmt.Println("The key has value,update value")
		obj.value = val
		obj.rw.Unlock()
		return 0
	} else {
		//		fmt.Printf("set value:%v\n", reflect.TypeOf(obj))
		newobj := newObject(val)
		db.DB[key] = newobj
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
