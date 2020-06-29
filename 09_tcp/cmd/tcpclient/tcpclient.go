package main

import (
	"io"
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
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		log.Println(err)
		return err
	}
	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			log.Println(cerr)
			if err == nil {
				err = cerr
			}
		}
	}(conn)

	_, err = conn.Write([]byte("hello!"))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
