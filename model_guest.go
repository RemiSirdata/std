package main

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

func loadGuest(mongo *Mongo) ([]Guest, error) {
	guests := []Guest{}
	err := mongo.Connect()
	if err == nil {
		err = mongo.Session.DB(DB_NAME).C(GUEST_COLLECTION_NAME).Find(map[string]string{}).Sort("last_update_status").All(&guests)
	}
	return guests, err
}

type Guest struct {
	Name             string `bson:"name"`
	Login            string `bson:"login"`
	Password         string `bson:"password"`
	IsAdmin          bool `bson:"is_admin"`
	LastConnection   time.Time `bson:"last_connection"`
	MaxGuests        int `bson:"max_guests"`
	HasAccessToParty bool `bson:"has_access_to_party"`

	Status           string `bson:"status"`
	GuestCount       int `bson:"guest_count"`
	LastUpdateStatus time.Time `bson:"last_update_status"`
	Message          string `bson:"message"`
}

func (g *Guest) UpdateRsvp(rsvp *Rsvp, mongo *Mongo) {
	g.LastConnection = time.Now()
	g.LastUpdateStatus = time.Now()
	g.Status = rsvp.Status
	g.GuestCount = rsvp.GuestCount
	if g.GuestCount > g.MaxGuests {
		g.GuestCount = g.MaxGuests
	}
	g.Message = rsvp.Message
	mongo.Connect()
	mongo.Session.DB(DB_NAME).C(GUEST_COLLECTION_NAME).Upsert(map[string]string{"login": g.Login}, bson.M{"$set": *g})
	mongo.Session.DB(DB_NAME).C(GUEST_COLLECTION_HISTORY).Insert(*g)
}
