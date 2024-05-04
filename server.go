package main

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jaganathanb/dapps-web-server/config"
	"github.com/jaganathanb/dapps-web-server/logging"
	"github.com/jaganathanb/dapps-web-server/middlewares"
)

func InitServer(cfg *config.Config) {
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()

	RegisterRoutes(r, cfg)

	logger := logging.NewLogger(cfg)
	logger.Info(logging.General, logging.Startup, "Started", nil)

	err := r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
	if err != nil {
		logger.Fatal(logging.General, logging.Startup, err.Error(), nil)
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	r.SetTrustedProxies(nil)

	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(gin.Logger(), gin.CustomRecovery(middlewares.ErrorHandler))

	r.Use(static.Serve("/", static.LocalFile(cfg.Server.StaticContentPath, true)))
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File(path.Join(cfg.Server.StaticContentPath, "index.html"))
		} else {
			c.File(path.Join(cfg.Server.StaticContentPath, path.Join(dir, file)))
		}

	})
}

func main() {
	cfg := config.GetConfig()

	InitServer(cfg)
}
