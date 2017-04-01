package main

import (
	"gopkg.in/kataras/iris.v6"
)

func listHandler(ctx *iris.Context, appContext *AppContext) {
	// logged username
	//username := ctx.GetString("user")
	ctx.Redirect("/#list", iris.StatusFound)
}
