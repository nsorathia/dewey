package logging

//START HERE
import (
	log "github.com/sirupsen/logrus"
	"os"
)

//Logger is a common Logging solution
var Logger *log.Logger

func init() {
	logger := log.New()
	logger.Out = os.Stdout
	logger.Level = log.InfoLevel
	logger.Formatter = &log.TextFormatter{}
	Logger = logger;
}
