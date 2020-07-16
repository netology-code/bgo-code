package main

import "log"

func endpoint() {
	panic("implement me")
}

func middleware() {
	endpoint()
}

func catch() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	middleware()
}

func main() {
	catch()
	log.Println("ok!")
}



