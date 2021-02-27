package helpers

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func MustJsonString(obj interface{}) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		logrus.WithError(err).Warn("method", "MustJsonString")
		return ""
	}

	jsonStr := string(bytes)

	return jsonStr
}
