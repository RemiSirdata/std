package main

import (
	"gopkg.in/mgo.v2"
	"github.com/inconshreveable/log15"
	"fmt"
)

// use savethedate
// db.createCollection("guest")
// db.guest.createIndex({"login":1}, {unique:true})

type Mongo struct {
	Address   string
	Session   *mgo.Session
	Logger    log15.Logger
	Connected bool
}

func (m *Mongo) Connect() error {
	m.Logger.Debug(fmt.Sprintf("Trying to connect to %s", m.Address))
	session, err := mgo.Dial(m.Address)
	if err != nil {
		m.Logger.Crit(fmt.Sprintf("Fail to connect to mongo : %s", err.Error()))
		return err
	}
	m.Logger.Info("Connected to mongo")
	m.Session = session
	m.Connected = true
	return nil
}

func newMongo(address string, logger log15.Logger) *Mongo {
	mongo := Mongo{
		Address: address,
		Logger:  logger,
	}
	mongo.Connect()
	return &mongo
}
