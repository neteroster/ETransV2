package main

import (
	"bytes"
	"io"
	"log"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		// Don't exit
		log.Println("[ERROR]", err.Error())
	}
}

func checkArgs() {
	if len(os.Args) != 2 {
		log.Fatalln("[FATAL] Usage:", os.Args[0], "<port>")
	}
}

func main() {
	checkArgs()
	inc, err := net.Listen("tcp", ":" + os.Args[1]) // Get the connection from client
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
	var receiveBytesContainer []byte = make([]byte, 10240) // Receive byte
	readbytes, err := conn.Read(receiveBytesContainer)
	if err != nil {
		checkError(err)
		conn.Close()
		return
	}
	fn := string(bytes.Split(bytes.Split(receiveBytesContainer, []byte("//etransv2-head//"))[0], []byte(";;"))[0])
	fs := string(bytes.Split(bytes.Split(receiveBytesContainer, []byte("//etransv2-head//"))[0], []byte(";;"))[1])
	log.Println("[INFO] Receiving: " + fn + "  Size: " + fs)
	ind := bytes.Index(receiveBytesContainer, []byte("//etransv2-head//"))
	realbytes := receiveBytesContainer[ind + 17:]
	f, err := os.Create(fn)
	checkError(err)
	defer f.Close()
	f.Write(realbytes[:readbytes - (ind + 17)])
	for {
		readbytes, err := conn.Read(receiveBytesContainer) // Read to receiveBytesContainer
		if err != nil && err == io.EOF {
			log.Println("[INFO]", fn+" Saved")
			break
		} else if err != nil && err != io.EOF {
			checkError(err)
			break
		}
		f.Write(receiveBytesContainer[:readbytes]) // Write to file
	}
}
