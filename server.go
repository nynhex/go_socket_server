package main

import (
    "bufio"
	  "io"
	  "log"
	  "net"
	  "os"
	  "strconv"
	  "strings"
)

const Message = "Yes? How Can I Help You"

func SocketServer(port int) {

	listen, error := net.Listen("tcp4", ":"+strconv.Itoa(port))
	defer listen.Close()
	if error != nil {
		log.Fatalf("Socket listen port %d failed,%s", port, error)
		os.Exit(1)
	}
	log.Printf("Begin listen port: %d", port)

	for {
		conn, error := listen.Accept()
		if error != nil {
			log.Fatalln(error)
			continue
		}
		go ConnectionHandler(conn)
	}

}

func ConnectionHandler(conn net.Conn) {

	defer conn.Close()

	var (
		buffer = make([]byte, 1024)
		read   = bufio.NewReader(conn)
		write   = bufio.NewWriter(conn)
	)

ILOOP:
	for {
		n, err := read.Read(buffer)
		data := string(buffer[:n])

		switch err {
		case io.EOF:
			break ILOOP
		case nil:
			log.Println("Receive:", data)
			if isTransportOver(data) {
				break ILOOP
			}

		default:
			log.Fatalf("Receive data failed:%s", err)
			return
		}

	}
	write.Write([]byte(Message))
	write.Flush()
	log.Printf("Send: %s", Message)

}

func isTransportOver(data string) (over bool) {
	over = strings.HasSuffix(data, "\r\n\r\n")
	return
}

func main() {

	port := 6666

	SocketServer(port)
}
