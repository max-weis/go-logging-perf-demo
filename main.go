package main

import (
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
)

func main() {
	zapLogger := initZap()
	zeroLogger := initZero()
	e := echo.New()
	c := jaegertracing.New(e, nil)
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	defer c.Close()

	e.GET("/zerolog", func(c echo.Context) error {
		zeroLogger.Print("zerologger")
		return c.String(http.StatusOK, "Hello, zerolog!")
	})

	e.GET("/zap", func(c echo.Context) error {
		zapLogger.Info("zap")
		return c.String(http.StatusOK, "Hello, zap!")
	})

	e.GET("/logrus", func(c echo.Context) error {
		logrus.Print("logrus")
		return c.String(http.StatusOK, "Hello, logrus!")
	})

	e.GET("/stdlib", func(c echo.Context) error {
		log.Println("stdlib")
		return c.String(http.StatusOK, "Hello, stdlib!")
	})

	e.Logger.Fatal(e.Start(":8080"))

}

func initZero() zerolog.Logger {
	return zerolog.New(os.Stderr)
}

func initZap() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return logger
}
