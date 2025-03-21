package main

import (
	"log"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/ndraoo/restapi-go/cmd/api"
	"github.com/ndraoo/restapi-go/db"
	"github.com/ndraoo/restapi-go/config"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPass, 
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	initStorage(db)

	server := api.NewAPIserver(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatalf("failed to setup server: %v", err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {	
		log.Fatalf("failed to connect to db: %v", err)
	}
	log.Println("connected to db")

}