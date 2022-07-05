package models

import (
	"context"
	"sync"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var PostgresCN_Slave = Conectar_Pg_DB_Slave()

var (
	once_pg_Slave sync.Once
	p_pg_Slave    *pgxpool.Pool
)

func Conectar_Pg_DB_Slave() *pgxpool.Pool {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	once_pg_Slave.Do(func() {
		urlString := "postgres://postgresxd2:asf45vcf235sfds435GFHg435fd3h2s@postgresql-slave:5432/postgresxd2?pool_max_conns=120"
		config, _ := pgxpool.ParseConfig(urlString)
		p_pg_Slave, _ = pgxpool.ConnectConfig(ctx, config)
	})
	return p_pg_Slave
}
