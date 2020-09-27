package memotest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

var HttpGetBody = httpGetBody

func incomingUrls() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			//"https://golang.org",
			"https://blog.csdn.net/lsh6688/article/details/9671349",
			"https://blog.csdn.net/defonds/article/details/17953849",
			"https://www.infoq.cn/article/yPB3Y2lv-DsFtRr5Cguv",
			"https://www.baidu.com",
			"https://www.ibm.com/developerworks/cn/devops/1609_liuhh_finance1/index.html",
			"https://developer.ibm.com/zh/technologies/devops/",
			"https://www.ibm.com/developerworks/cn/devops/1609_wangz_finance4/index.html",
			"https://www.ibm.com/developerworks/cn/devops/1609_baiyf_finance5/index.html",
			//"https://godoc.org",
			//"https://play.golang.org",
			//"http://gopl.io",
		} {
			ch <- url
		}
	}()

	return ch
}

type M interface {
	Get(key string) (interface{}, error)
}

func Sequential(t *testing.T, m M) {
	for url := range incomingUrls() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}

		fmt.Printf("%s, %s ,%d bytes \n", url, time.Since(start), len(value.([]byte)))
	}
}


func Concurrent(t *testing.T, m M) {
	var n sync.WaitGroup
	for url := range incomingUrls() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %s bytes\n", url, time.Since(start), len(value.([]byte)))
		}(url)
	}

	println("Begin to wait")
	n.Wait()
}














