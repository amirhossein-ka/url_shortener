package cmd

import (
	"flag"

	"github.com/amirhossein-ka/url_shortener/config"
	"github.com/amirhossein-ka/url_shortener/controller/gin"
	"github.com/amirhossein-ka/url_shortener/database"
	"github.com/amirhossein-ka/url_shortener/service"
)

var (
	cfg     = config.Config{}
	cfgPath string
)

func Start() error {
	if err := config.Parse(&cfg, cfgPath); err != nil {
		panic(err)
	}

	rdb, err := database.New(cfg)
	if err != nil {
		panic(err)
	}

	srv := service.New(rdb)

	ctrl := gin.New(cfg, srv)
	ctrl.Start(":8080")

	return nil
}

func Parse() {
	if !flag.Parsed() {
		flag.StringVar(&cfgPath, "config", "./config.json", "path to json config file")
		flag.StringVar(&cfgPath, "c", "./config.json", "path to json config file (short)")
	}
}
