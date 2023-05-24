package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/qu0ta/todo-app"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User
	if err := c.BindJSON(&input); err != nil {
		logrus.Fatalf("Incorrect input: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.CreateUser(input)
	if err != nil {
		logrus.Fatalf("Error while creating a user: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) signIn(c *gin.Context) {

}
