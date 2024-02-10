package main

import (
	"github.com/fmuharam25/tutorial-golang-http-json/database"
	"github.com/fmuharam25/tutorial-golang-http-json/web"
)

func main() {
	database.Migrate()
	web.Serve()
}
