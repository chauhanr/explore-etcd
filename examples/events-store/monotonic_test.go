package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestMonotonicSource(t *testing.T) {
	tm := time.Unix(1000000, 0)
	entropy := rand.New(rand.NewSource(tm.UnixNano()))
	ulidSource := NewMonotonicULIDsource(entropy)

	id, _ := ulidSource.New(tm)
	expected := "0000XSNJG0T8CNRGXPSBZSA1PY"
	if id.String() != expected {
		t.Errorf("Expected ulid %s but got %s\n", expected, id)

	}

}
