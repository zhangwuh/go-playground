package main

import (
	"net"
	"os"
	"fmt"
)

func main(){
	client_()
}

func client_() {
	conn, err := net.Dial("tcp", "localhost:3540")
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}
	if err != nil {
		fmt.Printf("%v..",err)
		return
	}
	conn.Write([]byte("aaaa\n"))

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("reply from server=", string(reply))

	defer conn.Close()
}
