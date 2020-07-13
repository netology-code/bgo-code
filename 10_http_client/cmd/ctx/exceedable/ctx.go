package main

import (
	"context"
	"log"
	"time"
)

func main() {
	parent := context.Background() // базовый пустой контекст
	exceedable, _ := context.WithTimeout(parent, 5 * time.Second) // производный контекст
	// функция отмены нам не нужна, таймаут автоматически отменит контекст через 5 секунд
	// WithDeadline работает "аналогично" с точность до различий Deadline и Timeout

	go func() {
		for {
			select {
			case <-exceedable.Done(): // если в канале есть сообщение, вычитываем его
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
	log.Println(context.DeadlineExceeded == exceedable.Err())
}



