package main

import (
	"database/sql"
	"fmt"

	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/configs"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	println("connected on database", db.Ping())
}
