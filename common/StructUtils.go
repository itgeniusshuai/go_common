package common

import "reflect"

// 获取对象对应的着重号内的标签，key为原名称，v为标签名称
func GetStructTag(obj interface{}, tagName string) map[string]string{
	t := reflect.TypeOf(obj)
	n := t.NumField()
	resMap := make(map[string]string)
	for i := 0; i < n; i++{
		st := t.Field(i)
		tag := st.Tag
		tj := tag.Get(tagName)
		resMap[st.Name] = tj
	}
	return resMap
}

func GetFieldTagValue(obj interface{},fieldName string,tagName string) string{
	vt := reflect.TypeOf(obj)
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
		if vt.Field(i).Tag.Get(tagName) == tagValue{
			return ValueToInterface(v.Field(i))
		}
	}
	return nil
}

