package controller_http

import (
	"amiera/src/domain/access_token"
	"amiera/src/service"
	"amiera/src/utils/utils_errors"
	"amiera/src/utils/utils_responses"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	GetAT(*gin.Context)
	GetOptionById(*gin.Context)
	CreateToken(*gin.Context)
	UpdateExpiration(*gin.Context)
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
	userId, _ := strconv.ParseInt(c.Param("user_id"), 10, 0)
	accessToken, err := handler.atService.GetById(userId)
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

func (handler *accessTokenHandler) GetOptionById(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Param("filter"), 10, 0)
	accToken, err := handler.atService.GetOptionById(userId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accToken)
}

func (handler *accessTokenHandler) CreateToken(c *gin.Context) {
	var accToken access_token.AccessToken

	if err := c.ShouldBind(&accToken); err != nil {
		errRest := utils_errors.CustomBadRequestError("invalid json body")
		c.JSON(errRest.Status, errRest)
		return
	}

	// read body
	fmt.Println("accToken", accToken)
	err := handler.atService.CreateToken(accToken)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, utils_responses.CustomSuccessResponse("token is created"))
}

func (handler *accessTokenHandler) UpdateExpiration (c *gin.Context) {
	var accToken access_token.AccessToken

	if err := c.ShouldBind(&accToken); err != nil {
		errRest := utils_errors.CustomBadRequestError("invalid json body")
		c.JSON(errRest.Status, errRest)
		return
	}

	fmt.Println("accToken", accToken)
	isPartial := c.Request.Method == http.MethodPatch
	err := handler.atService.UpdateExpiration(accToken, isPartial)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, utils_responses.CustomSuccessResponse("OK"))
}