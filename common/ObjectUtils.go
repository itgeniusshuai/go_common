package common

import (
	"reflect"
	"fmt"
)

func ObjToMapString(v interface{}) map[string]string{
	var m = make(map[string]string,0)
	tye := reflect.TypeOf(v)
	value := reflect.ValueOf(v)
	var fieldNum = tye.NumField()
	for i := 0; i < fieldNum; i++{
		f := tye.Field(i)
		key := f.Name
		v1 := value.Field(i)
		m[key] = v1.String()
	}
	return m
}


func ObjToTagMapString(v interface{},tagName string) map[string]string{
	var m = make(map[string]string,0)
	tye := reflect.TypeOf(v)
	value := reflect.ValueOf(v)
	var fieldNum = tye.NumField()
	for i := 0; i < fieldNum; i++{
		f := tye.Field(i)
		key := f.Tag.Get(tagName)
		if key == ""{
			key = f.Name
		}
		v1 := fmt.Sprintf("%v",value.Field(i))
		m[key] = v1

	}
	return m
}

func ObjToMap(v interface{}) map[string]interface{}{
	var m = make(map[string]interface{},0)
	tye := reflect.TypeOf(v)
	value := reflect.ValueOf(v)
	var fieldNum = tye.NumField()
	for i := 0; i < fieldNum; i++{
		f := tye.Field(i)
		key := f.Name
		v1 := value.Field(i)
		m[key] = ValueToInterface(v1)
	}
	return m
}

func MapToObj(m map[string]interface{},v interface{}) {
	newObjValue := reflect.ValueOf(v).Elem()

	for k,v := range m{
		f := newObjValue.FieldByName(k)
		if f.IsValid(){
			f.Set(reflect.ValueOf(v))
		}
	}
}



