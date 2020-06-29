package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {
	if err := execute(); err != nil {
		os.Exit(1)
	}
}

func execute() (err error) {
	listener, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		if cerr := listener.Close(); cerr != nil {
			log.Println(cerr)
			if err == nil {
				err = cerr
			}
		}
	}()
	for {
		conn, err := listener.Accept() // для клиентов
		if err != nil {
			log.Println(err)
			continue
		}
		handle(conn)
	}
}

func handle(conn net.Conn) {
	defer func() {
		if cerr := conn.Close(); cerr != nil {
			log.Println(cerr)
		}
	}()

	reader := bufio.NewReader(conn)
	const delim = '\n'
	line, err := reader.ReadString(delim)
	if err != nil {
		if err != io.EOF {
			log.Println(err)
		}
		log.Printf("received: %s\n", line)
		return
	}
	log.Printf("received: %s\n", line)

	username := "Василий"
	balance := "1 000.50" // FIXME: написать функцию-форматтер

	page, err := ioutil.ReadFile("web/template/index.html")
	if err != nil {
		log.Println(err)
		return
	}
	page = bytes.ReplaceAll(page, []byte("{username}"), []byte(username))
	page = bytes.ReplaceAll(page, []byte("{balance}"), []byte(balance))

	const CRLF = "\r\n"
	writer := bufio.NewWriter(conn)
	_, err = writer.WriteString("HTTP/1.1 200" + CRLF)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = writer.WriteString("Content-Type: text/html;charset=utf-8" + CRLF)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = writer.WriteString(fmt.Sprintf("Content-Length: %d", len(page)) + CRLF)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = writer.WriteString("Connection: close" + CRLF + CRLF)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = writer.Write(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = writer.Flush()
	if err != nil {
		log.Println(err)
		return
	}
}
