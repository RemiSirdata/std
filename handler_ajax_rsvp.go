package main

import (
	"gopkg.in/kataras/iris.v6"
	_ "github.com/kataras/go-template/django"
	"io/ioutil"
	"encoding/json"
)

func ajaxRsvpHandler(ctx *iris.Context, appContext *AppContext) {
	// logged username
	guest := appContext.FindUserByLogin(ctx.GetString("user"))

	body, _ := ioutil.ReadAll(ctx.Request.Body)

	rsvp := Rsvp{}
	err := json.Unmarshal(body, &rsvp)
	if err != nil {
		ctx.SetStatusCode(iris.StatusInternalServerError)
		return
	}
	guest.UpdateRsvp(&rsvp, appContext.Mongo)
	ctx.SetStatusCode(iris.StatusAccepted)
}

type Rsvp struct {
	Status     string `json:"status"`
	GuestCount int `json:"guest_count"`
	Message    string `json:"message"`
}
