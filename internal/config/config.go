package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadConfig() {
	err := godotenv.Load(".env")

	if err != nil {
		logrus.Panic("Couldn't get env file.")
	}

	setupLogger()
	// initializes settings manager, but we do not need the instance rn
	_ = GetSettingsManager()
}
