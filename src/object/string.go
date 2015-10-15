package object

type StrObj struct {
	str string
}

func NewStringObj() *StrObj {
	return &StrObj{
		str: "",
	}
}

func (strobj *StrObj) Len() int {
	return len(strobj.str)
}

func (str1 *StrObj) Less(str2 *StrObj) bool {
	return str1.str < str2.str
}

func (strobj *StrObj) Get() interface{} {
	return strobj.str
}

func (strobj *StrObj) Set(value string) error {
	strobj.str = value
	return nil
}

func (strobj *StrObj) String() string {
	return strobj.str
}
