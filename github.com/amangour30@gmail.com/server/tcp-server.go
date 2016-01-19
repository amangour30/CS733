package main

import (
        "fmt"
        "net"
		"strings"
		"bufio"
)

func main() {
        service := "127.0.0.1:8080"
        tcpAddr, err := net.ResolveTCPAddr("tcp", service)
        checkError(err)
        listener, err := net.ListenTCP("tcp", tcpAddr)
        checkError(err)
        for {
                conn, err := listener.Accept()
                //fmt.Println("Server listerning")
                _, err = conn.Read([]byte("HEAD"))
				message, _ := bufio.NewReader(conn).ReadString('\n')
				// output message received     
				fmt.Print("Message Received:", string(message))     
				// sample process for string received     
				newmessage := strings.ToUpper(message)     
				// send new string back to client     
				conn.Write([]byte(newmessage + "\n")) 
                if err != nil {
                        conn.Close()
                }
                if err != nil {
                        continue
                }
        }
}

func checkError(err error) {
        if err != nil {
                fmt.Println("Fatal error ", err.Error())
        }
}