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
		r.Get("", GetEvents)
		r.Get("/:id", GetEvent)
		r.Post("", CreateEvent)
		r.Put("/:id", UpdateEvent)
		r.Delete("/:id", DeleteEvent)
	})

	m.Get("/", func() string {
		return "good"
	})

	m.Run()
}
