package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DBLocal  *sql.DB
)

func ConnectDB() {
	var err error

	localDB := os.Getenv("LOCAL_DB")

	if localDB == "" {
		log.Fatal("❌ ENV database belum diset (LOCAL_DB kosong)")
	}

	DBLocal, err = sql.Open("mysql", localDB)
	if err != nil {
		log.Fatalf("❌ Gagal connect ke Jubelio: %v", err)
	}

	if err = DBLocal.Ping(); err != nil {
		log.Fatalf("❌ Ping Local gagal: %v", err)
	}

	fmt.Println("✅ Connection Successed")
}
