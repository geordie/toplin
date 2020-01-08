package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tylerb/graceful"

	"github.com/geordie/toplin/application"
)

func newConfig() (*viper.Viper, error) {

	c := viper.New()
	c.SetConfigName("config.yaml")
	c.AddConfigPath(".")
	c.SetDefault("cookie_secret", "FWnQhq62ylAMfOAH")
	c.SetDefault("http_addr", ":8880")
	c.SetDefault("http_cert_file", "")
	c.SetDefault("http_key_file", "")
	c.SetDefault("http_drain_interval", "1s")
	c.SetDefault("joplin_port", "41184")

	err := c.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	c.AutomaticEnv()

	return c, nil
}

func main() {
	config, err := newConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	app, err := application.New(config)
	if err != nil {
		logrus.Fatal(err)
	}

	middle, err := app.MiddlewareStruct()
	if err != nil {
		logrus.Fatal(err)
	}

	serverAddress := config.Get("http_addr").(string)

	certFile := config.Get("http_cert_file").(string)
	keyFile := config.Get("http_key_file").(string)
	drainIntervalString := config.Get("http_drain_interval").(string)

	drainInterval, err := time.ParseDuration(drainIntervalString)
	if err != nil {
		logrus.Fatal(err)
	}

	srv := &graceful.Server{
		Timeout: drainInterval,
		Server:  &http.Server{Addr: serverAddress, Handler: middle},
	}

	logrus.Infoln("Running HTTP server on " + serverAddress)

	if certFile != "" && keyFile != "" {
		err = srv.ListenAndServeTLS(certFile, keyFile)
	} else {
		err = srv.ListenAndServe()
	}

	if err != nil {
		logrus.Fatal(err)
	}
}
