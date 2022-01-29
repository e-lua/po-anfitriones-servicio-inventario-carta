package models

import (
	"context"
	"sync"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var PostgresCNExternal = Conectar_Pg_DB_External()

var (
	once_pg_external sync.Once
	p_pg_external    *pgxpool.Pool
)

func Conectar_Pg_DB_External() *pgxpool.Pool {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	once_pg_external.Do(func() {
		urlString_external := "postgres://postgresx4y:asd34Fg2DDFfd3saF3Fgge65sGGS45@http://c-carta.restoner-api.fun:7000/postgresx4y?pool_max_conns=50"
		config_external, _ := pgxpool.ParseConfig(urlString_external)
		p_pg_external, _ = pgxpool.ConnectConfig(ctx, config_external)
	})
	return p_pg_external
}
