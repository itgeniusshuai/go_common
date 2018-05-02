package common

import (
	"errors"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func Get(url string, params map[string]string, headers map[string]string) (string,error){
	if "" == url {
		return "",errors.New("Url can not be nil or empty string")
	}
	if params != nil && len(params) > 0{
		var paramArr []string
		for k,v := range params{
			paramArr = append(paramArr, k + "=" +v)
		}
		url = url + "?" + strings.Join(paramArr,"&")
	}
	client := &http.Client{}
	request,err := http.NewRequest("GET",url,nil)
	if err != nil{
		fmt.Println("Create get request failed err is ",err.Error() )
		return "",err
	}
	if headers != nil && len(headers) > 0{
		for k,v := range headers{
			request.Header.Set(k,v)
		}
	}
	response,err := client.Do(request)
	if err != nil{
		fmt.Println("Do get request failed err is ",err.Error() )
		return "",err
	}
	if response.StatusCode == 200{
		body,err := ioutil.ReadAll(response.Body)
		if err != nil{
			fmt.Println("Create post request failed err is ",err.Error() )
			return "",err
		}
		bodyStr := string(body)
		return bodyStr,nil
	}else{
		return "",errors.New("Failed status code is "+string(response.StatusCode))
	}
}

func Post(url string, params map[string]string, headers map[string]string) (string,error){
	if "" == url {
		return "",errors.New("Url can not be nil or empty string")
	}
	var requestBody string
	if params != nil && len(params) > 0{
		paramsBytes,_ := json.Marshal(params)
		requestBody = string(paramsBytes)
	}
	client := &http.Client{}
	request,err := http.NewRequest("POST",url,strings.NewReader(requestBody))
	if err != nil{
		fmt.Println("Create post request failed err is ",err.Error() )
		return "",err
	}
	if headers != nil && len(headers) > 0{
		for k,v := range headers{
			request.Header.Set(k,v)
		}
	}
	response,err := client.Do(request)
	if err != nil{
		fmt.Println("Do post request failed err is ",err.Error() )
		return "",err
	}
	if response.StatusCode == 200{
		body,err := ioutil.ReadAll(response.Body)
		if err != nil{
			fmt.Println("Read response body failed err is ",err.Error() )
			return "",err
		}
		bodyStr := string(body)
		return bodyStr,nil
	}else{
		return "",errors.New("Failed status code is "+string(response.StatusCode))
	}
}
