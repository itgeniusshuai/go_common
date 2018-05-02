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

func RemoveStrArr(arr []string,index int) []string{
	length := len(arr)
	if length < index{
		return arr
	}
	for i := index; i < length-1; i++{
		arr[i] = arr[i+1]
	}
	return arr[0:length-1]
}


func RemoveIntArr(arr []int,index int) []int{
	length := len(arr)
	if length < index{
		return arr
	}
	for i := index; i < length-1; i++{
		arr[i] = arr[i+1]
	}
	return arr[0:length-1]
}

