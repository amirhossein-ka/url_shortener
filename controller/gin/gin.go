package gin

import (
	"github.com/amirhossein-ka/url_shortener/config"
	"github.com/amirhossein-ka/url_shortener/controller"
	"github.com/amirhossein-ka/url_shortener/service"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	r       *gin.Engine
	handler *handler
}

type handler struct {
	service service.Service
}

func New(cfg config.Config, srv service.Service) controller.Controller {
	return &Gin{
		r: gin.New(),
		handler: &handler{
			service: srv,
		},
	}
}

func (g *Gin) Start(addr string) error {
	g.r.Use(gin.Logger())
	g.r.Use(gin.Recovery())

	g.routing()

	return g.r.Run(addr)
}

func (g *Gin) Stop() {

}
