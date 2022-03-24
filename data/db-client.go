package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/deeplink/config"
	query "github.com/deeplink/constant"
	_ "github.com/go-sql-driver/mysql"
)

type DbClient struct {
}

var Db *sql.DB

func NewDbClient() *DbClient {
	//Db, _ = connect()
	//createLinkTable(Db)
	return &DbClient{}
}

func dsn(dbName string) string {
	DB_USERNAME := config.Config("DB_USERNAME")
	DB_PASSWORD := config.Config("DB_PASSWORD")
	DB_HOST := config.Config("DB_HOST")
	DB_PORT := config.Config("DB_PORT")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, dbName)
}

func connect() (*sql.DB, error) {
	var err error
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		return nil, err
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	dbname := config.Config("DB_NAME")
	res, err := db.ExecContext(ctx, query.CREATE_DATABASE+dbname)
	if err != nil {
		return nil, err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	db.Close()
	db, err = sql.Open("mysql", dsn(dbname))
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, err
}

func createLinkTable(db *sql.DB) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query.CREATE_LINK_TABLE)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (p *DbClient) Add(query interface{}, args ...interface{}) (err error) {
	ins, err := Db.Prepare(query.(string))
	if err != nil {
		panic(err.Error())
	}
	ins.Exec(args...)

	return nil
}

func (p *DbClient) GetAll() (response interface{}, err error) {
	return response, nil
}
