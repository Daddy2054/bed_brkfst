package dbrepo

import (
	"bed_brkfst/internal/config"
	"bed_brkfst/internal/repository"
	"database/sql"
)
type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// type testDBRepo struct {
// 	App *config.AppConfig
// 	DB *sql.DB
// }

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

// func NewTestingsRepo(a *config.AppConfig) repository.DatabaseRepo {
// 	return &testDBRepo{
// 		App: a,
// 	}
// }
