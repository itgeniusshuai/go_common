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

