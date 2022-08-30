package handler

import (
	"github.com/gin-gonic/gin"
	todo "github.com/m0n7h0ff/course-todo-app"
	"net/http"
)

func (h *Handler) singUp(c *gin.Context) {
	var input todo.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) singIn(c *gin.Context) {

}
