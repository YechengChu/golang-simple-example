package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	startTime := time.Now()
	// make函数创建了一个传递string类型参数的channel
	myChannel := make(chan string)
	for _, url := range os.Args[1:] {
		// 创建一个goroutine，并且让函数在这个goroutine异步执行http.Get方法
		go fetch(url, myChannel) // start a goroutine
	} // for
	for range os.Args[1:] {
		fmt.Println(<-myChannel) // receive from channel myChannel
	} // for
	fmt.Printf("%.2fs elapsed\n", time.Since(startTime).Seconds())
} // main

func fetch(url string, myChannel chan<- string) {
	startTime := time.Now()
	response, errorMsg := http.Get(url)
	if errorMsg != nil {
		myChannel <- fmt.Sprint(errorMsg) // send to channel myChannel
	} // if
	// io.Copy会把响应的Body内容拷贝到ioutil.Discard输出流中
	nbytes, errMsg := io.Copy(ioutil.Discard, response.Body)
	response.Body.Close() // don't leak resources
	if errMsg != nil {
		myChannel <- fmt.Sprintf("Erro while loding %s: %v", url, errMsg)
		return
	} // if
	timeSpent := time.Since(startTime).Seconds()
	// 当请求返回内容时，fetch函数都会往myChannel这个channel里写入一个字符串
	myChannel <- fmt.Sprintf("%.2fs  %7d  %s", timeSpent, nbytes, url)
} // fetch
