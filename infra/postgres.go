package infra

import (
	"fmt"
	"github.com/code-and-chill/iskandar/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"net/url"
)

func PgConnect(db config.DBConfig, log logrus.FieldLogger) *gorm.DB {
	builder := url.URL{
		User:     url.UserPassword(db.Username, db.Password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", db.Host, db.Port),
		Path:     db.Database,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	conn, err := gorm.Open("postgres", builder.String())
	if err != nil {
		log.Fatalf(fmt.Sprintf("Unable to connect postgres with host: %s, with error\n%s", db.Host, err.Error()))
	}
	log.Info(fmt.Sprintf("Database connected successfully."))
	return conn
}

func PgDisconnect(db *gorm.DB) {
	err := db.Close()
	if err != nil {
		logrus.Fatal("Unable to close database")
	}
}
