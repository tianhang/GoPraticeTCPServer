package main

import (
	"fmt"
	"net"
	"os"
)

type myConn struct {
	conn net.Conn
	prefix string
}
func handleError(err error){
	if err != nil {
		fmt.Println("Error on read: ", err)
		os.Exit(-1)
	}
}
func handleConn(mconn *myConn){
	//fmt.Println("reading once from conn ...")
	var buf [1024]byte

	n,err := mconn.conn.Read(buf[:])
	handleError(err)

	fmt.Println(mconn.prefix, ":", string(buf[0:n]))
	mconn.conn.Close()
}

func main() {
	println("server start ...")
	ln,err := net.Listen("tcp",":15440")
	handleError(err)
	var maxRoutineNum = make(chan int ,5)
	connNum := 0;
	for{

		//fmt.Println("Waiting for a connection via Accept")
		conn,err := ln.Accept()
		handleError(err)

		mconn := &myConn{
			conn:conn,
			prefix: fmt.Sprintf("%d says", connNum),
		}
		maxRoutineNum <- 1
		go handleConn(mconn)
		<- maxRoutineNum
		connNum++
	}
	fmt.Println("Exiting")
}
