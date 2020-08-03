package main

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib" // при импорте пакета выполнится код, который "зарегистрирует" драйвер
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
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()

	//tag, err := db.ExecContext(ctx, `
	//	INSERT INTO cards(number, balance, issuer, holder, owner_id, status)
	//	VALUES ($1, $2, $3, $4, $5, $6)
	//`, "**** 0001", 100_00, "Visa", "PETR IVANOV", 1, "ACTIVE")
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//log.Println(tag.RowsAffected())

	//client := &Client{}
	//err = db.QueryRowContext(ctx, `
	//	SELECT id, full_name, birthday FROM clients WHERE id = $1
	//`, 1).Scan(&client.Id, &client.FullName, &client.Birthday)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//log.Println(client)

	//client := &Client{}
	//err = db.QueryRowContext(ctx, `
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
	//rows, err := db.QueryContext(ctx, `
	//	SELECT id, full_name, birthday FROM clients WHERE status = $1
	//`, "ACTIVE")
	//if err != nil {
	//	// TODO: отдельная обработка для ErrNoRows
	//	log.Println(err)
	//	return
	//}
	//defer func() {
	//	if cerr := rows.Close(); cerr != nil {
	//		log.Println(cerr)
	//	}
	//}()
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
	err = db.QueryRowContext(ctx, `
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
