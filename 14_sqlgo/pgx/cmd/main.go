package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

type Client struct {
	Id       int64
	FullName string
	Birthday time.Time
}

func main() {
	dsn := "postgres://app:pass@localhost:5432/db"
	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		log.Println(err)
		return
	}
	defer pool.Close()

	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Release()

	//tag, err := conn.Exec(ctx, `
	//	INSERT INTO card(number, balance, issuer, holder, owner_id, status)
	//	VALUES ($1, $2, $3, $4, $5, $6)
	//`, "**** 0001", 100_00, "Visa", "PETR IVANOV", 1, "ACTIVE")
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//log.Println(tag.RowsAffected())

	//client := &Client{}
	//err = conn.QueryRow(ctx, `
	//	SELECT id, full_name, birthday FROM clients WHERE id = $1
	//`, 1).Scan(&client.Id, &client.FullName, &client.Birthday)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//log.Println(client)

	//client := &Client{}
	//err = conn.QueryRow(ctx, `
	//	SELECT id, full_name, birthday FROM clients WHERE id = $1
	//`, 1).Scan(&client.Id, &client.FullName, &client.Birthday)
	//if err != nil {
	//	if err != pgx.ErrNoRows {
	//		log.Println(err)
	//		return
	//	}
	//}
	//log.Println(client)

	//clients := make([]*Client, 0)
	//rows, err := conn.Query(ctx, `
	//	SELECT id, full_name, birthday FROM clients WHERE status = $1
	//`, "ACTIVE")
	//if err != nil {
	//	// TODO: отдельная обработка для ErrNoRows
	//	log.Println(err)
	//	return
	//}
	//defer rows.Close()
	//for rows.Next() {
	//	client := &Client{}
	//	err = rows.Scan(&client.Id, &client.FullName, &client.Birthday)
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//	clients = append(clients, client)
	//}
	//err = rows.Err()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	var id int64
	err = conn.QueryRow(ctx, `
		INSERT INTO cards(number, balance, issuer, holder, owner_id, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`, "**** 0001", 100_00, "Visa", "PETR IVANOV", 1, "ACTIVE").Scan(&id)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(id)
}
