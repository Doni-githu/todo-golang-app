package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func (h *Handler) createList(сtx *gin.Context) {

}

func (h *Handler) getAllLists(c *gin.Context) {
	users := []User{
		{Name: "Doniyor", ID: 1},
		{Name: "Fotima", ID: 2},
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *Handler) getListById(сtx *gin.Context) {

}

func (h *Handler) updateList(сtx *gin.Context) {

}

func (h *Handler) deleteList(сtx *gin.Context) {

}
