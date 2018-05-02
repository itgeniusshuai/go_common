package common

import (
	"time"
	"fmt"
)

const maxNumInFixTime int = 5
const maxNumFixTime int64 = 5

const maxIntervalFixNum int = 5
const maxInterval int64 = 1

var maxNumArr = make(map[string][]int64,0)
var maxIntervalArr = make(map[string][]int64,0)


func MaxNumInFixedTime(ip string){
	arr := maxNumArr[ip]
	if arr == nil{
		arr = make([]int64,0)
	}
	ctime := time.Now().Unix()
	arr = append(arr, ctime)
	fmt.Println("arr ",arr)
	if len(arr)>=maxNumInFixTime{
		interval :=  ctime - arr[0]
		fmt.Println("interval : ",interval)
		if interval < maxNumFixTime{
			fmt.Println(ip)
		}
		arr = arr[1:]
	}
	maxNumArr[ip] = arr
}

func MaxIntervalFixNum(ip string){
	arr := maxIntervalArr[ip]
	ctime := time.Now().Unix()
	if arr == nil ||  ctime - arr[len(arr)-1] > maxInterval{
		arr = make([]int64,0)
	}else{
		if len(arr) >= maxIntervalFixNum-1{
			arr = arr[:maxIntervalFixNum-1]
			fmt.Println("ip : ",ip)
		}
	}
	arr = append(arr, ctime)
	maxIntervalArr[ip] = arr
}
