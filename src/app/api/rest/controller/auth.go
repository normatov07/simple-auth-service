package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/normatov07/auth-service/common/response"
	"github.com/normatov07/auth-service/core/action"
	"github.com/normatov07/auth-service/core/service"
	"github.com/normatov07/auth-service/db/postgres"
)

type AuthController struct{}

func (c *AuthController) Login(ctx *gin.Context) {
	var acn action.Login

	if err := ctx.Bind(&acn); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	token, err := service.GetUserService(new(postgres.UserRepo)).Login(acn)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse(gin.H{"token": token}))
}

func (c *AuthController) Register(ctx *gin.Context) {
	var acn action.Register

	if err := ctx.Bind(&acn); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	token, err := service.GetUserService(new(postgres.UserRepo)).Register(acn)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse(gin.H{"token": token}))
}
