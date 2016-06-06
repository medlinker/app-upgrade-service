package logger

import (
	"app-upgrade-service/config"

	log "github.com/cihub/seelog"
)

// Logger log it
var Logger log.LoggerInterface

// InitLog init log config
func InitLog() {
	logFile := config.GetString("logFilePath", true, "../conf/log.xml")
	var err error
	Logger, err = log.LoggerFromConfigAsFile(logFile)
	if err != nil {
		panic(err)
	}
}
