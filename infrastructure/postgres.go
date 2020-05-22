package infrastructure

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"net/url"
)

type DbConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func Connect(db DbConfig, log logrus.FieldLogger) *gorm.DB {
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

func DisConnect(db *gorm.DB) {
	err := db.Close()
	if err != nil {
		logrus.Fatal("Unable to close database")
	}
}
