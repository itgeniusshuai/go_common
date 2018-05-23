package common

// 小端法
func BytesToIntWithMin(bs []byte) int{
	bLen := len(bs)
	if(bLen > 4){
		bLen = 4
	}
	var i int
	var j uint
	for j = 0; j < uint(bLen); j++{
		i |= int(bs[j])<<(8*j)
	}
	return i
}

// 小端法
func BytesToInt64WithMin(bs []byte) int64{
	bLen := len(bs)
	if(bLen > 8){
		bLen = 8
	}
	var i int64
	var j uint
	for j = 0; j < uint(bLen); j++{
		i |= int64(bs[j])<<(8*j)
	}
	return i
}

// 大端法
func BytesToIntWithMax(bs []byte) int{
	bLen := len(bs)
	if(bLen > 4){
		bLen = 4
	}
	var i int
	var j uint
	for j = 0; j < uint(bLen); j++{
		i |= int(bs[bLen-int(j)-1])<<(8*j)
	}
	return i
}

// 大端法
func BytesToInt64WithMax(bs []byte) int64{
	bLen := len(bs)
	if(bLen > 8){
		bLen = 8
	}
	var i int64
	var j uint
	for j = 0; j < 8; j++{
		i |= int64(bs[bLen-int(j)-1])<<(8*j)
	}
	return i
}

func Max(v1 int, v2 int)int{
	if v1 < v2{
		return v2
	}
	return v1
}

func Min(v1 int, v2 int)int{
	if v1 > v2{
		return v2
	}
	return v1
}
