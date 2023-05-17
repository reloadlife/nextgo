package webserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/reloadlife/nextgo/internal/middleware"
	"github.com/reloadlife/nextgo/internal/template"
	"github.com/reloadlife/nextgo/nextjs"
	"github.com/reloadlife/nextgo/routes"
	"github.com/sarulabs/di"
	log "github.com/sirupsen/logrus"
	"io/fs"
	"net/http"
)

var (
	tpl = template.GetTemplate()
)

func run(handler *gin.Engine) {
	err := handler.Run(
		fmt.Sprintf("%s:%d",
			tpl.GetHost(), tpl.GetPort()),
	)
	if err != nil {
		log.Errorf("Failed to start HTTP server: %s", err.Error())
	}
}

func runSSL(handler *gin.Engine) {
	err := handler.RunTLS(
		fmt.Sprintf("%s:%d",
			tpl.GetHost(), tpl.GetSSLPort()),
		tpl.GetCertificate(),
		tpl.GetCertificateKey(),
	)
	if err != nil {
		log.Errorf("Failed to start SSL server: %s", err.Error())
	}
}

func SetupWebserver() *di.Def {
	return &di.Def{
		Name: "gin",
		Build: func(_ di.Container) (interface{}, error) {
			gin.SetMode(gin.DebugMode)

			if tpl.IsProduction() {
				gin.SetMode(gin.ReleaseMode)
			}

			handler := gin.New()
			handler.Use(gin.Recovery(), gin.Logger())
			handler.Use(middleware.Middlewares...)
			routes.Routes(handler)
			handler.NoRoute(routes.NoRoute)

			if tpl.IsProduction() {
				distFS, err := fs.Sub(nextjs.NextFS(), "out")
				if err != nil {
					log.Fatal(err)
				}
				handler.StaticFS("/", http.FS(distFS))
			} else {
				// overwrite no route handler for development,
				// it will reverse proxy everything that doesn't match any route
				// on to the NextJS server running by NPM.
				handler.NoRoute(nextjs.ReverseProxy)
			}
			go run(handler)

			if tpl.IsSSL() {
				go runSSL(handler)
			}
			return nil, nil
		},
	}
}
