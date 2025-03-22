package config

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func setupLogger() {
	sm := GetSettingsManager()

	logFileName, err := sm.GetSetting(LogFileName)

	if err != nil {
		logrus.Warn("Error while getting log file name : %s", err.Error())
	}

	file, err := os.OpenFile(logFileName.Value.(string), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Warn("Failed to log to file, using only console")
	}

	multiWriter := io.MultiWriter(os.Stdout, file)

	logrus.SetOutput(multiWriter)

	logLevelSetting, err := sm.GetSetting(LogLevel)

	if err != nil {
		logrus.Warn("Couldnt' extract log level : %s", err.Error())
	}

	logLevel, err := logrus.ParseLevel(logLevelSetting.Value.(string))

	if err != nil {
		logrus.Warn("Incorrect value for LOG_LEVEL, using default log level - info")
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.Level(logLevel))
	}

	logrus.Info("Logger is set...")
}
