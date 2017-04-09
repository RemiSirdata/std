package main

import (
	"gopkg.in/kataras/iris.v6"
	//"gopkg.in/kataras/iris.v6/middleware/basicauth"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
	"flag"
	"fmt"
	"gopkg.in/kataras/iris.v6/middleware/basicauth"
	"time"
	"os"
	"html/template"
	"encoding/csv"
	"strconv"
	"strings"
)

func main() {

	mongoAddr := flag.String("mongo-addr", "localhost", "Mongo address like server1.example.com,server2.example.com")
	createAdmin := flag.String("create-admin", "", "Create admin and stop")
	importFromFile := flag.String("user-file", "", "CSV file to import format name, maxguest, hasAccessToDinner, login, pass, isAdmin")
	bind := flag.String("bind", ":8888", "Bind port")
	flag.Parse()

	appContext := newAppContext(*mongoAddr)

	if *createAdmin != "" {
		password := "_" + *createAdmin
		err := appContext.Mongo.Session.DB(DB_NAME).C(GUEST_COLLECTION_NAME).Insert(Guest{
			Login:            *createAdmin,
			Password:         password,
			IsAdmin:          true,
			Status:           STATUS_WAITING_RESPONSE,
			MaxGuests:        2,
			GuestCount:       1,
			LastUpdateStatus: time.Now(),
			HasAccessToParty: true,
		})
		if err != nil {
			appContext.Logger.Crit(err.Error())
		} else {
			appContext.Logger.Info(fmt.Sprintf("User %s created with password %s", *createAdmin, password))
		}
		os.Exit(0)
	}

	if *importFromFile != "" {
		appContext.Logger.Info("Import user file")
		f, err := os.Open(*importFromFile)
		if err != nil {
			appContext.Logger.Crit(err.Error())
			os.Exit(1)
		}
		defer f.Close() // this needs to be after the err check
		reader := csv.NewReader(f)
		reader.Comma = ';'
		reader.LazyQuotes = true
		lines, err := reader.ReadAll()
		if err != nil {
			appContext.Logger.Crit(err.Error())
			os.Exit(1)
		}
		c := appContext.Mongo.Session.DB(DB_NAME).C(GUEST_COLLECTION_NAME)
		for _, line := range lines {
			if line[3] != "" {
				appContext.Logger.Info(fmt.Sprintf("%v", line))
				count, _ := strconv.Atoi(line[1])
				err := c.Insert(Guest{
					Name:             strings.TrimSpace(line[0]),
					Login:            strings.TrimSpace(line[3]),
					Password:         strings.TrimSpace(line[4]),
					IsAdmin:          line[5] == "1",
					Status:           STATUS_WAITING_RESPONSE,
					MaxGuests:        count,
					GuestCount:       0,
					HasAccessToParty: line[2] == "1",
				})
				if err != nil {
					appContext.Logger.Error(fmt.Sprintf("Fail to insert %s : %s", line[3], err.Error()))
				}
			}
		}
		os.Exit(0)
	}

	users := map[string]string{}
	for _, guest := range appContext.Guests {
		users[guest.Login] = guest.Password
	}
	authentication := basicauth.Default(users)

	app := iris.New(iris.Configuration{})
	app.StaticServe("./public", "/assets")

	tpl := view.HTML("./templates", ".html")
	tpl.Funcs(template.FuncMap{"N": N})

	app.Adapt(
		iris.DevLogger(),
		httprouter.New(),
		tpl,
	)
	app.Config.Charset = "UTF-8"

	routes := []Route{
		{ROUTE_HOME, iris.MethodGet, homeHandler},
		{ROUTE_LIST, iris.MethodGet, listHandler},
		{ROUTE_LIST_GUEST, iris.MethodGet, listGuestHandler},
		{ROUTE_AJAX_RSVP, iris.MethodPost, ajaxRsvpHandler},
	}

	for _, route := range routes {
		r := route
		appContext.Logger.Info(fmt.Sprintf("Add route %s", r.Path))
		app.HandleFunc(route.Method, route.Path, authentication, func(ctx *iris.Context) {
			r.Handler(ctx, appContext)
		})
	}

	app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Writef("404 ERROR PAGE")
	})

	app.Listen(*bind)
}

type Route struct {
	Path    string
	Method  string
	Handler Handler
}

type Handler func(ctx *iris.Context, appContext *AppContext)
