package common

func IntersectionNumArr(maxValue uint, NumArr *[]uint, NumArr2 ...*[]uint)*[]uint{
	if NumArr == nil{
		return nil
	}
	if NumArr2 == nil || len(NumArr2) == 0{
		return NumArr
	}
	baseBitSet := NumArrToBitSetArr(NumArr,maxValue)
	for _,e := range NumArr2{
		baseBitSet = IntersectionBitSet(NumArrToBitSetArr(e,maxValue),baseBitSet)
	}
	return BitSetArrToNumArr(baseBitSet)
}

func IntersectionBitSet( bitSetArr *[]int64, bitSetArr2 *[]int64)*[]int64{
	if len(*bitSetArr) > len(*bitSetArr2){
		tmpSetArr := bitSetArr
		bitSetArr = bitSetArr2
		bitSetArr2 = tmpSetArr
	}
	for i,e := range *bitSetArr {
		(*bitSetArr2)[i] = ((*bitSetArr2)[i] & e)
	}
	return bitSetArr2
}

// 将数字数组转为bitset数组
func NumArrToBitSetArr(arr *[]uint,maxValue uint)*[]int64{
	bitSetArr := make([]int64,(maxValue + 63)>>6)
	for _,v := range *arr{
		// 所在数组中的元素位置
		var eIndex uint = v >> 6
		// 元素对应的位置
		var bitIndex uint = v & 63
		var base int64 = 1
		bitSetArr[eIndex] = bitSetArr[eIndex] | (base << bitIndex)
	}
	return &bitSetArr
}

// 将bitset转为对应的数字的数组
func BitSetArrToNumArr(arr *[]int64) *[]uint{
	numArr := make([]uint,0)
	for i,num := range *arr{
		var base int64 = 1
		var j uint = 0
		for ; j < 64; j++{
			if ((base << j)&num) != 0{
				numArr = append(numArr, uint(i << 6) + j)
			}
		}
	}
	return &numArr
}
