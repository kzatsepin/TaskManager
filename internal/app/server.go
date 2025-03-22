package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kzatsepin/TaskManager/internal/config"
	"github.com/sirupsen/logrus"
)

func RunServer() error {
	config.LoadConfig()

	router := gin.Default()

	setupRoutes(router)

	sm := config.GetSettingsManager()

	ginServerIP, err := sm.GetSetting(config.GinServerIP)

	if err != nil {
		logrus.Fatalf("Cannot run Gin Web Server : %s", err.Error())
	}

	ginServerPort, err := sm.GetSetting(config.GinServerPort)

	if err != nil {
		logrus.Fatalf("Cannot run Gin Web Server : %s", err.Error())
	}

	fullServerAddress := ginServerIP.Value.(string) + ":" + ginServerPort.Value.(string)

	logrus.Info("Running Gin Web Server at %s", fullServerAddress)

	err = router.Run(fullServerAddress)

	if err != nil {
		logrus.Fatalf("Couldn't run Gin Web Server : %s", err.Error())
	}

	logrus.Info("Leave with success")
	return nil
}

func setupRoutes(r *gin.Engine) {
	logrus.Info("Setting up routes...")

	logrus.Info("Routes are set up, success")

}
