package main

import (
	//"encoding/json"
	"fmt"
	//"log"
        "strconv"
        "strings"
	//"time"
        "net/http"

        "github.com/go-martini/martini"
        "github.com/martini-contrib/render"
        "github.com/martini-contrib/auth"
)


func main() {
        m := martini.Classic()
        m.Use(Auth)
        m.Use(render.Renderer())
        m.MapTo(db, (*DB)(nil))

        m.Group("/events", func(r martini.Router){
            // EVENTS DATA
            r.Get("", GetEvents)
            r.Get("/:id", GetEvent)
            r.Post("", CreateEvent)
            //r.Put("/:id", UpdateEvent)
            //r.Delete("/:id", DeleteEvent)

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

        m.Post("/", Encoding, func(r render.Render) string {
                fmt.Println()
                return "WORK"
                //r.JSON(200, {})
        })
        m.Run()
}

func Encoding(res http.ResponseWriter, req *http.Request) {
    ct := strings.ToLower(req.Header.Get("Content-Type"))
    if ct != "application/json" {
        http.Error(res, "Content-Type must be `application/json`", 406)
    }
}

func Auth(res http.ResponseWriter, req *http.Request) {
    secret := "abc123"
    if auth.SecureCompare(req.Header.Get("Validation"), secret) != true {
        http.Error(res, "Unauthorized", 401)
    }
}

func GetEvents(req *http.Request, r render.Render, db DB) {
    r.JSON(200, db.GetAll())
}

func GetEvent(req *http.Request, r render.Render, db DB, params martini.Params) {
    key, ok := strconv.Atoi(params["id"])
    if ok != nil {
        fmt.Println(key)
    }
    r.JSON(200, db.Get(key))
}

func CreateEvent(req *http.Request, r render.Render, db DB, params martini.Params) {
    fmt.Println(req)
}
