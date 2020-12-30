package main

import (
	"fmt"
	"net"
	"sort"
)

//工人函数
func worker(ports chan int, results chan int) {
	for port := range ports {
		address := fmt.Sprintf("47.106.139.187:%d", port)
		conn, err := net.Dial("tcp", address)

		//如果错误信息不为空，则说明端口没读到
		if err != nil {
			results <- 0
			fmt.Printf("%d 端口关着呢...\n", port)
			continue
		}
		results <- port
		conn.Close()
	}
}

//异步扫描Tcp端口
func asyncTest() {
	ports := make(chan int, 1000)
	results := make(chan int)
	var openedPorts []int

	//创建1000个工人
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		//循环的端口1~65535
		for i := 1; i < 8000; i++ {
			//往管道塞值，工人救会自己取
			ports <- i
		}
	}()

	for i := 1; i < 8000; i++ {
		//取出结果
		port := <-results
		if port != 0 {
			openedPorts = append(openedPorts, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openedPorts)
	fmt.Printf("-----------------%d 个端口打开--------------", len(openedPorts))
	for _, p := range openedPorts {
		fmt.Printf("%d 端口是打开的！\n", p)
	}
}
