package main

import (
	"gopkg.in/kataras/iris.v6"
	"bytes"
	"fmt"
)

func listGuestHandler(ctx *iris.Context, appContext *AppContext) {
	// logged username
	guest := appContext.FindUserByLogin(ctx.GetString("user"))

	if !guest.IsAdmin {
		ctx.SetStatusCode(iris.StatusNotFound)
		return
	}

	guests, err := loadGuest(appContext.Mongo)

	var buffer bytes.Buffer
	buffer.WriteString(`<html><head><link rel="stylesheet" href="/assets/css/plugins/bootstrap.css"></head><body>
	<table class="table table-striped"><tr><th>Name</th><th>Login</th><th>Pass</th><th>Status</th><th>Guests</th><th>Max guests</th><th>Last Update Status</th><th>Message</th></tr>`)
	if err == nil {
		for _, guest := range guests {

			buffer.WriteString(fmt.Sprintf(`<tr>
			<td>%s</td>
			<td>%s</td>
			<td>%s</td>
			<td>%s</td>
			<td>%d</td>
			<td>%d</td>
			<td>%s</td>
			<td>%s</td>
		</tr>`, guest.Name, guest.Login, guest.Password, guest.Status, guest.GuestCount, guest.MaxGuests, guest.LastUpdateStatus.Format("2006-01-02 15:04:05"), guest.Message))
		}
	}
	buffer.WriteString("</table></body></html>")
	ctx.Write(buffer.Bytes())
}
