package csv

import (
	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

func init() {

}

// Logger provides the global logger
func Logger() *logrus.Logger {
	return log
}
