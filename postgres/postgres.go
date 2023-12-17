package postgres

import (
	"fmt"
	"time"
	"todo/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type client struct {
	db *sqlx.DB
}

func New() (*client, error) {
	connString := config.Value.Database.ConnectionString()
	conn, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("postgres.New: could not connect to postgres %s due to error %w", connString, err)
	}
	conn.SetMaxOpenConns(config.Value.Database.MaxPoolSize)
	conn.SetMaxIdleConns(config.Value.Database.MaxIdleConn)
	conn.SetConnMaxLifetime(time.Second * time.Duration(config.Value.Database.ConnMaxLifetimeSec))

	return &client{db: conn}, nil
}

func (c client) GetTodos() {

}

func (c client) CreateTodo(title string) {
	// query := "INSERT into todo (title) values(?) RETURNING id"
	// _ := c.db.MustExec(query, title)

}

func (c client) CompleteTodo() {

}

func (c client) DeleteTodo() {

}
