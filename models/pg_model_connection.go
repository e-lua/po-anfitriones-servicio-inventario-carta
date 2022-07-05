package models

import (
	"context"
	"sync"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

//var PostgresCN = Conectar_Pg_DB()

var (
	once_pg sync.Once
	p_pg    *pgxpool.Pool
)

func Conectar_Pg_DB() *pgxpool.Pool {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	once_pg.Do(func() {
		urlString := "postgres://postgresxd2:asf45vcf235sfds435GFHg435fd3h2s@postgres-master:5432/postgresxd2?pool_max_conns=120"
		config, _ := pgxpool.ParseConfig(urlString)
		p_pg, _ = pgxpool.ConnectConfig(ctx, config)
	})
	return p_pg
}
