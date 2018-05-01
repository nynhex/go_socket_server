package main

import (
    "log"
	   "net"
	   "strconv"
	   "strings"
)

const (
	Message       = "Hey there!"
	StopCharacter = "\r\n\r\n"
)


func SocketClient(ip string, port int) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)

	defer conn.Close()

	if err != nil {
		log.Fatalln(err)
	}

	conn.Write([]byte(Message))
	conn.Write([]byte(StopCharacter))
	log.Printf("Send: %s", Message)

	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)
	log.Printf("Receive: %s", buffer[:n])

}

func main() {

	var (
		ip_address   = "127.0.0.1"
		tcp_port = 6666
	)

	SocketClient(ip_address, tcp_port)

}
