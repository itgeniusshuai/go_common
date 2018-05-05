package common

// 小端法
func BytesToIntWithMin(bs []byte) int{
	bLen := len(bs)
	if(bLen < 4){
		for m := 0; m < 4-bLen; m ++{
			bs = append(bs, 0)
		}
	}
	var i int
	var j uint
	for j = 0; j < 4; j++{
		i |= int(bs[j])<<(8*j)
	}
	return i
}

// 小端法
func BytesToInt64WithMin(bs []byte) int64{
	bLen := len(bs)
	if(bLen < 8){
		for m := 0; m < 8-bLen; m ++{
			bs = append(bs, 0)
		}
	}
	var i int64
	var j uint
	for j = 0; j < 8; j++{
		i |= int64(bs[j])<<(8*j)
	}
	return i
}

// 大端法
func BytesToIntWithMax(bs []byte) int{
	var i int
	var j uint
	for j = 0; j < 4; j++{
		i |= int(bs[3-j])<<(8*j)
	}
	return i
}

// 大端法
func BytesToInt64WithMax(bs []byte) int64{
	var i int64
	var j uint
	for j = 0; j < 8; j++{
		i |= int64(bs[7-j])<<(8*j)
	}
	return i
}