package common

import "reflect"

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