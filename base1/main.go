package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Println("开启服务器：http://127.0.0.1:9999")
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	// ResponseWriter为返回的内容
	requestLogHandler(req)
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	requestLogHandler(req)
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

func requestLogHandler(req *http.Request) {
	addLog := fmt.Sprintf("访问的地址为：%s", req.RequestURI)
	fmt.Println(addLog)
}


