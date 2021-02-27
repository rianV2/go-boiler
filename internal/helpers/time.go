package helpers

import (
	"time"

	"github.com/sirupsen/logrus"
)

func GetBodTime(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func GetValidTimezone(tz string) *string {

	loc, err := time.LoadLocation(tz)
	if err != nil {
		logrus.WithError(err).Warn("method", "GetValidTimezone")
		return nil
	}

	return StringPointer(loc.String())
}
