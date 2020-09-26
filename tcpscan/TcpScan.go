package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	hostname := flag.String("hostname", "www.baidu.com", "hostname to test")
	portStart := flag.Int("start-port", 80, "the port on which the scanning starts")
	portEnd := flag.Int("end-port", 1000, "the port on which the scanning ends")
	timeout := flag.Duration("timeout", time.Millisecond * 200, "timeout")
	flag.Parse()

	//availPorts := make([]int, *portEnd - *portStart)
	availPorts := []int{}

	fmt.Printf("hostname=%s, portStart=%d, portEnd=%d \n", *hostname, *portStart, *portEnd)

	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	for port := *portStart; port <= *portEnd; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(*hostname, p, *timeout)
			if opened {
				mutex.Lock()
				availPorts = append(availPorts, p)
				mutex.Unlock()
			}
			wg.Done()
		}(port)
	}

	wg.Wait()
	fmt.Printf("Open ports: %v\n", availPorts)
}

func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		_= conn.Close()
		return true
	}

	return false
}
