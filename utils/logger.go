package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

// InitLogger inicializa o logger padrão
func InitLogger() {
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.InfoLevel)
}
