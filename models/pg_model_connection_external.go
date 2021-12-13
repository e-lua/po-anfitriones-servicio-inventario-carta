package models

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

var PostgresCNExternal = Conectar_Pg_DB_External()

var (
	once_pg_external sync.Once
	p_pg_external    *pgxpool.Pool
)

func Conectar_Pg_DB_External() *pgxpool.Pool {

	once_pg_external.Do(func() {
		urlString_external := "postgres://postgresx4y:postgresx4y@67.205.146.218:7000/postgresx4y?pool_max_conns=50"
		config_external, _ := pgxpool.ParseConfig(urlString_external)
		p_pg_external, _ = pgxpool.ConnectConfig(context.Background(), config_external)
	})
	return p_pg_external
}
