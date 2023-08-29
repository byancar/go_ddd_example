package user

import "github.com/gin-gonic/gin"

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) create(c *gin.Context) {

}
