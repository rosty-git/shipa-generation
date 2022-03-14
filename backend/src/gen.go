package src

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shipa-gen/src/cloudformation"
	"shipa-gen/src/crossplane"
	"shipa-gen/src/shipa"
)

func Generate(c *gin.Context) {
	var cfg shipa.Config

	if err := c.BindJSON(&cfg); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	var data *shipa.Result
	switch cfg.Provider {
	case "crossplane":
		data = crossplane.Generate(cfg)
	case "cloudformation":
		data = cloudformation.Generate(cfg)
	default:
		c.IndentedJSON(http.StatusBadRequest, errors.New("not supported provider"))
		return
	}

	if data == nil {
		c.IndentedJSON(http.StatusNoContent, errors.New("no data"))
		return
	}

	rawData := []byte(data.Content)
	extraHeaders := map[string]string{
		"Content-Disposition": fmt.Sprintf(`attachment; filename="%s"`, data.Name),
	}
	c.DataFromReader(http.StatusOK, int64(len(rawData)), "text/plain; charset=utf-8", bytes.NewBufferString(data.Content), extraHeaders)
}
