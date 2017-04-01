package main

import (
	"github.com/inconshreveable/log15"
	"os"
)

func newAppContext(mongoAddr string) *AppContext {

	logger := log15.New()
	logger.SetHandler(log15.LvlFilterHandler(log15.LvlInfo, log15.StdoutHandler))
	mongo := newMongo(mongoAddr, logger)
	guests, err := loadGuest(mongo)
	if err != nil {
		logger.Crit(err.Error())
		os.Exit(1)
	}
	appContext := AppContext{
		mongoAddr,
		mongo,
		guests,
		logger,
	}
	return &appContext
}

type AppContext struct {
	MongoAddr string
	Mongo     *Mongo

	Guests []Guest

	Logger log15.Logger
}

func (a *AppContext) FindUserByLogin(login string) *Guest {
	for _, guest := range a.Guests {
		if guest.Login == login {
			return &guest
		}
	}
	return &Guest{}
}
