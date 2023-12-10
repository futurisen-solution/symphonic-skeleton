package log

import (
	"github.com/fwidjaya20/symphonic/contracts/log"
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/sirupsen/logrus"
)

func Logger() log.Logger {
	config := facades.Config()
	logger := facades.Logger()

	if config.GetString("app.env") == "production" {
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "time",
				logrus.FieldKeyLevel: "severity",
				logrus.FieldKeyMsg:   "message"},
		})
	}

	return logger
}
