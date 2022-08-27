package gin

func (g *Gin) routing() {
	g.r.GET("/go/:path", g.handler.redirect)
	g.r.POST("/shortener", g.handler.shortener)
}
