package main

import (
    "sync"
    "time"
)

// DATABASE
type DB interface {
    Get(id int) *Event
    GetAll() []*Event
    Add(e *Event) (int, error)
    Update(e *Event) error
    Delete(id int)
}

type eventsDB struct {
    sync.RWMutex
    m   map[int]*Event
    seq int
}

var db DB

func init() {
    db = &eventsDB{
        m: make(map[int]*Event),
    }
    start := time.Date(2014, time.May, 20, 0, 0, 0, 0, time.UTC)
    end := time.Date(2015, time.May, 20, 0, 0, 0, 0, time.UTC)
    event := Event{1, "tester", "pushups", start, end, 1, []Record{}}
    db.Add(&event)
    db.Add(&event)
    db.Add(&event)
    db.Add(&event)
}

func (db *eventsDB) GetAll() []*Event {
    db.RLock()
    defer db.RUnlock()
    if len(db.m) == 0 {
        return nil
    }
    ar := make([]*Event, len(db.m))
    i := 0
    for _, v := range db.m {
      ar[i] = v
      i++
    }
    return ar
}

func (db *eventsDB) Get(id int) *Event {
    db.RLock()
    defer db.RUnlock()
    return db.m[id]
}

func (db *eventsDB) Add(e *Event) (int, error) {
    db.Lock()
    defer db.Unlock()
    db.seq++
    e.Id = db.seq
    db.m[e.Id] = e
    return e.Id, nil
}

func (db *eventsDB) Update(e *Event) error {
    db.Lock()
    defer db.Unlock()
    db.m[e.Id] = e
    return nil
}

func (db *eventsDB) Delete(id int) {
    db.Lock()
    defer db.Unlock()
    delete(db.m, id)
}

// MODELS
type Record struct {
        Id     int
	Date   time.Time
	Amount int
	Done   bool
}

type Event struct {
        Id        int
	Title     string
	Action    string
	Start     time.Time
	End       time.Time
	Increment int
	Days      []Record
}
