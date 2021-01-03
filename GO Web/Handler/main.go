package main

import (
	"net/http"
)

//自己的handler，类似于一个路由处理
type helloHandler struct{}

//实现这个ServeHTTP就创建一个自己的handler
func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home!"))
}

func main() {
	mh := helloHandler{}

	//添加一个自己的handler到DefaultServeMux的下一级，需要传入的是指针
	http.Handle("/hello", &mh)
	//将与ServeHTTP签名一致的函数使用HandlerFunc转成一个Handler
	http.Handle("/home", http.HandlerFunc(home))
	//监听端口，nil表示默认路由转发（DefaultServeMux（本身也是一个handler，可以用自己的代替））
	http.ListenAndServe("localhost:8080", nil)
}
