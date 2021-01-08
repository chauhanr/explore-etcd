package main

import "time"

//FakeEvent is an event we will send to the etcd store.
type FakeEvent struct {
	ID        string    `json:"id"`
	TimeStamp time.Time `json:"timestamp"`
	Data      int       `json:"data"`
}
