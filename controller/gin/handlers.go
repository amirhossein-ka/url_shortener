package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) redirect(c *gin.Context) {
	path := c.Param("path")
	url, err := h.service.GetUrlAddr(c.Request.Context(), path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "err": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, url)
}

func (h *handler) shortener(c *gin.Context) {
	var url ShortenerReq
	if err := c.BindJSON(&url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "err": err.Error()})
		return
	}

	key, err := h.service.ShortUrlAddress(c.Request.Context(), url.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "err": err.Error()})
		return
	}

	fmt.Println(c.Request.Host)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"path": fmt.Sprintf("http://%s/go/%s", c.Request.Host, key),
	})
}
