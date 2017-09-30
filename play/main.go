package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	fmt.Println("start...tasks")
	var wg sync.WaitGroup

	wg.Add(1)
	go Get(&wg, "http://ip.cn/")
	wg.Add(1)
	go Get(&wg, "http://httpbin.org/ip")

	if waitTimeout(&wg, 5.9e8) {
		fmt.Println("timeout")
	}
}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}

func Get(wg *sync.WaitGroup, addr string) {
	req, _ := http.NewRequest("GET", addr, nil)
	req.Header.Add("User-Agent", "HTTPie/0.9.9")
	client := new(http.Client)
	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	wg.Done()
}
