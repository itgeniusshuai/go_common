package common

import (
	"strconv"
	"bytes"
)

// 是否是空串
func IsNull(str string) bool{
	return str == ""
}

// 是否不为空串
func IsNotNull(str string) bool{
	return !IsNull(str)
}

// 字符串转int
func StrToInt(str string) (int, error){
	return strconv.Atoi(str)

}

// 字符串转int8
func StrToInt8(str string, base int) (int8, error){
	num,err := strconv.Atoi(str)
	if err != nil{
		return 0,err
	}
	return int8(num),nil
}

// 字符串转int16
func StrToInt16(str string) (int16, error){
	num,err := strconv.Atoi(str)
	if err != nil{
		return 0,err
	}
	return int16(num),nil
}

// 字符串转int16
func StrToInt32(str string) (int32, error){
	num,err := strconv.Atoi(str)
	if err != nil{
		return 0,err
	}
	return int32(num),nil
}

// 字符串转int64
func StrToInt64(str string) (int64, error){
	return strconv.ParseInt(str,10,64)
}

// 字符串转int
func StrToUint(str string) (uint, error){
	num,err := strconv.Atoi(str)
	if err != nil{
		return 0,err
	}
	return uint(num),err

}

// 字符串转uint8
func StrToUint8(str string) (uint8, error){
	num,err := strconv.ParseUint(str,10,8)
	if err != nil{
		return 0,err
	}
	return uint8(num),nil
}

// 字符串转uint16
func StrToUint16(str string) (uint16, error){
	num,err := strconv.ParseUint(str,10,16)
	if err != nil{
		return 0,err
	}
	return uint16(num),nil
}

// 字符串转uint16
func StrToUint32(str string) (uint32, error){
	num,err := strconv.ParseUint(str,10,32)
	if err != nil{
		return 0,err
	}
	return uint32(num),nil
}

// 字符串转uint64
func StrToUint64(str string) (uint64, error){
	num,err := strconv.ParseUint(str,10,64)
	if err != nil{
		return 0,err
	}

	return uint64(num),err
}

// 字符串转float64
func StrToFloat64(str string) (float64, error){
	return strconv.ParseFloat(str, 64)
}

// 字符串转float32
func StrToFloat32(str string) (float32, error){
	num, err := strconv.ParseFloat(str, 32)
	if err != nil{
		return 0,err
	}
	return float32(num),nil
}

// float64转string
func Float64ToStr(f float64) string{
	return strconv.FormatFloat(f, 'E', -1, 64)
}

// float32转string
func Float32ToStr(f float32) string{
	return Float64ToStr(float64(f))
}



// []uint8转string
func Uint8sToStr(us []uint8) string{
	ba := []byte{}
	for _, b := range us {
		ba = append(ba, byte(b))
	}
	return string(ba)
}

// []uint转string
func UintsToStr(us []uint) string{
	ba := []byte{}
	for _, b := range us {
		ba = append(ba, byte(b))
	}
	return string(ba)
}

// int转string
func IntToStr(num int) string{
	return strconv.Itoa(num)
}

// int8转string
func Int8ToStr(num int8) string{
	return IntToStr(int(num))
}

// int8转string
func Int16ToStr(num int16) string{
	return IntToStr(int(num))
}

// int8转string
func Int32ToStr(num int32) string{
	return IntToStr(int(num))
}

// int8转string
func Int64ToStr(num int64) string{
	return strconv.FormatInt(num,64)
}

// int转string
func UintToStr(num uint) string{
	return IntToStr(int(num))
}

// int8转string
func Uint8ToStr(num uint8) string{
	return IntToStr(int(num))
}

// int8转string
func Uint16ToStr(num uint16) string{
	return IntToStr(int(num))
}

// int8转string
func Uint32ToStr(num uint32) string{
	return IntToStr(int(num))
}

// int8转string
func Uint64ToStr(num uint64) string{
	return strconv.FormatUint(num,64)
}

// string转[]byte
func StrToBytes(str string) []byte{
	return []byte(str)
}

// string转[]uint8
func StrToUint8s(str string) []uint8{
	return []uint8(str)
}

func IntToStrHex(num int) string{
	return strconv.FormatInt(int64(num),16)
}

func IntsToStrHex(nums []int) string{
	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲

	for _,num := range nums{
		buffer.WriteString(IntToStrHex(num))
	}
	return buffer.String()
}

func ByteToStr(b byte) string{
	return string(b)
}

func ByteToStrHex(b byte) string{
	return IntToStrHex(int(b))
}

// []byte转string
func BytesToStr(bs []byte) string{
	return string(bs)
}

func BytesToStrHex(bs []byte) string{
	var buffer bytes.Buffer
	for _,b := range bs{
		buffer.WriteString(ByteToStrHex(b))
	}
	return buffer.String()
}

