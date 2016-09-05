package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln("[FATAL]", err) // will exit when Error happened
	}
}

func checkArgs() {
	if len(os.Args) != 3 {
		log.Fatalln("[FATAL] Usage:", os.Args[0], "<server-ip:port>", "<filepath>")
	}
}

func sendFile(c net.Conn, filepth string) {
	f, err := os.Open(filepth)
	checkError(err)
	defer f.Close()
	fstat, err := f.Stat() // get stat
	checkError(err)
	filename := fstat.Name() //get name
	filesize := fstat.Size() //get size
	writeString := filename + ";;" + strconv.FormatInt(filesize, 10) + "//etransv2-head//"
	c.Write([]byte(writeString))
	buf := make([]byte, 10240)
	bfRd := bufio.NewReader(f)
	for {
		n, err := bfRd.Read(buf)
		if err != nil {
			break // break when finish reading
		}
		c.Write(buf[:n])
	}
}

func main() {
	checkArgs()
	conn, err := net.Dial("tcp", os.Args[1]) //connect to server
	checkError(err)
	log.Println("[INFO] Successfully connected to server")
	defer conn.Close() // finally close the connection
	sendFile(conn, os.Args[2])
	log.Println("[INFO] Successfully send file")
}
