package common


// 指定索引增加元素,索引大于范围添加到最后，索引小于0放到前面
func AddIntArr(arr []int, index int,e int) ([]int){
	if(len(arr) <= index){
		return AddIntArrAfter(arr,e)
	}
	if(index < 0){
		return AddIntArrFront(arr,e)
	}
	length := len(arr) + 1
	retArr := make([]int, length)
	for i := 0; i < index; i++{
		retArr[i] = arr[i]
	}
	retArr[index] = e
	for i := index; i < length - 1; i++{
		retArr[i + 1] = arr[i]
	}
	return retArr
}

// 添加到后面
func AddIntArrAfter(arr []int,e int) []int{
	if arr == nil{
		arr = make([]int,1)
		arr[0] = e
		return arr
	}
	arr = append(arr, e)
	return arr
}

// 添加到前面
func AddIntArrFront(arr []int,e int) []int{
	if arr == nil{
		arr = make([]int,1)
		arr[0] = e
		return arr
	}
	retArr := make([]int,0,len(arr)+1)
	retArr = append(retArr, e)
	retArr = append(retArr,arr...)
	return retArr
}

// 指定索引增加元素,索引大于范围添加到最后，索引小于0放到前面
func AddStrArr(arr []string, index int,e string) ([]string){
	if(len(arr) <= index){
		return AddStrArrAfter(arr,e)
	}
	if(index < 0){
		return AddStrArrFront(arr,e)
	}
	length := len(arr) + 1
	retArr := make([]string, length)
	for i := 0; i < index; i++{
		retArr[i] = arr[i]
	}
	retArr[index] = e
	for i := index; i < length - 1; i++{
		retArr[i + 1] = arr[i]
	}
	return retArr
}

// 添加到后面
func AddStrArrAfter(arr []string,e string) []string{
	if arr == nil{
		arr = make([]string,1)
		arr[0] = e
		return arr
	}
	arr = append(arr, e)
	return arr
}

// 添加到前面
func AddStrArrFront(arr []string,e string) []string{
	if arr == nil{
		arr = make([]string,1)
		arr[0] = e
		return arr
	}
	retArr := make([]string,0,len(arr)+1)
	retArr = append(retArr, e)
	retArr = append(retArr,arr...)
	return retArr
}

// 根据索引移除
func RemoveStrArr(arr []string,index int) []string{
	length := len(arr)
	if length < index{
		return arr
	}
	sArr := CopyStrArr(arr,0,index)
	eArr := CopyStrArr(arr,index+1,length)
	return append(sArr, eArr...)
}

// 根据索引移除
func RemoveIntArr(arr []int,index int) []int{
	length := len(arr)
	if length < index{
		return arr
	}
	sArr := CopyIntArr(arr,0,index)
	eArr := CopyIntArr(arr,index+1,length)
	return append(sArr, eArr...)
}

// 范围复制
func CopyIntArr(arr []int, start int, end int) []int{
	length := len(arr)
	if start >= end || start > length || end <= 0{
		return nil
	}
	if start <0 {
		start = 0
	}
	if end > length{
		end = length
	}
	sliceLength := end - start
	retArr := make([]int,sliceLength)
	for i := 0; i < sliceLength; i++{
		retArr[i] = arr[start+i]
	}
	return retArr
}

// 范围复制
func CopyStrArr(arr []string, start int, end int) []string{
	length := len(arr)
	if start >= end || start > length || end <= 0{
		return nil
	}
	if start <0 {
		start = 0
	}
	if end > length {
		end = length
	}
	sliceLength := end - start
	retArr := make([]string,sliceLength)
	for i := 0; i < sliceLength; i++{
		retArr[i] = arr[start+i]
	}
	return retArr
}

// 某个元素的索引，从起始位置开始搜寻
func IndexOfIntArr(arr []int, v int) int{
	return IndexOfIntArrLeft(arr,v,0)
}

// 某个元素的索引，从指定索引开始
func IndexOfIntArrLeft(arr []int, v int, start int) int{
	if start < 0 {
		start = 0
	}
	length := len(arr)
	if start >= length{
		return -1
	}
	for i := start; i < len(arr); i++{
		if arr[i] == v{
			return i
		}
	}
	return -1
}

// 某个元素索引，从结尾开始搜寻
func LastIndexOfIntArr(arr []int, v int) int{
	length := len(arr)
	for i := length-1; i >= 0; i--{
		if arr[i] == v{
			return i
		}
	}
	return -1
}

// 某个元素索引，从起始开始搜寻
func IndexOfStrArr(arr []string, v string) int{
	return IndexOfStrArrLeft(arr,v,0)
}

// 某个元素的索引，从指定索引开始
func IndexOfStrArrLeft(arr []string, v string, start int) int{
	if start < 0 {
		start = 0
	}
	length := len(arr)
	if start >= length{
		return -1
	}
	for i := start; i < len(arr); i++{
		if arr[i] == v{
			return i
		}
	}
	return -1
}

// 某个元素索引，从结尾开始搜寻
func LastIndexOfStrArr(arr []string, v string) int{
	length := len(arr)
	for i := length-1; i >= 0; i--{
		if arr[i] == v{
			return i
		}
	}
	return -1
}

// 移除某个元素，元素为从左侧开始第一个
func RemoveIntByValueArr(arr []int, v int) []int{
	index := IndexOfIntArr(arr,v)
	if index == -1{
		return arr
	}
	return RemoveIntArr(arr,index)
}

// 移除某个元素，元素为从左侧开始第一个
func RemoveStrByValueArr(arr []string, v string) []string{
	index := IndexOfStrArr(arr,v)
	if index == -1{
		return arr
	}
	return RemoveStrArr(arr,index)
}

// 是否包含元素
func IsContainIntArr(arr []int, v int) bool{
	return (IndexOfIntArr(arr,v) != -1)
}

// 是否包含元素
func IsContainStrArr(arr []string, v string) bool{
	return (IndexOfStrArr(arr,v) != -1)
}

// 是否为空
func IsEmptyIntArr(arr []int) bool{
	return arr == nil || len(arr) == 0
}

// 是否为空
func IsEmptyStrArr(arr []string) bool{
	return arr == nil || len(arr) == 0
}

func AddAllIntArrSlice(desc []int, src []int,start int, end int) []int{
	srcLength := len(src)
	if start < 0{
		start = 0
	}
	if end < start{
		return desc
	}
	if end > srcLength{
		end = srcLength
	}
	return append(desc,src[start:end]...)
}

func AddAllIntArr(desc []int, src []int ) []int{
	return append(desc,src...)
}

func AddAllStrArrSlice(desc []string, src []string,start int, end int) []string{
	srcLength := len(src)
	if start < 0{
		start = 0
	}
	if end < start{
		return desc
	}
	if end > srcLength{
		end = srcLength
	}
	return append(desc,src[start:end]...)
}

func AddAllStrArr(desc []string, src []string ) []string{
	return append(desc,src...)
}

