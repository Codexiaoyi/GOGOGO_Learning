package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const Count = 5

type resData struct {
	res *http.Response
	err error
}

func doRequest() {
	//造一个客户端
	transport := http.Transport{}
	client := http.Client{
		Transport: &transport,
	}

	//response数据通道
	resChan := make(chan *resData, 1)
	defer close(resChan)
	//造一个请求对象
	res, err := http.NewRequest("POST", "http://localhost:3000/api/v1/account/login", bytes.NewBufferString("{\"email\": \"1111@qq.com\",\"password\": \"1111\"}"))
	if err != nil {
		fmt.Printf("new request failed,err:%v\n", err)
		return
	}

	//加上context
	//res = res.WithContext(ctx)

	//执行请求操作
	response, err := client.Do(res)
	//fmt.Printf("client.do response:%v,err:%v\n", response, err)
	rd := &resData{
		res: response,
		err: err,
	}
	resChan <- rd

	select {
	// case <-ctx.Done():
	// 	//接收到context对象的cancel指令后
	// 	fmt.Println("call api timeout")
	case result := <-resChan:
		//只要没结束就一直从通道内取值
		//fmt.Println("call api success")
		if result.err != nil {
			fmt.Printf("call api failed,err:%v\n", result.err)
			return
		}
		defer result.res.Body.Close()
		data, _ := ioutil.ReadAll(result.res.Body)
		fmt.Printf("response:%s\n", string(data))
	}

}

var a = 0

func main() {
	var wg sync.WaitGroup
	wg.Add(1000)
	defer wg.Wait()
	for i := 0; i < 1000; i++ {
		//新goroutine跑请求
		go func() {
			//ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10000)
			//defer cancel()
			doRequest()
			wg.Done()
		}()
	}
}
