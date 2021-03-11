package controller_http

import (
	"amiera/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	GetAT(*gin.Context)
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
	acId, _ := strconv.ParseInt(c.Param("user_id"), 10, 0)
	accessToken, err := handler.atService.GetById(acId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) GetAT(c *gin.Context) {
	accessToken, err := handler.atService.GetAT()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}