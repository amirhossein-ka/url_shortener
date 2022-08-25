package gin

import (
	"github.com/amirhossein-ka/url_shortener/config"
	"github.com/amirhossein-ka/url_shortener/controller"
	"github.com/amirhossein-ka/url_shortener/service"
)

type Gin struct {
    service *service.Service
    cfg config.Config
}

func New(cfg config.Config, srv *service.Service) controller.Controller { 

    return &Gin{}
}


func (g *Gin) Start()  {
    
}

func (g *Gin) Stop()  {
    
}
