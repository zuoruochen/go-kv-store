package db

import (
	//"errors"
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
	db map[string]*object
	sync.RWMutex
}

const (
	CREATE = iota
	UPDATE
)


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
		db: make(map[string]*object),
	}
}

// new a object. Param is a specific object type
// now there are three types object: *MapObj,*StrObj,*ListObj
// we can call NewMapObj(),NewStringObj() or NewListObj() to get a specific object
func newObject(value Value) *object {
	return &object{
		rw:    new(sync.RWMutex),
		value: value,
	}
}

func (db *MyDB) GetValue(key string) (Value, error) {
	obj, ok := db.db[key]
	if ok {
		obj.rw.RLock()
		val := obj.value
		obj.rw.RUnlock()
		return val, nil
	} else {
		return nil, fmt.Errorf("[ %s ] is not in db!", key)
	}
}


func (db *MyDB) SetValue(key string, val Value) int {
	obj, ok := db.db[key]
	if ok {
		obj.rw.Lock()
		fmt.Println("The key has value,update value")
		obj.value = val
		obj.rw.Unlock()
		return UPDATE
	} else {
		//		fmt.Printf("set value:%v\n", reflect.TypeOf(obj))
		newobj := newObject(val)
		db.db[key] = newobj
		return CREATE
	}
}

// return val: true reps delete key successfully
//			   false  reps there is no key in db
func (db *MyDB) DelKey(key string) bool {
	_, ok := db.db[key]
	if ok {
		delete(db.db, key)
		return true
	}
	return false
}
