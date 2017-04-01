package main

import (
	"gopkg.in/kataras/iris.v6"
	//_ "github.com/kataras/go-template/django"
	"fmt"
)

func homeHandler(ctx *iris.Context, appContext *AppContext) {
	// logged username
	//username := ctx.GetString("user")
	guest := appContext.FindUserByLogin(ctx.GetString("user"))

	appContext.Logger.Info("Home handler")

	selectStatus := map[string]string{
		STATUS_RECEPTION: "Je serai présent au vin d'honneur",
		STATUS_REFUSED:   "je ne pourrai pas être présent",
	}
	if guest.HasAccessToParty {
		selectStatus = map[string]string{
			STATUS_ACCEPT_ALL: "Je serai présent au vin d'honneur et au dîner",
			STATUS_RECEPTION:  "Je serai présent au vin d'honneur",
			STATUS_DINNER:     "Je serai présent au dîner",
			STATUS_REFUSED:    "je ne pourrai pas être présent",
		}
	}

	err := ctx.Render("index.html", Page{
		AppContext:   appContext,
		HotelList:    getHotelList(),
		Guest:        guest,
		SelectStatus: selectStatus,
	}, iris.RenderOptions{"gzip": true})
	if err != nil {
		appContext.Logger.Error(fmt.Sprintf("Error rendering template %s", err))
	}
}
