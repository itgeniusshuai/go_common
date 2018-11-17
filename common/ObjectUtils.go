package common

import "reflect"

func ObjToMapString(v interface{}) map[string]string{
	var m = make(map[string]string,0)
	tye := reflect.TypeOf(v)
	value := reflect.ValueOf(v)
	var fieldNum = tye.NumField()
	for i := 0; i < fieldNum; i++{
		f := tye.Field(i)
		key := f.Name
		v1 := value.Field(i).String()
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
		m[key] = v1.Interface()

	}
	return nil
}

func MapToObj(m map[string]string,v interface{},tagName string) {


}
