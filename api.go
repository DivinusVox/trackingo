package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/render"
)

/* MIDDLEWARE */
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

/* EVENTS API */
func GetEvents(r render.Render, db DB) {
	r.JSON(200, db.GetAll())
}

func GetEvent(r render.Render, db DB, params martini.Params) {
	key, ok := strconv.Atoi(params["id"])
	if ok != nil {
		fmt.Println(key)
	}
	r.JSON(200, db.Get(key))
}

func CreateEvent(req *http.Request, r render.Render, db DB) {
	t := DecodeJSONEvent(req)
	db.Add(t)
	r.JSON(201, t)
}

func UpdateEvent(req *http.Request, r render.Render, db DB, params martini.Params) {
	t := DecodeJSONEvent(req)
	db.Update(t)
	r.JSON(200, t)
}

func DeleteEvent(req *http.Request, r render.Render, db DB, params martini.Params) {
	key, ok := strconv.Atoi(params["id"])
	if ok != nil {
		fmt.Println(key)
	}
	t := db.Get(key)
	db.Delete(key)
	r.JSON(200, t)
}

/* EVENTS HELPER FUNCTIONS */
func DecodeJSONEvent(req *http.Request) *Event {
	decoder := json.NewDecoder(req.Body)
	var t Event
	decoder.Decode(&t)
	fmt.Println(t)
	return &t
}
