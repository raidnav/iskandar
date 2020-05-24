package infra

import (
	"crypto/tls"
	"fmt"
	"github.com/code-and-chill/iskandar/config"
	"github.com/sirupsen/logrus"
	"net"
	"time"

	"gopkg.in/mgo.v2"
)

func CosmosConnect(cfg config.DBConfig, log logrus.FieldLogger) *mgo.Session {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{fmt.Sprintf("%s.%s:%d", cfg.Database, cfg.Host, cfg.Port)},
		Timeout:  60 * time.Second,
		Database: cfg.Database,
		Username: cfg.Username,
		Password: cfg.Password,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalf("Unable to connect to mongodb.", err.Error())
	}
	return session
}

func CosmosDisconnect(session *mgo.Session) {
	session.Close()
}
