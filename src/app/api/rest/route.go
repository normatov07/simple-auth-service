package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/normatov07/auth-service/app/api/rest/controller"
)

func (s *Server) Routes() http.Handler {
	r := gin.Default()

	var authCtr controller.AuthController

	auth := r.Group("/api")
	{
		auth.POST("/login", authCtr.Login)
		auth.POST("/register", authCtr.Register)
		//TODO: Also implement refresh token
	}

	return r
}
