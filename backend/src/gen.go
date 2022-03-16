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

type GenerateAppsResponse struct {
	Apps   []AppResponse `json:"apps"`
	Errors []AppError    `json:"errors,omitempty"`
}

type AppError struct {
	AppName string `json:"appName"`
	Error   string `json:"error"`
}

type AppResponse struct {
	AppName string       `json:"appName"`
	Files   []FileResult `json:"files"`
}

type FileResult struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

func GenerateApps(c *gin.Context) {
	var cfg shipa.AppsConfig

	if err := c.BindJSON(&cfg); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if len(cfg.Apps) == 0 {
		c.IndentedJSON(http.StatusBadRequest, errors.New("empty input data"))
		return
	}

	var out GenerateAppsResponse
	for _, app := range cfg.Apps {
		file, err := generateApp(app)
		if err != nil {
			out.Errors = append(out.Errors, AppError{
				AppName: app.AppName,
				Error:   err.Error(),
			})
			continue
		}

		out.Apps = append(out.Apps, AppResponse{
			AppName: app.AppName,
			Files: []FileResult{
				{Filename: file.Name, Content: file.Content},
			},
		})
	}

	c.JSON(http.StatusOK, out)
}

func generateApp(cfg shipa.Config) (*shipa.Result, error) {
	var data *shipa.Result
	switch cfg.Provider {
	case "crossplane":
		data = crossplane.Generate(cfg)
	case "cloudformation":
		data = cloudformation.Generate(cfg)
	default:
		return nil, fmt.Errorf("not supported provider: %s", cfg.Provider)
	}

	return data, nil
}
