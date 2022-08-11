package main

import (
	"flag"
	"fmt"
	"multi-socket/common"
	"net"
	"time"
)

var (
	address = "0.0.0.0:8080"
	count   = 10
	c       = make(chan int8)
)

func init() {
	flag.StringVar(&address, "a", address, "address")
	flag.IntVar(&count, "c", count, "count")
	flag.Parse()
}

func main() {
	go func() {
		for {
			fmt.Println("connect cont: ", common.Count)
			time.Sleep(1 * time.Second)
		}
	}()

	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err)
	}

	for i := 0; i < count; i++ {
		conn, err := net.DialTCP("tcp", nil, addr)
		if err != nil {
			panic(err)
		}

		go doMain(conn)
	}

	for i := 0; i < count; i++ {
		_ = <-c
	}
}

func doMain(conn *net.TCPConn) {
	defer func() {
		common.D()
		_ = conn.Close()
		c <- 1
	}()
	common.I()

	if _, err := conn.Write(common.Data); err != nil {
		return
	}

	buf := make([]byte, 1)
	for {
		_ = conn.SetReadDeadline(time.Now().Add(3 * time.Second))
		if _, err := conn.Read(buf); err != nil {
			break
		}
		time.Sleep(1 * time.Second)
		if _, err := conn.Write(common.Data); err != nil {
			return
		}
	}
}
