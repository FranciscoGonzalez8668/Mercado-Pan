package main

import (
	"pan/app"
	"pan/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
