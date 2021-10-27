package main

import (
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"log"
	"os"
)

var (
	zeroLogger zerolog.Logger
	zapLogger  *zap.Logger
)

func main() {
	// init logger
	zapLogger = initZap()
	zeroLogger = initZero()

	// init echo
	e := echo.New()
	c := jaegertracing.New(e, nil)
	defer c.Close()
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	// set logger
	loggerType := os.Getenv("LOGGER")
	var handler echo.HandlerFunc

	switch loggerType {
	case "zap":
		handler = zapHandler
	case "logrus":
		handler = logrusHandler
	case "zero":
		handler = zeroHandler
	case "stdlib":
		handler = stdlibHandler
	}

	e.GET("/", handler)

	e.Logger.Fatal(e.Start(":8080"))
}

func zeroHandler(c echo.Context) error {
	zeroLogger.Print("zerologger")
	return nil
}

func zapHandler(c echo.Context) error {
	zapLogger.Info("zap")
	return nil
}

func logrusHandler(c echo.Context) error {
	logrus.Print("logrus")
	return nil
}

func stdlibHandler(c echo.Context) error {
	log.Println("stdlib")
	return nil
}

func initZero() zerolog.Logger {
	return zerolog.New(os.Stderr)
}

func initZap() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return logger
}
