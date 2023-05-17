package nextjs

import (
	"embed"
	"fmt"
	"github.com/reloadlife/nextgo/internal/template"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os/exec"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di"
	log "github.com/sirupsen/logrus"
)

var (
	frontendPort = "3000"
	logger       = log.StandardLogger()
)

func init() {
	frontendPort = strconv.Itoa(template.GetTemplate().GetFrontendPort())
	logger.SetFormatter(&log.TextFormatter{})
}

//go:embed all:out
var nextFS embed.FS

func preDev() {
	log.Infof("Copying next.config.dev.js to next.config.js")
	cmd := exec.Command("cp", "next.config.dev.js", "next.config.js")
	cmd.Dir = "./nextjs"
	cmd.Stdout = logger.Out
	cmd.Stderr = logger.Out
	if err := cmd.Start(); err != nil {
		panic(err)
	}
}

func newFrontendServiceDev() (interface{}, error) {
	log.Infof("Starting frontend service in development mode")
	preDev()
	cmd := exec.Command("npm", "run", "dev", "--", "--port="+frontendPort)
	cmd.Dir = "./nextjs"
	cmd.Stdout = logger.Out
	cmd.Stderr = logger.Out
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return nil, nil
}

func newFrontendServiceProd() (interface{}, error) {
	log.Infof("Starting frontend service in Production mode")
	return nil, nil
}

func ReverseProxy(c *gin.Context) {
	remote, _ := url.Parse(fmt.Sprintf("http://127.0.0.1:%v", frontendPort))
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL = c.Request.URL
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
	}
	proxy.ServeHTTP(c.Writer, c.Request)
	proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, e error) {
		log.Errorf("Error while proxying: %s", e.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("Gateway is not there."))
	}
}

func NextFS() embed.FS {
	return nextFS
}

func SetupFrontendService() *di.Def {
	return &di.Def{
		Name: "frontend",
		Build: func(_ di.Container) (interface{}, error) {
			if template.GetTemplate().IsProduction() {
				return newFrontendServiceProd()
			}
			return newFrontendServiceDev()
		},
	}
}
