package logging

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

var Logging *zap.SugaredLogger
var Logger *zap.Logger

var ZapLogger *zap.Logger

func init() {
	zap.RegisterSink("winfile", newWinFileSink)
}

func Init() {
	var err error

	zap.RegisterSink("winfile", newWinFileSink)
	cfg := zap.NewProductionConfig()

	cfg.Level = zap.NewAtomicLevel()
	cfg.OutputPaths = []string{
		"winfile:///" + filepath.Join(getPWD(), "logs/app.log"),
		"stdout",
	}
	cfg.ErrorOutputPaths = []string{
		"winfile:///" + filepath.Join(getPWD(), "logs/app.log"),
		"stderr",
	}

	Logger, err = cfg.Build()

	if err != nil {
		panic(err)
	}
	//defer Logger.Sync() // flushes buffer, if any
	Logging = Logger.Sugar()
}

func ZapInstanceForGin() {
	var err error

	cfg := zap.NewProductionConfig()

	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	cfg.OutputPaths = []string{
		"winfile:///" + filepath.Join(getPWD(), "logs/gin.log"),
		//"stdout",
	}
	cfg.ErrorOutputPaths = []string{
		"winfile:///" + filepath.Join(getPWD(), "logs/gin.log"),
		"stderr",
	}

	ZapLogger, err = cfg.Build()

	if err != nil {
		panic(err)
	}
	//defer ZapLogger.Sync() // flushes buffer, if any
}

func newWinFileSink(u *url.URL) (zap.Sink, error) {
	// Remove leading slash left by url.Parse()
	return os.OpenFile(u.Path[1:], os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
}

func getPWD() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}
