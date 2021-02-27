// +build integration unit

package test

import (
	"testing"
	"time"

	"github.com/icrowley/fake"
	"github.com/remnv/go-boiler/internal/helpers"
	"github.com/remnv/go-boiler/internal/model"
)

func FakePlayer(t *testing.T, cb func(player model.Player) model.Player) model.Player {
	t.Helper()
	today := getBod(time.Now().UTC())
	fakePl := model.Player{
		UserId:    helpers.StringPointer(fake.CharactersN(10)),
		Name:      helpers.StringPointer(fake.WordsN(4)),
		Level:     helpers.StringPointer("1"),
		Job:       helpers.StringPointer(fake.CharactersN(5)),
		CreatedAt: &today,
		UpdatedAt: &today,
	}
	if cb != nil {
		fakePl = cb(fakePl)
	}
	return fakePl
}

// getBod gets time at the beginning of the day.
func getBod(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}
