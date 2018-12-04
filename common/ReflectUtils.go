package common

import (
	"reflect"
	"fmt"
)

func ValueToInterface(v reflect.Value)interface{}{
	var result interface{}
	switch v.Kind().String() {
		case "int":result = v.Int();break;
		case "int8":result = v.Int();break;
		case "int16":result = v.Int();break;
		case "int32":result = v.Int();break;
		case "int64":result = v.Int();break;
		case "uint":result = v.Uint();break;
		case "uint8":result = v.Uint();break;
		case "uint16":result = v.Uint();break;
		case "uint32":result = v.Uint();break;
		case "uint64":result = v.Uint();break;
		case "string":result = v.String();break;
		case "float32":result = v.Float();break;
		case "float64":result = v.Float();break;
	}
	return result
}

func FuncReflectExec(execFunc interface{},params ...interface{})([]interface{},error){
	vt := reflect.ValueOf(execFunc)
	fmt.Println(vt.Kind())
	paramValues := make([]reflect.Value,0)
	//if vt.Type().NumIn() != len(params){
	//	return nil,errors.New("参数个数不对")
	//}
	for _,param := range params{
		paramValues = append(paramValues, reflect.ValueOf(param))
	}
	resValues := vt.Call(paramValues)
	var res = make([]interface{},0)
	for _,resValue := range resValues{
		res = append(res, ValueToInterface(resValue))
	}
	return res,nil
}