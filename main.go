package main

import (
	//"fmt"
	//"log"
	//"time"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(Auth)
	m.Use(render.Renderer())
	m.MapTo(db, (*DB)(nil))

	m.Group("/events", func(r martini.Router) {
		// EVENTS DATA
		r.Get("", GetEvents)
		r.Get("/:id", GetEvent)
		r.Post("", CreateEvent)
		r.Put("/:id", UpdateEvent)
		r.Delete("/:id", DeleteEvent)

		// DAILY DATA
		//r.Get("/:id/days", GetDays)
		//r.Get("/:id/days/:index", GetDay)
		//r.Post("/:id/days/", CreateDay)
		//r.Put("/:id/days/:index", UpdateDay)
		//r.Delete("/:id/days/:index", DeleteDay)

	})
	//events := []Event()

	m.Get("/", func() string {
		return "good"
	})

	m.Run()
}
