package controller_http

import (
	"amiera/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
}

type accessTokenHandler struct {
	atService service.Service
}

func NewHandler(atService service.Service) AccessTokenHandler {
	return &accessTokenHandler{
		atService: atService,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	//acId := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := handler.atService.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}