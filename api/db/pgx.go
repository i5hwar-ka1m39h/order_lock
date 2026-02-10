package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)



func Nex(ctx context.Context, dburl string) *pgxpool.Pool{
	pool, err:=pgxpool.New(ctx, dburl)

	if err != nil{
		log.Fatalln("error occured while generation db pool", err)
	}

	if err := pool.Ping(ctx); err != nil{
		log.Fatalln("error occured while pinging db", err)
	}

	return pool
}