package Handler

import (
	"github.com/kataras/iris"
	"sync/Database"
)

func DonateHandler(ctx iris.Context) {

	link := "pgsql"
	sql := "select field_de_store_id from opp.orders"
	res := Database.Querys(link, sql)

	ctx.JSON(res)


	// let's pass a value to the next handler
	// Values is the way handlers(or middleware) are communicating between each other.
	ctx.Values().Set("donate_url", "https://github.com/kataras/iris#-people")
	ctx.Next() // in order to execute the next handler in the chain, look donate route.
}

func DonateFinishHandler(ctx iris.Context) {
	link := "pgsql"
	insert := "insert into opp.time_interval (start_time,end_time,interval) values ('12:00:00','23:00:00',49)"
	res := Database.Exec(link, insert)
	ctx.JSON(res)
	// values can be any type of object so we could cast the value to a string
	// but iris provides an easy to do that, if donate_url is not defined, then it returns an empty string instead.
	donateURL := ctx.Values().GetString("donate_url")
	//ctx.Writef(donateURL)
	ctx.Application().Logger().Infof("donate_url value was: " + donateURL)
	//ctx.Writef("\n\nDonate sent(?).")
}
