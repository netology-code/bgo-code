package main

import (
	"context"
	"log"
	"time"
)

func main() {
	parent := context.Background() // базовый пустой контекст
	cancellable, cancel := context.WithCancel(parent) // производный контекст + функция отмены

	go func() {
		for {
			select {
			case <-cancellable.Done(): // если в канале есть сообщение, вычитываем его
				log.Println("cancel signal received")
				return // не забываем завершить горутину
			default: // если нет, то не блокируясь выполняем default
				time.Sleep(time.Second)
				log.Println("do part of work")
				// важно: без default будет блокировка
			}
		}
	}()

	time.Sleep(10 * time.Second)
	cancel() // при вызове cancel() устанавливается ошибка Canceled
	log.Println(context.Canceled == cancellable.Err())

	time.Sleep(10 * time.Second) // чтобы успеть напечатать cancel signal received
}



