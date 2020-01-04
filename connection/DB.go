package connection

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"testing/config"
	_ "github.com/lib/pq"
)

func Connection() *sqlx.DB {
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable timezone='Asia/Tehran'",
		config.PSConfig.User,
		config.PSConfig.Pass,
		config.PSConfig.Host,
		config.PSConfig.Port,
		config.PSConfig.Name))

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return db
}
