package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
)

func main() {
        if len(os.Args) != 2 {
                fmt.Println("Usage: ", os.Args[0], "host")
                os.Exit(1)
        }
        host := os.Args[1]
        conn, err := net.Dial("tcp", host+":8080")
        checkError(err)
        for {
                		reader := bufio.NewReader(os.Stdin)     
						fmt.Print("Text to send: ")     
						text, _ := reader.ReadString('\n')     // send to
						fmt.Fprintf(conn, text + "\n")     // listen for
						message, _ :=
						bufio.NewReader(conn).ReadString('\n')     
						fmt.Print("Message from server: "+message)
        }
}
func checkError(err error) {
        if err != nil {
                fmt.Println("Fatal error ", err.Error())
        }
}