package common

import "reflect"
import "fmt"

// 获取对象对应的着重号内的标签，key为原名称，v为标签名称
func GetStructTag(vt reflect.Type, tagName string) map[string]string{
	n := vt.NumField()
	resMap := make(map[string]string)
	for i := 0; i < n; i++{
		st := vt.Field(i)
		tag := st.Tag
		tj := tag.Get(tagName)
		resMap[st.Name] = tj
	}
	return resMap
}

func GetFieldTagValue(vt reflect.Type,fieldName string,tagName string) string{
	f,b := vt.FieldByName(fieldName)
	if b {
		return f.Tag.Get(tagName)
	}
	return ""
}

func GetValueByField(obj interface{},fieldName string) interface{}{
	vt := reflect.ValueOf(obj)
	f := vt.FieldByName(fieldName)
	return ValueToInterface(f)
}

func GetFieldValueByFieldTag(obj interface{},tagName string,tagValue string) interface{}{
	v := reflect.ValueOf(obj)
	vt := v.Type()
	n := vt.NumField()
	for i := 0; i < n; i++{
		fmt.Println(vt.Field(i).Tag.Get(tagName))
		if vt.Field(i).Tag.Get(tagName) == tagValue{
			return ValueToInterface(v.Field(i))
		}
	}
	return nil
}

func SetFieldByTagName(obj interface{}, tagName string, tagValue string, value interface{}){
	t := reflect.TypeOf(obj)
	tv := reflect.ValueOf(obj)
	n := t.NumField()
	for i := 0; i < n; i++{
		st := t.Field(i)
		tag := st.Tag
		tj := tag.Get(tagName)
		if tj == tagValue{
			tv.Field(i).Set(reflect.ValueOf(value))
		}
	}
}