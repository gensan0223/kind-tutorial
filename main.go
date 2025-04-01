package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres-service"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "mydb"
)

func getDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func fetchData(w http.ResponseWriter, r *http.Request) {
	db, err := getDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// サンプルクエリ：テーブルから情報を取得
	var id int
	var name string
	err = db.QueryRow("SELECT id, name FROM user LIMIT 1").Scan(&id, &name)
	if err != nil {
		log.Fatal(err)
	}

	// 取得した情報をログに出力
	log.Printf("Fetched Data: id=%d, name=%s", id, name)

	// レスポンスを返す
	fmt.Fprintf(w, "Fetched Data: id=%d, name=%s", id, name)
}

func main() {
	http.HandleFunc("/fetch", fetchData)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

