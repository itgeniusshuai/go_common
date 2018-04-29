package common

import "encoding/json"

// 对象转json
func JsonObjToStr(v interface{}) (string,error){
	bs,err := json.Marshal(v)
	if err != nil{
		return "",err
	}
	return BytesToStr(bs),nil
}

// json转对象
func JsonStrToObj(str string, v interface{}) error{
	err := json.Unmarshal(StrToBytes(str),v)
	return err
}

