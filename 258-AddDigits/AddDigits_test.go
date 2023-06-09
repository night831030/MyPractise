package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestAddDigits(t *testing.T) {
	seedNum := time.Now().UnixNano()
	rand.Seed(seedNum)
	date := rand.Intn(1001)
	expectans := (date-1)%9 + 1

	if addDigits(date) != expectans {
		t.Errorf("func is fail, expect %d but get %d", expectans, addDigits(date))
	}
}
