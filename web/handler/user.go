package handler

import (
	"net/http"
	"startup-backend-api/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Index(c *gin.Context) {

	users, err := h.userService.GetAllUsers()

	if err != nil {

	}

	c.HTML(http.StatusOK, "user_index.html", gin.H{"users": users})

}