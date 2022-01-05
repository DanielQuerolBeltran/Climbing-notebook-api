package bootstrap

import (
	"database/sql"
	//"fmt"
	"log"

	"github.com/DanielQuerolBeltran/Climbing-notebook-api/internal/platform/server"
	repo "github.com/DanielQuerolBeltran/Climbing-notebook-api/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host = "localhost"
	port = 1323

	dbUser = "tester"
	dbPass = "secret"
	dbHost = "db"
	dbPort = "3306"
	dbName = "test"
)

func Run() error {
	/*mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	log.Println("MYSQL URI", mysqlURI)

	db, err := sql.Open("mysql", mysqlURI)*/
	db, err := sql.Open("mysql", "tester:secret@tcp(test_db:3306)/test")

	if err != nil {
		return err
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println("ERROR DATABASE: ", err)
	}

	climbRepository := repo.NewClimbRepository(db)

	srv := server.New(host, port, climbRepository)
	return srv.Run()
}
