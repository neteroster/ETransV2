package main

import (
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func checkError(err error) {
	if err != nil {
		// Don't exit
		log.Println("[ERROR]", err.Error())
	}
}

func checkArgs() {
	if len(os.Args) != 2 {
		log.Fatalln("[FATAL]Usage:", os.Args[0], "<port>")
	}
}

func main() {
	checkArgs()
	inc, err := net.Listen("tcp", ":"+os.Args[1]) // Get the connection from client
	checkError(err)
	for {
		conn, err := inc.Accept() // Accept client
		checkError(err)
		go handleClient(conn) // To handle connection
	}
}

func handleClient(conn net.Conn) {
	log.Println("[INFO]", conn.RemoteAddr(), "Connected.")
	defer conn.Close()
	var dt []byte = make([]byte, 1024) // Receive byte
	timenow := strconv.FormatInt(time.Now().UnixNano(), 10)
	f, err := os.Create(timenow) // filename is time
	defer f.Close()
	for {
		_, err = conn.Read(dt) // Read to dt
		if err != nil && err == io.EOF {
			log.Println("[INFO]", timenow, "Saved")
			break
		} else if err != nil && err != io.EOF {
			checkError(err)
			break
		}
		f.Write(dt) //write to file
	}
}
