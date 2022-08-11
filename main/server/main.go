package main

import (
	"flag"
	"fmt"
	"multi-socket/common"
	"net"
	"time"
)

var (
	port = 8080
)

func init() {
	flag.IntVar(&port, "p", port, "port")
	flag.Parse()
}

func main() {
	go func() {
		for {
			fmt.Println("connect cont: ", common.Count)
			time.Sleep(1 * time.Second)
		}
	}()

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", addr)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			panic(err)
		}

		go doMain(conn)
	}
}

func doMain(conn *net.TCPConn) {
	defer func() {
		common.D()
		_ = conn.Close()
	}()
	common.I()

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
