package main

import (
	"context"
	"log"
	"time"
)

func main() {
	parent := context.Background() // базовый пустой контекст
	cancelable, cancel := context.WithCancel(parent)
	exceedable, _ := context.WithTimeout(cancelable, 5 * time.Second)

	go func() {
		for {
			select {
			case <-exceedable.Done():
				log.Println("cancel signal received")
				return
			default:
				time.Sleep(time.Second)
				log.Println("do part of work")
			}
		}
	}()
	time.Sleep(3 * time.Second)
	// если на родительском вызвать cancel, то
	// все вложенные контексты тоже отменятся
	// но наоборот
	cancel()
	time.Sleep(3 * time.Second)
}
