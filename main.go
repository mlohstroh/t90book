package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/mlohstroh/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load()
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)

	port := 5000
	if len(os.Getenv("PORT")) > 0 {
		aPort, err := strconv.Atoi(os.Getenv("PORT"))
		if err == nil {
			port = aPort
		}
	}

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		log, err := log.WithFields(
			log.Fields{
				"ip":        param.ClientIP,
				"timestamp": param.TimeStamp,
				"method":    param.Method,
				"path":      param.Path,
				"status":    param.StatusCode,
				"latency":   param.Latency,
				"ua":        param.Request.UserAgent(),
				"error":     param.ErrorMessage,
			},
		).String()

		if err != nil {
			return err.Error()
		}

		return log
	}))

	router.Use(gin.Recovery())
	setupRoutes(router)
	router.Run(fmt.Sprintf(":%d", port))
}

func alive(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func setupRoutes(router *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	router.Use(cors.New(corsConfig))

	pprof.Register(router)
	router.GET("/api/current-id", getCurrentID)

	router.GET("/api/_alive", alive)
	router.Static("/public", "public")
	router.Static("/static", "public/static")
	router.StaticFile("/", "public/index.html")
	router.NoRoute(func(c *gin.Context) {
		c.File("public/index.html")
	})
}
