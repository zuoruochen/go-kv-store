# go-kv-store
> a key-value store server and client
> Three value types: string ,map and list

## value type
+ string : key	value
+ map    : key	elem1,value1,elem2,value2,elem3,value3...
+ list   : key	value1,value2,value3,value4...

## command list
+ "help","list all directive"
+ "select","select db num"
+ "set", "set key value"
+ "get", "get value by key"
+ "exist", "if the key is exist"
+ "del", "delete key"
+ "cmap", "create a map key"
+ "mget", "get value from map key value"
+ "clist", "create a list key"
+ "lget", "get value from list key value,parm is index"