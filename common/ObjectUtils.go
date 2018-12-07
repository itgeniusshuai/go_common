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

func ObjToMapByTagName(v interface{},tagName string) map[string]interface{}{
	var m = make(map[string]interface{},0)
	tye := reflect.TypeOf(v)
	value := reflect.ValueOf(v)
	var fieldNum = tye.NumField()
	for i := 0; i < fieldNum; i++{
		f := tye.Field(i)
		key := f.Name
		tagValue := f.Tag.Get(tagName)
		if tagValue != ""{
			key = tagName
		}
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

func InterfacePtrToInterface(v interface{})interface{}{
	var rv interface{}
	switch v1 := v.(type) {
	case *int:
		rv = *v1
	case *int8:
		rv = *v1
	case *int16:
		rv = *v1
	case *int32:
		rv = *v1
	case *int64:
		rv = *v1
	case *uint:
		rv = *v1
	case *uint8:
		rv = *v1
	case *uint16:
		rv = *v1
	case *uint32:
		rv = *v1
	case *float32:
		rv = *v1
	case *float64:
		rv = *v1
	case *string:
		rv = *v1
	}
	return rv
}



