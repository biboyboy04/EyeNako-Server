package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/biboyboy04/EyeNako-Server/cmd/api"
	"github.com/biboyboy04/EyeNako-Server/config"
	"github.com/biboyboy04/EyeNako-Server/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Hello Go")

	db, err := db.NewMySQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Check if db is working/connected
	initStorage(db)

	//refactor, change to env for server addrs
	addr := "localhost:5555"
	server := api.NewAPIServer(addr, db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping() 
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB Sucessfully Connected")
} 