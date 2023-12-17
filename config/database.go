package config

import "fmt"

type Database struct {
	Host               string
	Port               int
	Name               string
	Username           string
	Password           string
	MaxPoolSize        int
	MaxIdleConn        int
	ConnMaxLifetimeSec int
}

func (db Database) ConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s port=%d sslmode=disable", db.Name, db.Username, db.Password, db.Host, db.Port)
}
